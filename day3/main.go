package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Struct(Model)
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Isbn   string  `json:"isbn"`
	Author *Author `json:"author"`
}

//Author Struct(Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

//GetBooks
func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Request getBooks with no parameter")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	log.Println("Response getBooks", books)
}

//Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Create Single Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

//Update Single Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			book.ID = item.ID
			books = append(books, book)
			break
		}

	}
	json.NewEncoder(w).Encode(book)
}

//Delete Single Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock book data
	books = append(books, Book{ID: "1", Title: "A Lost Love", Isbn: "nep-010203", Author: &Author{Firstname: "Love", Lastname: "Joshi"}})
	books = append(books, Book{ID: "2", Title: "A Lost Love2", Isbn: "nep-010204", Author: &Author{Firstname: "Love", Lastname: "Joshi"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
