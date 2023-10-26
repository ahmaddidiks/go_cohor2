package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Coundn't get file from form data", http.StatusBadRequest)
			return
		}
		defer file.Close()

		path := "/uploads"
		perm := os.ModePerm
		os.MkdirAll(path, perm)
		filePath := filepath.Join(path, handler.Filename)

		// open new file store from uploaded file
		newFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Coundn't Create File", http.StatusBadRequest)
			return
		}
		defer newFile.Close()

		// copy uploaded file to new file
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, "Coundn't Copy Uploaded File to New File", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "File Upload Successfull")

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server started at :2020")
	http.ListenAndServe(":2020", nil)
}
