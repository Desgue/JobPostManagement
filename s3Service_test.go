package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/localstack"
)

func TestNewS3Client(t *testing.T) {
	var ctx = context.Background()
	container := newTestContainer(ctx, t)
	defer container.localstack.Terminate(ctx)

	client, err := NewS3Client(container.url, "us-east-1", "test", "test", "test")
	if err != nil {
		t.Fatal(err)
	}
	output, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		t.Fatal(err)
	}
	if len(output.Buckets) != 0 {
		t.Fatalf("expected 0 buckets, got %d", len(output.Buckets))
	}
}

// Test Setup

type container struct {
	localstack *localstack.LocalStackContainer
	url        string
}

func newTestContainer(ctx context.Context, t *testing.T) *container {
	localstack, err := localstack.RunContainer(ctx, testcontainers.WithImage("localstack/localstack:latest"))
	if err != nil {
		t.Fatal(err)
	}
	provider, err := testcontainers.NewDockerProvider()
	if err != nil {
		t.Fatal(err)
	}
	host, err := provider.DaemonHost(ctx)
	if err != nil {
		t.Fatal(err)
	}
	port, err := localstack.MappedPort(ctx, nat.Port("4566/tcp"))
	if err != nil {
		t.Fatal(err)
	}
	url := fmt.Sprintf("http://%s:%s", host, port.Port())
	return &container{localstack: localstack, url: url}
}
