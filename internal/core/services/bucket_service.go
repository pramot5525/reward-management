package services

import "github.com/nocnoc-thailand/reward-management/internal/core/ports/output"

type bucketService struct {
	storage output.StoragePort
}

func NewBucketService(storage output.StoragePort) *bucketService {
	return &bucketService{storage: storage}
}

func (s *bucketService) UploadImage(filename string, data []byte, contentType string) (string, error) {
	return s.storage.Upload(filename, data, contentType)
}

func (s *bucketService) PingStorage() error {
	return s.storage.Ping()
}
