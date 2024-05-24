package main

import (
	"fmt"
	"net/http"
	"Converter/handlers"
)

const filesDirectory = "./Files"

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		handlers.UploadHandler(w, r, filesDirectory)
	})	
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Front/index.html")
	})

	fmt.Println("Server started on :8080")
	
	http.ListenAndServe(":8080", nil)
}
