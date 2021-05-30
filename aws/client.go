package aws

import (
	"log"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	r53 "github.com/aws/aws-sdk-go-v2/service/route53"
)

// type Client struct {
// 	config config.Config
// }

func NewS3Client() (*s3.Client, error) {
	ctx := context.TODO()
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	return client, err
}

func NewR53Client() (*r53.Client, error) {
	ctx := context.TODO()
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Create an Amazon S3 service client
	client := r53.NewFromConfig(cfg)

	return client, err
}