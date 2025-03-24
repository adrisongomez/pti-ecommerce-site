package utils

import (
	"context"
	"fmt"

	internalUtils "github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetResourceURL(bucket, region, key string) string {
	if internalUtils.IsProduction() {
		return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, key)
	}
	return fmt.Sprintf("http://%s.s3.%s.localhost:4566/%s", bucket, region, key)
}

func CreateObjectOnS3(ctx context.Context, bucket, key string, size int64) (string, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("dummy", "dummy", "dummy")),
		config.WithBaseEndpoint("http://localhost:4566"),
	)

	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	presignClient := s3.NewPresignClient(client)

	presignedReq, err := presignClient.PresignPutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			// ACL:    types.ObjectCannedACLPublicRead,
		},
	)

	if err != nil {
		return "", err
	}

	return presignedReq.URL, nil
}
