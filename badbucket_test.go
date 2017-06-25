package badbucket

import (
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var TestBucket = "yourbucketname"

func TestIsBucketListable(t *testing.T) {
	log.Printf("Testing if bucket %s is listable", TestBucket)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		log.Panicf("Could not create session: %s", err)
	}
	svc := s3.New(sess)
	listable, _ := IsBucketListable(svc, TestBucket)
	if listable != true {
		t.Fail()
	}
}

func TestIsBucketListableNonExistent(t *testing.T) {
	log.Printf("Testing handling of non-existent buckets")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		log.Panicf("Could not create session: %s", err)
	}
	svc := s3.New(sess)
	_, err = IsBucketListable(svc, "thisbucketdoesnotexist123123123")
	if err == nil {
		t.Fail()
	}
}

func TestAreFilesReadable(t *testing.T) {
	var isDownloadable bool
	log.Printf("Testing readable bucket")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		log.Panicf("Could not create session: %s", err)
	}
	//svc := s3.New(sess)
	isDownloadable, _ = AreFilesReadable(sess, TestBucket)
	if isDownloadable != true {
		t.Fail()
	}

}

func TestDetectBucketRegion(t *testing.T) {
	log.Printf("Testing region detection")
	region := DetectBucketRegion("badbucket-test")
	if region != "eu-west-1" {
		t.Fail()
	}
}
