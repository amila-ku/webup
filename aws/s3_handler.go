package aws

import (
	"context"
	"fmt"
	//"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3CreateBucketAPI defines the interface for the CreateBucket function.
// We use this interface to test the function using a mocked service.
type S3CreateBucketAPI interface {
	CreateBucket(ctx context.Context,
		params *s3.CreateBucketInput,
		optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
}



func MakeBucket(c context.Context, bucketname string) {
	if bucketname == "" {
		fmt.Println("You must supply a bucket name.")
		return
	}
	input := &s3.CreateBucketInput{
		Bucket: &bucketname,
	}

	client, err := NewS3Client() 
	if err != nil {
		fmt.Println("Could not create s3 client")
		fmt.Println(err)
	}

	_, err = makeBucket(c, client, input)
	if err != nil {
		fmt.Println("Could not create bucket " + bucketname)
		fmt.Println(err)
	}


}

// makeBucket creates an Amazon Simple Storage Service (Amazon S3) bucket.
// Inputs:
//     c is the context of the method call, which includes the AWS Region
//     api is the interface that defines the method call
//     input defines the input arguments to the service call.
// Output:
//     If success, a CreateBucketOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to CreateBucket.
func makeBucket(c context.Context, api S3CreateBucketAPI, input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	return api.CreateBucket(c, input)
}

// putBucketConfig creates an Amazon Simple Storage Service (Amazon S3) bucket configuration.
// Inputs:
//     c is the context of the method call, which includes the AWS Region
//     bucketname is the s3 resource name configuration is to be applied
// Output:
//     If success, a CreateBucketOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to CreateBucket.
func putBucketConfg(c context.Context, bucketname string) (*s3.PutBucketWebsiteOutput, error) {
	config := &s3.PutBucketWebsiteInput{
		Bucket: &bucketname,
		
	}
}