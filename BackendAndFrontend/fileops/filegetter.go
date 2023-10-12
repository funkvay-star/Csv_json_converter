package fileops

import (
	"net/http"
	"mime/multipart"
)

type FileGetter struct{}

func NewFileGetter() *FileGetter {
	return &FileGetter{}
}

func (fg *FileGetter) GetFile(r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile("file")
}
