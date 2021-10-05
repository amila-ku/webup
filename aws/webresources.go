package aws

import (
	"context"
	"log"
	"errors"
)

func MakeWebResources(c context.Context, webSiteName string) error {

	// creates s3 bucket and setup bucket for webhosting
	bucket, err := MakeBucket(context.TODO(), webSiteName)
	if err != nil {
		log.Println("Error setting up s3 bucket")
		log.Fatal(err)
		return errors.New("Could not create s3 bucket")
	}
	log.Println("Bucket created" + bucket)

	// create route53 rules

}