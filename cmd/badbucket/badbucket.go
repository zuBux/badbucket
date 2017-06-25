package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/zuBux/badbucket"
)

var bucket = flag.String("bucket", "", "S3 bucket name")
var regionPtr string

func init() {
	// example with short version for long flag
	flag.StringVar(bucket, "b", "", "S3 bucket name")
	flag.StringVar(&regionPtr, "region", "", "AWS region")
	flag.StringVar(&regionPtr, "r", "", "AWS region")
}

func main() {
	flag.Parse()
	// Initialize a session that the SDK will use to load configuration,
	// credentials, and region from the shared config file. (~/.aws/config).
	svc := createS3Session(regionPtr)
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(regionPtr)})
	list, err := badbucket.IsBucketListable(svc, *bucket)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("Listable: %t\n", list)

	read, err := badbucket.AreFilesReadable(sess, *bucket)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("Readable: %t\n", read)

	write, err := badbucket.IsBucketWriteable(sess, *bucket)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("Writeable: %t\n", write)
}

func createS3Session(region string) *s3.S3 {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		log.Panicf("Could not create session: %s", err)
	}
	svc := s3.New(sess)
	return svc
}
