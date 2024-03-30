package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port    string
	ConnStr string
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	Port = os.Getenv("PORT")
	ConnStr = os.Getenv("CONN_STR")

}

func main() {
	postgresStore, err := NewPostgresStore(ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer postgresStore.Close()

	companyStore := NewCompanyStore(postgresStore)
	jobStore := NewJobStore(postgresStore)

	companyService := NewCompanyService(companyStore)
	jobService := NewJobService(jobStore)

	companiesController := NewCompanyController(companyService)
	jobsController := NewJobController(jobService)

	server := NewAPIServer(Port, companiesController, jobsController)
	log.Fatal(server.Start())
}
