package aws

import (
	"context"
	"testing"


	"github.com/stretchr/testify/assert"
)

func TestMakeWebResources(t *testing.T) {

	// creates a s3 client
	s3client := &S3BucketImpl{}

	// creates s3 bucket and setup bucket for webhosting
	bucket, err := MakeBucket(context.TODO(), s3client, "abc.com")
	assert.Nil(t, err)
	assert.Equal(t, "abc.com", bucket)

	// creates a r53 client
	r53client := &R53APIImpl{}

	// create route53 rules
	dns, err := MakeRoutes(context.TODO(), r53client, "abc-test.com", "Z1TI4H711TUAOG")
	assert.Nil(t, err)
	assert.Equal(t, "abc-test.com", dns)
}
