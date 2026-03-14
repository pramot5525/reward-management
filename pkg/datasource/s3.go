package datasource

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewS3Client returns an S3 client.
// When localEndpoint is set (local env), it points to LocalStack instead of AWS.
func NewS3Client(localEndpoint string) (*s3.Client, error) {
	if localEndpoint != "" {
		cfg, err := config.LoadDefaultConfig(context.Background(),
			config.WithRegion("ap-southeast-1"),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
		)
		if err != nil {
			return nil, err
		}
		return s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(localEndpoint)
			o.UsePathStyle = true // required for LocalStack
		}), nil
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(cfg), nil
}
