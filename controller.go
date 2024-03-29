package main

import "net/http"

type Controller interface {
	RegisterRoutes(mux *http.ServeMux)
}

type CompanyController struct {
}

func (h *CompanyController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/companies", h.GetCompanies)
	mux.HandleFunc("/company/{id}", h.GetCompany)
}

func (c CompanyController) GetCompanies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting companies"))
}

func (c CompanyController) GetCompany(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting company"))
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
	w.Write([]byte("Updating job"))
}

func (c JobController) PublishJob(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Publishing job"))
}

func (c JobController) ArchiveJob(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Archiving job"))
}

func (c JobController) DeleteJobDraft(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a job draft"))
}

func (c JobController) GetFeed(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting feed"))
}
