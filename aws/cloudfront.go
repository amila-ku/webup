package aws

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cf "github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cftypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

// CF3API defines the interface for the CloudFront function.
// We use this interface to test the function using a mocked service.
type CFAPI interface {
	CreateDistributionWithTags(ctx context.Context, params *cf.CreateDistributionWithTagsInput, optFns ...func(*cf.Options)) (*cf.CreateDistributionWithTagsOutput, error)
}

// NewCFClient initializes a new aws CloudFront client.
func NewCFClient() (*cf.Client, error) {
	ctx := context.TODO()
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Create an Amazon Cloudfront service client
	client := cf.NewFromConfig(cfg)

	return client, err
}


func createDistribution(c context.Context, api CFAPI, input *cf.CreateDistributionWithTagsInput, hostedzone string) (*cf.CreateDistributionWithTagsOutput, error) {
	return api.CreateDistributionWithTags(c, input)
}