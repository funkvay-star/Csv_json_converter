package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"mime/multipart"
)

// FileGetter is responsible for getting files from the request
type FileGetter struct{}

// NewFileGetter creates a new FileGetter
func NewFileGetter() *FileGetter {
	return &FileGetter{}
}

// GetFile retrieves the file from the request
func (fg *FileGetter) GetFile(r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile("file")
}

// FileUploader is responsible for saving files to a directory
type FileUploader struct {
	Directory string
}

// NewFileUploader creates a new FileUploader
func NewFileUploader(directory string) *FileUploader {
	return &FileUploader{Directory: directory}
}

// SaveFile saves the file to the specified directory
func (fu *FileUploader) SaveFile(file multipart.File, header *multipart.FileHeader) error {
	dst, err := os.Create(fmt.Sprintf("%s/%s", fu.Directory, header.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fg := NewFileGetter()
	file, header, err := fg.GetFile(r)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fu := NewFileUploader("../Files")
	if err := fu.SaveFile(file, header); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"message":"File successfully uploaded"}`)
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../Front/index.html")
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
