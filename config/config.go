package config

import (
	"os"
	"log"
)

type Config struct {
	UploadDirectory string
}

func LoadConfig() *Config {
	uploadDirectory, exists := os.LookupEnv("UPLOAD_DIRECTORY")
	if !exists {
		uploadDirectory = "./uploads" // default upload directory
	}

	return &Config{
		UploadDirectory: uploadDirectory,
	}
}

func (c *Config) Validate() {
	if _, err := os.Stat(c.UploadDirectory); os.IsNotExist(err) {
		log.Printf("Upload directory does not exist, creating: %s", c.UploadDirectory)
		err := os.MkdirAll(c.UploadDirectory, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create upload directory: %v", err)
		}
	}
}
