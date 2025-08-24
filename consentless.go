package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/consentless.js", serveConsentlessJS)
	http.HandleFunc("/counter.js", serveConsentlessJS) // For backward compatibility

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Get the X-Referrer header
		url := r.Header.Get("X-Referrer")

		// Use minute precision in local time
		ts := time.Now().Format("2006-01-02 15:04")

		// Print CSV line to stdout
		fmt.Printf("%s,%s\n", url, ts)

		// Respond with JSON
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  true,
			"referrer": url,
		})
	})

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start HTTP server
	log.Printf("Starting HTTP server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("http server error: %v", err)
	}
}

// serveConsentlessJS serves the consentless.js script with dynamic domain and protocol
func serveConsentlessJS(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Get the host from the request to make the domain dynamic
	host := r.Host
	if host == "" {
		host = "localhost:8080" // fallback
	}

	// Determine protocol (check if TLS is being used)
	protocol := "http"
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		protocol = "https"
	}

	// Generate JavaScript with dynamic domain and protocol
	jsCode := fmt.Sprintf(`var xhr = new XMLHttpRequest();
xhr.open('GET', "%s://%s?rand=" + Math.random(), true);
xhr.setRequestHeader('X-Referrer', window.location.href);
xhr.send();
`, protocol, host)

	// Serve JavaScript
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write([]byte(jsCode))
}
