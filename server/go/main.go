package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

type AdderData struct {
	Sha256 string `json:"sha256"`
}

type ErrorData struct {
	Message string `json:"message"`
}

type AdderRequestBody struct {
	A int
	B int
}

func writeJsonError(errorData ErrorData, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorData)
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
		writeJsonError(ErrorData{"Sorry, only POST method is supported."}, w)
		return
	} else if err := r.ParseForm(); err != nil {
		writeJsonError(ErrorData{"Error Parsing POST Data"}, w)
		return
	}
	var adderReqBody AdderRequestBody
	err := json.NewDecoder(r.Body).Decode(&adderReqBody)
	if err != nil {
		writeJsonError(ErrorData{"You should send only number as parameters"}, w)
		return
	}

	hashAdder := sha256.Sum256([]byte(strconv.Itoa(adderReqBody.A + adderReqBody.B)))
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
