package filegetter

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type LocalGetter struct {
    cachedFilePath string
}

// NewLocalGetter initializes and returns a new LocalGetter instance
func NewLocalGetter() *LocalGetter {
    return &LocalGetter{}
}

// GetFile fetches a file from a specified directory and caches its path.
func (lg *LocalGetter) GetFile() error {
    directoryPath := "../BackendAndFrontend/Files"

    files, err := ioutil.ReadDir(directoryPath)
    if err != nil {
        return err
    }

    for _, file := range files {
        if !file.IsDir() {
            lg.cachedFilePath = filepath.Join(directoryPath, file.Name())
            return nil
        }
    }
    return fmt.Errorf("no regular files found in directory: %s", directoryPath)
}

// GetFileName extracts and returns the name of the cached file (without the extension).
func (lg *LocalGetter) GetFileName() (string, error) {
    if lg.cachedFilePath == "" {
        return "", fmt.Errorf("file path is empty")
    }
    return strings.TrimSuffix(filepath.Base(lg.cachedFilePath), filepath.Ext(lg.cachedFilePath)), nil
}

// RetrieveFile returns the content of the cached file.
func (lg *LocalGetter) RetrieveFile() ([]byte, error) {
    return ioutil.ReadFile(lg.cachedFilePath)
}

// RemoveFile removes the cached file from the filesystem.
func (lg *LocalGetter) RemoveFile() error {
    if err := os.Remove(lg.cachedFilePath); err != nil {
        return err
    }
    lg.cachedFilePath = ""
    return nil
}

// CleanupCache cleans up the cached file path.
func (lg *LocalGetter) CleanupCache() {
    lg.cachedFilePath = ""
}
