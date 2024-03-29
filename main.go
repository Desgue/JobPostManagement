package main

func main() {
	companiesController := &CompanyController{}
	jobsController := &JobController{}
	server := NewAPIServer(":8080", companiesController, jobsController)
	server.Start()
}
