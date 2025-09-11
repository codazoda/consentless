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
		// Prevent caching of beacon responses
		w.Header().Set("Cache-Control", "no-store, max-age=0")

		// Determine the URL being counted (single method: query param `u`)
		url := r.URL.Query().Get("u")
		if url == "" {
			// Nothing to log; return no content
			w.WriteHeader(http.StatusNoContent)
			return
		}

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

	// Generate minimal JavaScript with dynamic domain and protocol.
	jsCode := fmt.Sprintf(`(function(){
  var img=new Image();
  img.src="%s://%s/?u="+encodeURIComponent(window.location.href)+"&rand="+Math.random();
})();`, protocol, host)

	// Serve JavaScript
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write([]byte(jsCode))
}
