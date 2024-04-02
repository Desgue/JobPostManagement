package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port    string
	ConnStr string
	s3Cfg   S3Config
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	Port = os.Getenv("PORT")
	ConnStr = os.Getenv("CONN_STR")
	s3Cfg = S3Config{
		Url:       os.Getenv("S3_URL"),
		Region:    os.Getenv("S3_REGION"),
		AccessKey: os.Getenv("S3_ACCESS_KEY"),
		SecretKey: os.Getenv("S3_SECRET_KEY"),
		Token:     os.Getenv("S3_TOKEN"),
	}

}

func main() {
	postgresStore, err := NewPostgresStore(ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer postgresStore.Close()

	companyStore := NewCompanyStore(postgresStore)
	jobStore := NewJobStore(postgresStore)

	s3Client, err := NewS3Client(s3Cfg)
	if err != nil {
		log.Fatal(err)
	}

	companyService := NewCompanyService(companyStore)
	jobService := NewJobService(jobStore, s3Client)

	companiesController := NewCompanyController(companyService)
	jobsController := NewJobController(jobService)

	server := NewAPIServer(Port, companiesController, jobsController)
	log.Fatal(server.Start())
}
