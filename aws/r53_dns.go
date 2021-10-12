package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	r53 "github.com/aws/aws-sdk-go-v2/service/route53"
	r53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/aws-sdk-go/service/route53"
)

// CreateHostedZoneAPI defines the interface for the CreateHostedZone function.
// We use this interface to test the function using a mocked service.
type R53CreateHostedZoneAPI interface {
	CreateHostedZone(ctx context.Context,
		params *r53.CreateHostedZoneInput,
		optFns ...func(*r53.Options)) (*r53.CreateHostedZoneOutput, error)
}

// ChangeResourceRecordSetsAPI defines the interface for the CreateHostedZone function.
// We use this interface to test the function using a mocked service.
type R53ChangeResourceRecordSetsAPI interface {
	ChangeResourceRecordSets(ctx context.Context,
		params *r53.ChangeResourceRecordSetsInput,
		optFns ...func(*r53.Options)) (*r53.ChangeResourceRecordSetsOutput, error)
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
	
	input := &r53.ChangeResourceRecordSetsInput{
		ChangeBatch: &r53types.ChangeBatch{
			Changes: []*r53types.Change{
				{
					Action: aws.String("CREATE"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String("example.com"),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.44"),
							},
						},
						TTL:  aws.Int64(60),
						Type: aws.String("A"),
					},
				},
			},
			Comment: aws.String("Web server for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}
}

func createRoute(c context.Context, api R53ChangeResourceRecordSetsAPI, input *r53.ChangeResourceRecordSetsInput, hostedzone string) (*r53.ChangeResourceRecordSetsOutput, error) {
	return api.ChangeResourceRecordSets(c, input )
}
