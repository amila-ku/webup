package aws

import (
	"context"
	"fmt"
	"testing"

	r53 "github.com/aws/aws-sdk-go-v2/service/route53"
	r53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/stretchr/testify/assert"
)

// R53APIImpl is for implementing testable r53 client without calling aws services
type R53APIImpl struct{}

func (dt R53APIImpl) ChangeResourceRecordSets(ctx context.Context,
	params *r53.ChangeResourceRecordSetsInput,
	optFns ...func(*r53.Options)) (*r53.ChangeResourceRecordSetsOutput, error) {

	output := &r53.ChangeResourceRecordSetsOutput{
		ChangeInfo: &r53types.ChangeInfo{},
		ResultMetadata: middleware.Metadata{},
	}

	return output, nil
	//return output, errors.New("random error")
}

func TestMakeRoutes(t *testing.T) {
	api := &R53APIImpl{}
	_, err := MakeRoutes(context.TODO(), api, "s3-website-eu-west-1.amazonaws.com", "abc-test.com", "Z1TI4H711TUAOG")
	if err != nil {
		fmt.Println("failed")
	}
	assert.Nil(t, err)
}
