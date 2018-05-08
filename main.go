package main

import (
	"github.com/gorilla/mux"
	"log"
	//"math/rand"
	//"strconv"
	"net/http"
	//"encoding/json"
)

//Book Struct (Model)
type Book struct {
	ID 		string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"` //this have it's own struct
}

//Author Struct
type Author struct {
	firstName		string `json:firstname"`
	lastName		string `json:firstname"`
}

// Init books var as a slice Book struct
var books []BookArr

// Get All Books //
func getBooks(w http.ResponseWriter, r *http.Request) { //Any function as a route handler must have these 2 properties

}

// Get Single Book //
func getBook(w http.ResponseWriter, r *http.Request) { 
	log.Println("Single book request")

}

// Create a new Book
func createBook(w http.ResponseWriter, r *http.Request) { 

}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) { 

}

// Delete single book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}


func main() {
	
	//Start Router
	r := mux.NewRouter()
	log.Println("Server Online P:8000")

	//Mock data @todo - implement DB
	books = append()books, Book{
		ID: 	 "1",
		Isbn:  "448743",
		Title: "Book One",
		Autor: &Author{firstName: "John", lastName: "Doe"}})

		books = append()books, Book{
			ID: 	 "2",
			Isbn:  "328743",
			Title: "Book Two",
			Autor: &Author{firstName: "Carl", lastName: "Johnson"}})

	
	//Router Handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books{id}", getBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}