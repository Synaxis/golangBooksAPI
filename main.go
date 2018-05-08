package main

import (
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"strconv"
	"net/http"
	"encoding/json"
)

//Book Struct (Model)
type Book Struct {
	ID 		string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`

}

func main() {
	
	//Start Router

	r := mux.NewRouter()
	
	//Router Handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books{id}", getBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}