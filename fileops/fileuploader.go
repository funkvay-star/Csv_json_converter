package fileops

import (
	"fmt"
	"io"
	"os"
	"mime/multipart"
)

type FileUploader struct {
	Directory string
}

func NewFileUploader(directory string) *FileUploader {
	return &FileUploader{Directory: directory}
}

func (fu *FileUploader) SaveFile(file io.Reader, header *multipart.FileHeader) error {
	dst, err := os.Create(fmt.Sprintf("%s/%s", fu.Directory, header.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}
