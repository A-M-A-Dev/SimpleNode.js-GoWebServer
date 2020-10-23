package main

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"encoding/json"
	"crypto/sha256"
	"encoding/hex"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

type AdderData struct {
	Sha256 string `json:"sha256"`
}
type ErrorData struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func writeJsonError(errorData ErrorData, w http.ResponseWriter) {
	jsonData, err := json.Marshal(errorData)
	if err != nil {
		fmt.Fprintf(w, "Error : %s", err)
	}
	w.Write(jsonData)
}

func writeJsonAdderData(adderData AdderData, w http.ResponseWriter) {
	dataJson, err := json.Marshal(adderData)
	if err != nil {
		fmt.Fprintf(w, "Error : %s", err)
	}
	w.Write(dataJson)
}

func adder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		writeJsonError(ErrorData{true, "Sorry, only POST method is supported."} ,w)
		return
	} else if err := r.ParseForm(); err != nil {
		writeJsonError(ErrorData{true, "Error Parsing POST Data"} ,w)
		return
	}
	a, err1 := strconv.Atoi(r.FormValue("a"))
	b, err2 := strconv.Atoi(r.FormValue("b"))
	if err1 != nil || err2 != nil {
		writeJsonError(ErrorData{true, "You should send only number as parameters"} ,w)
		return
	}

	hashAdder := sha256.Sum256([]byte(strconv.Itoa(a + b)))
	hashAdderString := hex.EncodeToString(hashAdder[:])
	data := AdderData{hashAdderString}
	writeJsonAdderData(data, w)
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
	http.HandleFunc("/helloworld/go/adder", adder)
	http.HandleFunc("/", notFound)
	http.ListenAndServe(":8080", nil)
}
