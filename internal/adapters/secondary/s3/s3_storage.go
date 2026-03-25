package s3

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
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
	_, err := s.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}
	return s.cdnURL + key, nil
}

func (s *s3Storage) Ping() error {
	_, err := s.client.HeadBucket(context.Background(), &s3.HeadBucketInput{
		Bucket: aws.String(s.bucket),
	})
	return err
}
