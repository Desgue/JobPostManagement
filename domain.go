package main

import "time"

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
	Id          string
	CompanyId   string
	Title       string
	Description string
	Location    string
	Notes       string
	Status      JobStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
