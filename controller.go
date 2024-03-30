package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller interface {
	RegisterRoutes(mux *http.ServeMux)
}

type CompanyController struct {
	service *CompanyService
}

func NewCompanyController(companyService *CompanyService) *CompanyController {
	return &CompanyController{service: companyService}
}
func (h *CompanyController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /companies", h.GetCompanies)
	mux.HandleFunc("GET /company/{id}", h.GetCompany)
}

func (c CompanyController) GetCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := c.service.GetCompanies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(companies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (c CompanyController) GetCompany(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	company, err := c.service.GetCompany(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

type JobController struct {
	JobService *JobService
}

func NewJobController(jobService *JobService) *JobController {
	return &JobController{JobService: jobService}
}

func (c *JobController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /feed", c.GetFeed)
	mux.HandleFunc("POST /job", c.PostJobDraft)
	mux.HandleFunc("PUT /job/{id}", c.UpdateJob)
	mux.HandleFunc("PUT /job/{id}/publish", c.PublishJob)
	mux.HandleFunc("PUT /job/{id}/archive", c.ArchiveJob)
	mux.HandleFunc("DELETE /job/{id}", c.DeleteJobDraft)

}

func (c JobController) PostJobDraft(w http.ResponseWriter, r *http.Request) {
	var job JobRequest
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.JobService.CreateJob(job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Job draft created"))
}

func (c JobController) UpdateJob(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte(fmt.Sprintf("Updating Job with ID %s", id)))
}

func (c JobController) PublishJob(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte(fmt.Sprintf("Publishing Job with ID %s", id)))
}

func (c JobController) ArchiveJob(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte(fmt.Sprintf("Archiving Job with ID %s", id)))
}

func (c JobController) DeleteJobDraft(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte(fmt.Sprintf("Deleting Job with ID %s", id)))
}

func (c JobController) GetFeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Getting feed"))

}
