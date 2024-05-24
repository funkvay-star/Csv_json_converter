package handlers

import (
	"fmt"
	"net/http"
	"Csv_json_converter/fileops"
)

func UploadHandler(w http.ResponseWriter, r *http.Request, directoryPath string) {
	if r.Method != http.MethodPost {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	fg := fileops.NewFileGetter()
	file, header, err := fg.GetFile(r)
	if err != nil {
		fmt.Println("Error getting file from request:", err)
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to parse form")
		return
	}
	defer file.Close()

	fu := fileops.NewFileUploader(directoryPath)
	if err := fu.SaveFile(file, header); err != nil {
		fmt.Println("Error saving file:", err)
		WriteErrorResponse(w, http.StatusInternalServerError, "Error saving file")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"message":"File successfully uploaded"}`)
}
