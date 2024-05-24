package logs

import (
	"log"
	"os"
	"io"
)

// Initializes the logging configuration.
func Init(logFile string) error {
	// Create or open the log file
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Log output to both file and stdout
	multi := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multi)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return nil
}

// Logs an informational message.
func Info(message string) {
	log.Println("INFO: " + message)
}

// Logs an error message.
func Error(err error) {
	log.Println("ERROR: " + err.Error())
}
