package aws

import (
	"fmt"
	"context"
	"log"
	r53 "github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/config"
)

// CreateHostedZoneAPI defines the interface for the CreateHostedZone function.
// We use this interface to test the function using a mocked service.
type R53CreateHostedZoneAPI interface {
	CreateHostedZone(ctx context.Context,
		params *r53.CreateHostedZoneInput,
		optFns ...func(*r53.Options)) (*r53.CreateHostedZoneOutput, error)
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

func MakeRoutes() error {
	
}

func createRoute(c context.Context, api R53CreateHostedZoneAPI, input *r53.CreateHostedZoneInput) (*r53.CreateHostedZoneOutput, error) {
	return api.CreateHostedZone(c, input )
}
