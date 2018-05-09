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
type Book struct {
	ID 		string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"` //this have it's own struct
}

//Author Struct
type Author struct {
	Firstname		string `json:"firstname"`
	Lastname		string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books //
func getAllBooks(w http.ResponseWriter, r *http.Request) { //Any function as a route handler must have these 2 properties
	log.Println("All books request")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

// Get Single Book //
func getSingleBook(w http.ResponseWriter, r *http.Request) { 
	log.Println("Single books request")
	// set params
	params := mux.Vars(r) // Get params
	//Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new Book
func createBook(w http.ResponseWriter, r *http.Request) { 
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //generate a random ID
	//Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) { 

}

// Delete single book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {		
			books = append(books[:index], books[index+1:]...)// todo check this
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main() {
	
	//Start Router
	r := mux.NewRouter()
	log.Println("Server Online P:8000")

	//Mock data @todo - implement DB
	books = append(books, Book{
		ID: 	 "1",
		Isbn:  "448743",
		Title: "Book One",
		Author: &Author{Firstname: "John", Lastname: "Doe"},
	})

		books = append(books, Book{
			ID: 	 "2",
			Isbn:  "328743",
			Title: "Book Two",
			Author: &Author{Firstname: "Carl", Lastname: "Johnson"},
		})

	
	//Router Handlers
	r.HandleFunc("/api/books", getAllBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getSingleBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}