var xhr = new XMLHttpRequest();
xhr.open('GET', "https://consentless.joeldare.com?rand=" + Math.random(), true);
xhr.setRequestHeader('X-Referrer', window.location.href);
xhr.send();
