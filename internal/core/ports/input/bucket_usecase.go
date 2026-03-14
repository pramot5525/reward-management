package input

type BucketUsecase interface {
	UploadImage(filename string, data []byte, contentType string) (string, error)
	PingStorage() error
}
