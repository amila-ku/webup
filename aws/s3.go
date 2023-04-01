package aws

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"log"
	"os"
	"path/filepath"
)

// S3BucketAPI defines the interface for the S3 functionaliteis with Create Bucket, PutBucketWebiste, PutObject functions.
// We use this interface to test the function using a mocked service.
type S3BucketAPI interface {
	CreateBucket(ctx context.Context,
		params *s3.CreateBucketInput,
		optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
	PutBucketWebsite(ctx context.Context,
		params *s3.PutBucketWebsiteInput,
		optFns ...func(*s3.Options)) (*s3.PutBucketWebsiteOutput, error)
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

// S3PutObjectAPI defines the interface for the PutObject function.
// We use this interface to test the function using a mocked service.
// type S3PutObjectAPI interface {
// 	PutObject(ctx context.Context,
// 		params *s3.PutObjectInput,
// 		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
// }

// // S3PutBucketWebsiteAPI defines the interface for the CreateBucket function.
// // We use this interface to test the function using a mocked service.
// type S3PutBucketWebsiteAPI interface {
// 	PutBucketWebsite(ctx context.Context,
// 		params *s3.PutBucketWebsiteInput,
// 		optFns ...func(*s3.Options)) (*s3.PutBucketWebsiteOutput, error)
// }

// NewS3Client initializes a new aws s3 client.
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

// MakeBucket is used to create an s3 bucket with website config
// input: website name
func MakeBucket(c context.Context, client S3BucketAPI, bucketname string) (string, error) {
	if bucketname == "" {
		log.Println("You must supply a bucket name.")
		return "", errors.New("empty  bucket name")
	}
	input := &s3.CreateBucketInput{
		Bucket: &bucketname,
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintEuCentral1,
		},
	}

	_, err := createBucket(c, client, input)
	if err != nil {
		// log.Println("Could not create bucket " + bucketname)
		// log.Fatal(err)
		return "", err
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
		log.Println("bucket " + bucketname + " updated with website configuration")
		log.Println(err)
		return "", err
	}

	return bucketname, nil

}

func UploadFile(c context.Context, client S3BucketAPI, filename, bucketname string) error {
	if bucketname == "" || filename == "" {
		log.Println("You must supply a bucket name (-b BUCKET) and file name (-f FILE)")
		return errors.New("websitename or filename not suplied")
	}

	file, err := os.Open(filename)
	log.Printf("Opened file: %s", filename)

	if err != nil {
		log.Println("Unable to open file " + filename)
		return err
	}

	defer file.Close()

	filename = "index.html"

	// set ACL and other parameters according to
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/s3#PutObjectInput
	// more on Canned ACL : https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#CannedACL
	input := &s3.PutObjectInput{
		Bucket: &bucketname,
		Key:    &filename,
		Body:   file,
		ACL:    types.ObjectCannedACLPublicRead,
	}

	log.Printf("Trying to upload file: %s to s3", filename)

	st, err := putFile(c, client, input)
	if err != nil {
		log.Fatalln("Unable to upload file " + filename)
		return err
	}
	log.Printf("Uploaded file: %s to s3, object ETag: %v\n", filename, st.ETag)

	return nil
}

// fileWalk is used for identifying files in a given path excluding directories
type fileWalk chan string

func (f fileWalk) Walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		f <- path
	}
	return nil
}

// UploadFolder implemnted using https://aws.github.io/aws-sdk-go-v2/docs/sdk-utilities/s3/
func UploadFolder(c context.Context, client S3BucketAPI, localPath, bucketname string) error {
	if bucketname == "" || localPath == "" {
		return errors.New("Bucket name (-b BUCKET) and file name (-f FILE) is required")
	}

	walker := make(fileWalk)
	go func() {
		// Gather the files to upload by walking the path recursively
		if err := filepath.Walk(localPath, walker.Walk); err != nil {
			log.Fatalln("Walk failed:", err)
		}
		close(walker)
	}()

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalln("error:", err)
	}

	// prefix, err := os.Getwd()
	// if err != nil {
	// 	log.Println("Failed getting current path", err)
	// 	prefix = ""

	// }
	prefix := ""

	// For each file found walking, upload it to Amazon S3
	uploader := manager.NewUploader(s3.NewFromConfig(cfg))
	for path := range walker {
		rel, err := filepath.Rel(localPath, path)
		if err != nil {
			log.Fatalln("Unable to get relative path:", path, err)
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			log.Println("Failed opening file", path, err)
			continue
		}
		defer file.Close()
		result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: &bucketname,
			Key:    aws.String(filepath.Join(prefix, rel)),
			Body:   file,
			// more on Canned ACL : https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#CannedACL
			ACL: types.ObjectCannedACLPublicRead,
		})
		if err != nil {
			log.Fatalln("Failed to upload", path, err)
			continue
		}
		log.Println("Uploaded", path, result.Location)
	}

	return nil

}

// createBucket creates an Amazon Simple Storage Service (Amazon S3) bucket.
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region
//	api is the interface that defines the method call
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a CreateBucketOutput object containing the result of the service call and nil.
//	Otherwise, nil and an error from the call to CreateBucket.
func createBucket(c context.Context, api S3BucketAPI, input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	return api.CreateBucket(c, input)
}

// putBucketConfig creates an Amazon Simple Storage Service (Amazon S3) bucket configuration.
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region
//	bucketname is the s3 resource name configuration is to be applied
//
// Output:
//
//	If success, a CreateBucketOutput object containing the result of the service call and nil.
//	Otherwise, nil and an error from the call to CreateBucket.
func putBucketConfg(c context.Context, bucketname string, api S3BucketAPI, input *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	return api.PutBucketWebsite(c, input)
}

// putFile uploads a file to an Amazon Simple Storage Service (Amazon S3) bucket
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region
//	api is the interface that defines the method call
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a PutObjectOutput object containing the result of the service call and nil
//	Otherwise, nil and an error from the call to PutObject
func putFile(c context.Context, api S3BucketAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}
