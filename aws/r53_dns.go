package aws

import (
	"context"
	"log"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	r53 "github.com/aws/aws-sdk-go-v2/service/route53"
	r53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
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

func MakeRoutes(c context.Context, s3websiteendpoint, dnsname, zoneid string) (string, error) {
	
	input := &r53.ChangeResourceRecordSetsInput{
		ChangeBatch: &r53types.ChangeBatch{
			Changes: []r53types.Change{
				{
					Action: r53types.ChangeActionCreate,
					ResourceRecordSet: &r53types.ResourceRecordSet{
						Name: aws.String(dnsname),
						AliasTarget: &r53types.AliasTarget{
							DNSName: &s3websiteendpoint,
							HostedZoneId: hostedZoneIdByS3EndpointRegion("eu-central-1"),
						},
						Type: r53types.RRTypeA,
					},
					
				},
			},
			Comment: aws.String("Web server for example.com"),
		},
		HostedZoneId: aws.String(zoneid),
	}

	client, err := NewR53Client()
	if err != nil {
		log.Println("Could not create R53 client")
		log.Fatal(err)
		return "", errors.New("Could not connect to aws R53")
	}

	_, err = createRoute(c, client, input, zoneid)
	if err != nil {
		log.Println("Could not create record ")
		log.Fatal(err)
		return "", errors.New("Could not create R53 record")
	}

	return dnsname, nil
}

func hostedZoneIdByS3EndpointRegion(region string) *string {
	zoneid := string(zonemap[region])
	return &zoneid
}

func createRoute(c context.Context, api R53ChangeResourceRecordSetsAPI, input *r53.ChangeResourceRecordSetsInput, hostedzone string) (*r53.ChangeResourceRecordSetsOutput, error) {
	return api.ChangeResourceRecordSets(c, input )
}
