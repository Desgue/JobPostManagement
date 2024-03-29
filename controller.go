package main

import (
	"fmt"
	"net/http"
)

type Controller interface {
	RegisterRoutes(mux *http.ServeMux)
}

type CompanyController struct {
	CompanyService *CompanyService
	JobService     *JobService
}

func (h *CompanyController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /companies", h.GetCompanies)
	mux.HandleFunc("GET /company/{id}", h.GetCompany)
}

func (c CompanyController) GetCompanies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting companies"))
}

func (c CompanyController) GetCompany(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte(fmt.Sprintf("Getting company %s", id)))
}

type JobController struct {
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
	w.Write([]byte("Posting job draft"))
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
