package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrMissingTitle       = errors.New("missing job title")
	ErrMissingDescription = errors.New("missing job description")
	ErrMissingLocation    = errors.New("missing job location")
	ErrInvalidStatus      = errors.New("invalid job status")
	ErrInvalidCompanyID   = errors.New("invalid company id")
)

type Company struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type JobStatus string

func (s JobStatus) IsValid() bool {
	switch s {
	case JobStatusDraft, JobStatusPublished, JobStatusArchived, JobStatusRejected:
		return true
	default:
		return false
	}
}

const (
	JobStatusDraft     JobStatus = "draft"
	JobStatusPublished JobStatus = "published"
	JobStatusArchived  JobStatus = "archived"
	JobStatusRejected  JobStatus = "rejected"
)

type Job struct {
	Id          string    `json:"id"`
	CompanyId   string    `json:"companyId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Notes       string    `json:"notes"`
	Status      JobStatus `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type JobRequest struct {
	Title       string    `json:"title"`
	CompanyId   string    `json:"companyId"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Notes       string    `json:"notes"`
	Status      JobStatus `json:"status"`
}

func (j *JobRequest) Validate() error {
	if _, err := uuid.Parse(j.CompanyId); err != nil {
		return ErrInvalidCompanyID
	}
	if j.Title == "" {
		return ErrMissingTitle
	}
	if j.Description == "" {
		return ErrMissingDescription
	}
	if j.Location == "" {
		return ErrMissingLocation
	}
	if !j.Status.IsValid() {
		return ErrInvalidStatus
	}
	return nil
}
