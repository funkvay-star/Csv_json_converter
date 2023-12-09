package filegetter

// IFileGetter represents an interface for file operations
type IFileGetter interface {
    GetFile() error
    GetFileName() (string, error)
    RetrieveFile() ([]byte, error)
    RemoveFile() error
    CleanupCache()
}
