package aws

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"log"
)

// S3CreateBucketAPI defines the interface for the CreateBucket function.
// We use this interface to test the function using a mocked service.
type S3CreateBucketAPI interface {
	CreateBucket(ctx context.Context,
		params *s3.CreateBucketInput,
		optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
}

// S3CreateBucketAPI defines the interface for the CreateBucket function.
// We use this interface to test the function using a mocked service.
type S3PutBucketWebsiteAPI interface {
	PutBucketWebsite(ctx context.Context,
		params *s3.PutBucketWebsiteInput,
		optFns ...func(*s3.Options)) (*s3.PutBucketWebsiteOutput, error)
}

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

func MakeBucket(c context.Context, bucketname string) (string, error) {
	if bucketname == "" {
		fmt.Println("You must supply a bucket name.")
		return "", errors.New("empty  bucket name")
	}
	input := &s3.CreateBucketInput{
		Bucket: &bucketname,
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintEuCentral1,
		} ,
	}

	client, err := NewS3Client()
	if err != nil {
		log.Println("Could not create s3 client")
		log.Fatal(err)
		return "", errors.New("Could not connect to aws s3")
	}

	_, err = createBucket(c, client, input)
	if err != nil {
		log.Println("Could not create bucket " + bucketname)
		log.Fatal(err)
		return "", errors.New("Could not create s3 bucket")
	}

	webinput := &s3.PutBucketWebsiteInput{
		Bucket: &bucketname,
		WebsiteConfiguration: &types.WebsiteConfiguration{
			ErrorDocument: &types.ErrorDocument{
				Key: aws.String("error.html"),
			},
			IndexDocument: &types.IndexDocument{
				Suffix: aws.String("index.html"),
			},
		},
	}

	_, err = putBucketConfg(c, bucketname, client, webinput)
	if err != nil {
		fmt.Println("bucket " + bucketname + " updated with website configuration")
		fmt.Println(err)
		return "", errors.New("Could not update s3 bucket")
	}

}

// createBucket creates an Amazon Simple Storage Service (Amazon S3) bucket.
// Inputs:
//     c is the context of the method call, which includes the AWS Region
//     api is the interface that defines the method call
//     input defines the input arguments to the service call.
// Output:
//     If success, a CreateBucketOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to CreateBucket.
func createBucket(c context.Context, api S3CreateBucketAPI, input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	return api.CreateBucket(c, input)
}

// putBucketConfig creates an Amazon Simple Storage Service (Amazon S3) bucket configuration.
// Inputs:
//     c is the context of the method call, which includes the AWS Region
//     bucketname is the s3 resource name configuration is to be applied
// Output:
//     If success, a CreateBucketOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to CreateBucket.
func putBucketConfg(c context.Context, bucketname string, api S3PutBucketWebsiteAPI, input *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	return api.PutBucketWebsite(c, input)
}
