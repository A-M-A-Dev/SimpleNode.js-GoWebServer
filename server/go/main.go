package main

import (
	"fmt"
	"net/http"
	"strings"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	accepts := strings.ToLower(r.Header.Get("Accept"))

	if strings.Contains(accepts, "html") {
		redirectURL := fmt.Sprintf("/helloworld/404.html?url=%s%s%s", r.URL.Scheme, r.Host, r.URL.Path)
		http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
	} else if strings.Contains(accepts, "json") {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		http.Error(w, "{\"error\": \"Not found\"}", http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func main() {
	fmt.Println("start a simple web server...")
	http.HandleFunc("/helloworld/go", helloWorld)
	http.HandleFunc("/", notFound)
	http.ListenAndServe(":8080", nil)
}
