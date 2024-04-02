package main

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	Url       string
	Region    string
	AccessKey string
	SecretKey string
	Token     string
}

func NewS3Client(
	opt S3Config,
) (*s3.Client, error) {

	customResolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, opts ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           opt.Url,
				SigningRegion: opt.Region,
			}, nil
		})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(opt.Region),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(opt.AccessKey, opt.SecretKey, opt.Token)),
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	return client, nil
}

func ReadFileFromBucket(bucketName string, client *s3.Client, fileName string) ([]byte, error) {
	output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &fileName,
	})
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(output.Body)
	return buf.Bytes(), nil
}

func ListBucketFiles(bucketName string, client *s3.Client) ([]string, error) {
	var files []string
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	})
	if err != nil {
		return nil, err
	}
	for _, obj := range output.Contents {
		files = append(files, *obj.Key)
	}
	return files, nil
}

func AddFileToBucket(bucketName string, client *s3.Client, fileName string, fileContent []byte) error {
	body := bytes.NewReader(fileContent)
	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &fileName,
		Body:   body,
	})
	return err
}
