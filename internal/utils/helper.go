package utils

import (
	"log"
	"os"
)

// Ensures that a directory exists, and creates it if it doesn't.
func EnsureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// Logs the error if it's not nil.
func LogIfError(err error) {
	if err != nil {
		log.Println(err)
	}
}
