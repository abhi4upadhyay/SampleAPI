package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//LineA
	r := mux.NewRouter().StrictSlash(true)

	//Line B
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello World, This is a sample API listening on Port :8081")
	}).Methods(http.MethodGet)

	//Line C
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8081", r))
}