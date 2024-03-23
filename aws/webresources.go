package aws

import (
	"context"
	"errors"
	"log"
	"github.com/aws/aws-sdk-go-v2/config"
)

// MakeWebResources is used to create an s3 bucket with website config and add required dns entries in Route53
// input: website name, route 53 hosted zone id
func NewWebResources(c context.Context, webSiteName string, route53HostedZoneID string, skipBucketCreation bool) error {

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// creates s3 bucket if skip bucket creation is not true
	if !skipBucketCreation {
		// creates a s3 client
		s3client, err := NewS3Client()
		if err != nil {
			log.Println("Could not create s3 client")
			log.Fatal(err)
			return errors.New("Could not connect to aws s3")
		}

		// creates s3 bucket and setup bucket for webhosting
		bucket, err := MakeBucket(c, s3client, webSiteName, cfg.Region)
		if err != nil {
			log.Println("Error setting up s3 bucket")
			log.Fatal(err)
			return errors.New("Could not create s3 bucket")
		}
		log.Printf("Bucket %s created \n", bucket)
	}

	// creates a R53 client
	r53client, err := NewR53Client()
	if err != nil {
		log.Println("Could not create R53 client")
		log.Fatal(err)
		return errors.New("Could not connect to aws R53")
	}

	// create route53 rules
	// example hosted zone value "Z1TI4H711TUGO5"
	dns, err := MakeRoutes(c, r53client, webSiteName, route53HostedZoneID, cfg.Region)
	if err != nil {
		log.Println("Error setting up R53 entries for " + webSiteName)
		log.Fatal(err)
		return errors.New("Could not create s3 bucket")
	}
	log.Printf("DNS %s created ", dns)

	return nil
}

// UploadContent is used to upload a file to a s3 bucket with website config.
// input: website name
func UploadContent(c context.Context, webSiteName string) error {

	// creates a s3 client
	s3client, err := NewS3Client()
	if err != nil {
		log.Println("Could not create s3 client")
		log.Fatal(err)
		return errors.New("Could not connect to aws s3")
	}

	err = UploadFile(c, s3client, "webcontent/index.html", webSiteName)
	if err != nil {
		log.Fatal(err)
		return errors.New("Could not upload to aws s3")
	}

	return nil
}

// UploadContentFolder is used to upload multiple files to a s3 bucket with website config.
// input: website name
func UploadContentFolder(c context.Context, uploadPath, webSiteName string) error {

	// creates a s3 client
	s3client, err := NewS3Client()
	if err != nil {
		return errors.New("could not create  aws s3 client")
	}

	err = UploadFolder(c, s3client, uploadPath, webSiteName)
	if err != nil {
		log.Fatal(err)
		return errors.New("Could not upload folder content to aws s3")
	}

	return nil
}
