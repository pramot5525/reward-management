package output

type StoragePort interface {
	Upload(key string, data []byte, contentType string) (string, error)
	Ping() error
}
