package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	fmt.Println("start a simple web server...")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
