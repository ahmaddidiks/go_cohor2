package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int
	Title  string
	Stock  int
	Author string
}

var books = []Book{
	{ID: 1, Title: "Love For Imperfect Things", Stock: 20, Author: "Haemin Sunim"},
	{ID: 2, Title: "The Power of Nunchi", Stock: 30, Author: "Euny Hong"},
	{ID: 3, Title: "Winter in Tokyo", Stock: 40, Author: "Ilana Tan"},
}

var PORT = ":9090"

func main() {
	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/book", createBook)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(books)
		return
	}

	// using html/template package
	// if r.Method == "GET" {
	// 	tpl, err := template.ParseFiles("template.html")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	tpl.Execute(w, books)
	// 	return
	// }

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		title := r.FormValue("title")
		stock := r.FormValue("stock")
		author := r.FormValue("author")

		convertStock, err := strconv.Atoi(stock)
		if err != nil {
			http.Error(w, "Invalid stock", http.StatusBadRequest)
		}

		newBook := Book{
			ID:     len(books) + 1,
			Title:  title,
			Stock:  convertStock,
			Author: author,
		}

		books = append(books, newBook)

		json.NewEncoder(w).Encode(books)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
