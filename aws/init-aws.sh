#!/usr/bin/env bash

awslocal s3api create-bucket --bucket reward-bucket

# public read policy
awslocal s3api put-bucket-policy \
  --bucket reward-bucket \
  --policy file:///etc/localstack/init/ready.d/bucket_policy.json

# static website (for CDN-like URL access)
awslocal s3 website s3://reward-bucket/ \
  --index-document index.html \
  --error-document error.html

# CORS
awslocal s3api put-bucket-cors \
  --bucket reward-bucket \
  --cors-configuration file:///etc/localstack/init/ready.d/cors-config.json

echo "✅ LocalStack S3 bucket 'reward-bucket' ready"
