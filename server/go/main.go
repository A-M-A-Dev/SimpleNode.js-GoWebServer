package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	redirectURL := fmt.Sprintf("/helloworld/404.html?url=%s%s%s", r.URL.Scheme, r.Host, r.URL.Path)
	http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
}

func main() {
	fmt.Println("start a simple web server...")
	http.HandleFunc("/helloworld/go", helloWorld)
	http.HandleFunc("/", notFound)
	http.ListenAndServe(":8080", nil)
}
