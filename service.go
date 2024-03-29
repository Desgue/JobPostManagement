package main

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
}

func NewJobService(jobStore *JobStore) *JobService {
	return &JobService{JobStore: jobStore}
}

func (s *JobService) GetFeed() ([]Job, error) {
	return s.JobStore.GetFeed()
}

func (s *JobService) GetJob(id string) (Job, error) {
	return s.JobStore.GetJob(id)
}

func (s *JobService) CreateJob(job Job) error {
	return s.JobStore.CreateJob(job)
}

func (s *JobService) UpdateJob(id string, job Job) error {
	return s.JobStore.UpdateJob(id, job)
}

func (s *JobService) PublishJob(id string) error {
	return s.JobStore.PublishJob(id)
}

func (s *JobService) ArchiveJob(id string) error {
	return s.JobStore.ArchiveJob(id)
}

func (s *JobService) DeleteJob(id string) error {
	return s.JobStore.DeleteJob(id)
}
