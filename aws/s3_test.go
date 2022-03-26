package aws

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/middleware"
	"github.com/stretchr/testify/assert"
)

// S3BucketImpl is for implementing testable s3 client without calling aws services
type S3BucketImpl struct{}

func (dt S3BucketImpl) CreateBucket(ctx context.Context,
	params *s3.CreateBucketInput,
	optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {

	output := &s3.CreateBucketOutput{
		Location: aws.String("us-west-2"),
	}

	return output, nil

}

func (dt S3BucketImpl) PutBucketWebsite(ctx context.Context,
	params *s3.PutBucketWebsiteInput,
	optFns ...func(*s3.Options)) (*s3.PutBucketWebsiteOutput, error) {

	output := &s3.PutBucketWebsiteOutput{
		ResultMetadata: middleware.Metadata{},
	}

	return output, nil

}

func TestMakeBucket(t *testing.T) {
	api := &S3BucketImpl{}

	_, err := MakeBucket(context.TODO(), api, "abc.com")

	assert.Nil(t, err)
}
