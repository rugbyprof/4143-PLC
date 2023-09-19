package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", uploadForm).Methods("GET")
	r.HandleFunc("/upload", uploadImage).Methods("POST")

	http.Handle("/", r)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func uploadForm(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Image Downloader</title>
		</head>
		<body>
			<h1>Image Downloader</h1>
			<form action="/upload" method="post" enctype="multipart/form-data">
				URL to Image: <input type="text" name="image_url"><br>
				<input type="submit" value="Download Image">
			</form>
		</body>
		</html>
	`
	fmt.Fprintf(w, html)
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Print("in function")
	r.ParseForm()
	url := r.FormValue("image_url")

	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch the image: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Failed to fetch the image: HTTP status %d", resp.StatusCode), http.StatusInternalServerError)
		return
	}

	// Extract the file name from the URL
	fileName := url[strings.LastIndex(url, "/")+1:]

	file, err := os.Create(fileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save the image: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Image downloaded successfully as %s", fileName)
}
