package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Rating float32 `json:"rating"`
}

// To store books
var books []Book

func main() {
	books = append(books,
		Book{Id: "0", Name: "Jurassic Park", Author: " Michael Chrichton", Rating: 4.1},
		Book{Id: "1", Name: "The Martian", Author: "Andy Weir", Rating: 4.3})
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/books", getAllBooks).Methods("GET")
	router.HandleFunc("/api/v1/books", addBook).Methods("POST")
	router.HandleFunc("/api/v1/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/v1/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		fmt.Println("error while encoding...", err)
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newBook := Book{}
	err1 := json.NewDecoder(r.Body).Decode(&newBook)
	if err1 != nil {
		fmt.Println("error while decoding...", err1)
	}
	books = append(books, newBook)
	err2 := json.NewEncoder(w).Encode(newBook)
	if err2 != nil {
		fmt.Println("error while encoding...", err2)
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	bookIdToBeUpdated := params["id"]

	bookToBeUpdated := Book{}
	err1 := json.NewDecoder(r.Body).Decode(&bookToBeUpdated)
	if err1 != nil {
		fmt.Println("error while decoding...", err1)
	}

	for index, book := range books {
		if book.Id == bookIdToBeUpdated {
			books[index] = bookToBeUpdated
			break
		}
	}

	err2 := json.NewEncoder(w).Encode(books)
	if err2 != nil {
		fmt.Println("error while encoding...", err2)
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	bookIdToBeDeleted := params["id"]

	for index, book := range books {
		if book.Id == bookIdToBeDeleted {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	err2 := json.NewEncoder(w).Encode(books)
	if err2 != nil {
		fmt.Println("error while encoding...", err2)
	}
}
