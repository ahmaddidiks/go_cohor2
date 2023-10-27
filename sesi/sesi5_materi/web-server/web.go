package main

import (
	"fmt"
	"net/http"
)

var PORT = ":9090"

func main() {
	http.HandleFunc("/", greet)

	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Helloooo!"
	fmt.Fprintf(w, msg)
}
