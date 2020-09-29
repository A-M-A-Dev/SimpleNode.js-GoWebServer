package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// .....
	fmt.Fprintf(w, "hello world from Go")
}

func main() {
	fmt.Println("start a simple web server...")
	http.HandleFunc("/hw", helloWorld)
	http.ListenAndServe(":8080", nil)
}
