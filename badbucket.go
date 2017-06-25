package badbucket

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//IsBucketListable checks if a bucket is listable
func IsBucketListable(s *s3.S3, bucketName string) (bool, error) {
	params := &s3.ListObjectsInput{Bucket: aws.String(bucketName)}
	_, err := s.ListObjects(params)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Get error details
			if awsErr.Code() == "NoSuchBucket" {
				return false, err
			} else {
				return false, nil
			}
		}
	}
	return true, nil
}

//AreFilesReadable lists objects in an s3 bucket and tries to download the first one
func AreFilesReadable(s *session.Session, bucketName string) (bool, error) {
	svc := s3.New(s)
	downloader := s3manager.NewDownloader(s)
	params := &s3.ListObjectsInput{Bucket: aws.String(bucketName)}
	resp, err := svc.ListObjects(params)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Get error details
			if awsErr.Code() == "NoSuchBucket" {
				return false, err
			} else {
				return false, nil
			}
		}
	}
	if resp.Contents == nil {
		return false, fmt.Errorf("No files in bucket")
	}
	testKey := resp.Contents[0].Key
	file, err := os.Create(*testKey)

	if err != nil {
		log.Panicf("Unable to open file %q, %v", "testfile", err)
	}

	defer file.Close()
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(*testKey),
		})

	if err != nil {
		log.Printf("Unable to download item %q, %v", *testKey, err)
		return false, nil
	}

	return true, nil
}

func IsBucketWriteable(s *session.Session, bucketName string) (bool, error) {
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(s)
	d1 := []byte("thiswasgeneratedbybadbycket\n")
	err := ioutil.WriteFile("/tmp/buckettest.txt", d1, 0644)

	f, err := os.Open("/tmp/buckettest.txt")
	defer f.Close()
	if err != nil {
		return false, fmt.Errorf("failed to open file %q, %v", f, err)
	}

	// Upload the file to S3.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("badbucketpoc.txt"),
		Body:   f,
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Get error details
			if awsErr.Code() == "AccessDenied" {
				return false, nil
			} else {
				return false, err
			}
		}
	}
	return true, nil
}
