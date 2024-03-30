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
	companiesController := &CompanyController{}
	jobsController := &JobController{}
	server := NewAPIServer(Port, companiesController, jobsController)
	server.Start()
}
