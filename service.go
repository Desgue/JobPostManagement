package main

import (
	"database/sql"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

// CompanyService handles the business logic for the company domain
type CompanyService struct {
	CompanyStore *CompanyStore
}

func NewCompanyService(companyStore *CompanyStore) *CompanyService {
	return &CompanyService{CompanyStore: companyStore}
}

func (s *CompanyService) GetCompanies() ([]Company, error) {
	return s.CompanyStore.GetCompanies()
}

func (s *CompanyService) GetCompany(id string) (Company, error) {
	return s.CompanyStore.GetCompany(id)
}

// JobService handles the business logic for the job domain
type JobService struct {
	JobStore *JobStore
	s3Client *s3.Client
}

func NewJobService(jobStore *JobStore, s3Client *s3.Client) *JobService {
	return &JobService{JobStore: jobStore, s3Client: s3Client}
}

func (s *JobService) GetFeed() ([]Job, error) {
	/* 	var feed []Job
	   	output, err := ReadFileFromBucket("JobPostings", s.s3Client, "published_jobs.json")
	   	if err != nil {
	   		return nil, err
	   	}
	   	if err := json.Unmarshal(output, &feed); err != nil {
	   		return nil, err
	   	} */
	return s.JobStore.GetFeed()
}

func (s *JobService) CreateJob(job *JobRequest) error {
	job.Status = JobStatusDraft
	if err := job.Validate(); err != nil {
		return err
	}
	fmt.Println(job)

	return s.JobStore.CreateJob(job)
}

func (s *JobService) UpdateJob(id string, job JobRequest) error {
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid job id")
	}
	if err := job.Validate(); err != nil {
		return err
	}

	return s.JobStore.UpdateJob(id, job)
}

func (s *JobService) PublishJob(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid job id")
	}
	job, err := s.JobStore.GetJob(id)
	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("job not found")
	}
	if err != nil {
		return err
	}
	if job.Status != JobStatusDraft {
		return fmt.Errorf("job status must be draft to be published")
	}

	return s.JobStore.PublishJob(id)
}

func (s *JobService) ArchiveJob(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid job id")
	}
	job, err := s.JobStore.GetJob(id)
	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("job not found")
	}
	if err != nil {
		return err
	}
	if job.Status != JobStatusPublished {
		return fmt.Errorf("job status must be published to be archived")
	}

	return s.JobStore.ArchiveJob(id)
}

func (s *JobService) DeleteJob(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid job id")
	}
	_, err := s.JobStore.GetJob(id)
	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("job not found")
	}
	if err != nil {
		return err
	}

	return s.JobStore.DeleteJob(id)
}
