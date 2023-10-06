package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv"
)

type Book struct {
	ID int
	Title string
	Stock int
	Author string
}

var books = []Book{
	{ID: 1, Title: "Buku 1", Stock: 20, Author: "Author1"},
	{ID: 2, Title: "Buku 2", Stock: 30, Author: "Author2"},
}


func main() {
   http.HandleFunc("books", getBooks)
  //  http.HandleFunc("add-book", createBook)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(books)
		return
	}
	http.Error(w, "invalid method", http.StatusBadRequest)
}

// func createBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
	
// 	if r.Method == "POST" {
// 		title := r.FormValue("title")
// 		stock := r.FormValue("stock")
// 		author := r.FormValue("author")

// 		convertSTock, err := strconv.Atoi(stock)
// 		if err !=

// 	}
// }