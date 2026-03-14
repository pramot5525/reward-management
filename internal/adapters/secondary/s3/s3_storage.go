package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Storage struct {
	client *s3.Client
	bucket string
	cdnURL string
}

func NewS3Storage(client *s3.Client, bucket, cdnURL string) *s3Storage {
	return &s3Storage{client: client, bucket: bucket, cdnURL: cdnURL}
}

func (s *s3Storage) Upload(key string, data []byte, contentType string) (string, error) {
	// TODO: implement PutObject
	_ = context.Background()
	return s.cdnURL + key, nil
}

func (s *s3Storage) Ping() error {
	// TODO: implement HeadBucket
	return nil
}
