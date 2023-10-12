package handlers

import (
	"fmt"
	"net/http"
	"Converter/fileops"
)

func UploadHandler(w http.ResponseWriter, r *http.Request, directoryPath string) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    fg := fileops.NewFileGetter()
    file, header, err := fg.GetFile(r)
    if err != nil {
        fmt.Println("Error getting file from request:", err)
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }
    defer file.Close()

    fu := fileops.NewFileUploader(directoryPath)
    if err := fu.SaveFile(file, header); err != nil {
        fmt.Println("Error saving file:", err)
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `{"message":"File successfully uploaded"}`)
}
