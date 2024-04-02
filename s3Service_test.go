package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/localstack"
)

var (
	cont       *container
	ctx        context.Context
	bucketName = "test"
	opt        S3Config
)

func TestMain(m *testing.M) {
	var err error
	ctx = context.Background()
	cont, err = newTestContainer(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer cont.localstack.Terminate(ctx)
	opt = S3Config{
		Url:       cont.url,
		Region:    "us-east-1",
		AccessKey: "test",
		SecretKey: "test",
		Token:     "test",
	}

	client, err := NewS3Client(opt)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}
func TestNewS3Client(t *testing.T) {
	client, err := NewS3Client(opt)
	if err != nil {
		t.Fatal(err)
	}
	output, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		t.Fatal(err)
	}
	if len(output.Buckets) != 1 {
		t.Fatalf("expected 0 buckets, got %d", len(output.Buckets))
	}
	if *output.Buckets[0].Name != "test" {
		t.Fatalf("expected bucket name to be %s, got %s", bucketName, *output.Buckets[0].Name)
	}
}

func TestAddFileToBucket(t *testing.T) {
	client, err := NewS3Client(opt)
	if err != nil {
		t.Fatal(err)
	}
	job := Job{
		Id:          "f6bd9c06-9079-461b-8fd5-a5b868f3377d",
		Title:       "Software Engineer",
		CompanyId:   "f6bd9c06-9079-461b-8fd5-a5b868f3377d",
		Description: "Software Engineer at Google",
		Location:    "Mountain View, CA",
		Notes:       "Must have 5 years of experience",
		Status:      JobStatusPublished,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	jobs := []Job{job}
	b, err := json.Marshal(jobs)
	if err != nil {
		t.Fatal(err)
	}
	err = AddFileToBucket(bucketName, client, "published_jobs.json", b)
	if err != nil {
		t.Fatal(err)
	}
	files, err := ListBucketFiles(bucketName, client)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}
	if files[0] != "published_jobs.json" {
		t.Fatalf("expected file name to be published_jobs.json, got %s", files[0])
	}

}

// The principle function to be tested, that will actually be implemented in the service, all the other functions were to setup the buckets to this test
func TestReadFileFromBucket(t *testing.T) {
	client, err := NewS3Client(opt)
	if err != nil {
		t.Fatal(err)
	}
	output, err := ReadFileFromBucket(bucketName, client, "published_jobs.json")
	if err != nil {
		t.Fatal(err)
	}

	var jobs []Job
	err = json.Unmarshal(output, &jobs)
	if err != nil {
		t.Fatal(err)
	}
	if len(jobs) != 1 {
		t.Fatalf("expected 1 job, got %d", len(jobs))
	}
	if jobs[0].Id != "f6bd9c06-9079-461b-8fd5-a5b868f3377d" {
		t.Fatalf("expected job id to be f6bd9c06-9079-461b-8fd5-a5b868f3377d, got %s", jobs[0].Id)
	}

}

// Test Setup

type container struct {
	localstack *localstack.LocalStackContainer
	url        string
}

func newTestContainer(ctx context.Context) (*container, error) {
	localstack, err := localstack.RunContainer(ctx, testcontainers.WithImage("localstack/localstack:latest"))
	if err != nil {
		return nil, err
	}
	provider, err := testcontainers.NewDockerProvider()
	if err != nil {
		return nil, err
	}
	host, err := provider.DaemonHost(ctx)
	if err != nil {
		return nil, err
	}
	port, err := localstack.MappedPort(ctx, nat.Port("4566/tcp"))
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("http://%s:%s", host, port.Port())
	return &container{localstack: localstack, url: url}, nil
}
