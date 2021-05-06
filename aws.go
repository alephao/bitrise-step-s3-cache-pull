package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsS3 struct {
	sess       *session.Session
	bucketName string
}

func NewAwsS3(awsRegion, accessKey, secretKey, bucket string) *AwsS3 {
	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")
	aws.NewConfig()
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: creds,
	}))
	return &AwsS3{
		sess:       sess,
		bucketName: bucket,
	}
}

func (a *AwsS3) CacheExists(key string) (bool, string) {
	svc := s3.New(a.sess)
	obj, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket:  aws.String(a.bucketName),
		MaxKeys: aws.Int64(1),
		Prefix:  aws.String(key),
	})

	if err != nil {
		log.Printf("An error occurred when hitting the cache: %s. Assuming there is no cache\n", err.Error())
		return false, ""
	}

	if len(obj.Contents) > 0 {
		return true, *obj.Contents[0].Key
	} else {
		return false, ""
	}
}

func (a *AwsS3) Download(key, outputPath string) (int64, error) {
	downloader := s3manager.NewDownloader(a.sess)

	downloadedFile, err := os.Create(outputPath)
	if err != nil {
		return 0, err
	}
	defer downloadedFile.Close()

	return downloader.Download(
		downloadedFile,
		&s3.GetObjectInput{
			Bucket: aws.String(a.bucketName),
			Key:    aws.String(key),
		},
	)
}

func (a *AwsS3) UploadToAws(key, pathToFile string) error {
	uploader := s3manager.NewUploader(a.sess)

	f, err := os.Open(pathToFile)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", pathToFile, err)
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(a.bucketName),
		Key:    aws.String(key),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	return nil
}
