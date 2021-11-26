package aws

import (
	"context"
	"errors"
	"log"
)

// MakeWebResources is used to create an s3 bucket with website config and add required dns entries in Route53
// input: website name
func MakeWebResources(c context.Context, webSiteName string) error {

	// creates a s3 client
	s3client, err := NewS3Client()
	if err != nil {
		log.Println("Could not create s3 client")
		log.Fatal(err)
		return errors.New("Could not connect to aws s3")
	}

	// creates s3 bucket and setup bucket for webhosting
	bucket, err := MakeBucket(c, s3client, webSiteName)
	if err != nil {
		log.Println("Error setting up s3 bucket")
		log.Fatal(err)
		return errors.New("Could not create s3 bucket")
	}
	log.Println("Bucket created" + bucket)

	// creates a R53 client
	r53client, err := NewR53Client()
	if err != nil {
		log.Println("Could not create R53 client")
		log.Fatal(err)
		return errors.New("Could not connect to aws R53")
	}

	// create route53 rules
	dns, err := MakeRoutes(c, r53client, "s3-website-eu-west-1.amazonaws.com", webSiteName, "Z1TI4H711TUGO5")
	if err != nil {
		log.Println("Error setting up s3 bucket")
		log.Fatal(err)
		return errors.New("Could not create s3 bucket")
	}
	log.Println("DNS created" + dns)

	return nil
}
