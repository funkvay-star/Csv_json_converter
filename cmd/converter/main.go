package main

import (
    "log"
    "net/http"
    "os"
    "Csv_json_converter/config"
    "Csv_json_converter/handlers"
    "Csv_json_converter/internal/utils"
    "Csv_json_converter/logs"
)

func main() {
    cfg := config.LoadConfig()

    // Initialize logging
    err := logs.Init("app.log")
    if err != nil {
        log.Fatalf("Could not initialize log: %v", err)
    }

    // Check if upload directory exists
    err = utils.EnsureDir(cfg.UploadDirectory)
    if err != nil {
        logs.Error(err)
        os.Exit(1)
    }

    http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
        handlers.UploadHandler(w, r, cfg.UploadDirectory)
    })

    // Serve static files from the Front directory
    http.Handle("/", http.FileServer(http.Dir("./Front")))

    logs.Info("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        logs.Error(err)
    }
}
