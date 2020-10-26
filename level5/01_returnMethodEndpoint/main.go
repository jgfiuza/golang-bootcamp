package main

import (
	"log"
	"net/http"
)

func returnMethod(w http.ResponseWriter, r *http.Request) {
	var message []byte

	switch r.Method {
	case "GET":
		message = []byte(`{"message": "get called"}`)
	case "POST":
		message = []byte(`{"message": "post called"}`)
	default:
		message = []byte(`{"message": "unknown method called"}`)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

func handleRequests() {
	http.HandleFunc("/", returnMethod)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequests()
}
