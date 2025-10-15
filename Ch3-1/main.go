package main

// using a router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getRequest(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a GET\n")

}

func postRequest(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a POST\n")

}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a DELETE\n")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getRequest).Methods("GET")
	r.HandleFunc("/", postRequest).Methods("POST")
	r.HandleFunc("/", deleteRequest).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server started and listening on localhost:9003 ")
	log.Fatal(http.ListenAndServe(":9003", nil))
}