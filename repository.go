package main

// CompanyStore is responsible for interacting with the postgress database and handling the company data
type CompanyStore struct {
	conn string
}

func NewCompanyStore(conn string) *CompanyStore {
	return &CompanyStore{conn: conn}
}

func (s *CompanyStore) GetCompanies() ([]Company, error) {
	return []Company{}, nil
}

func (s *CompanyStore) GetCompany(id string) (Company, error) {
	return Company{}, nil
}

// JobStore is responsible for interacting with the postgress database and handling the job data
type JobStore struct {
	conn string
}

func NewJobStore(conn string) *JobStore {
	return &JobStore{conn: conn}
}

// GetFeed queries a S3 bucket and return a json file with all the jobs that are published
func (s *JobStore) GetFeed() ([]Job, error) {
	return []Job{}, nil
}

// All the other methods will query the database directly
func (s *JobStore) GetJob(id string) (Job, error) {
	return Job{}, nil
}

func (s *JobStore) CreateJob(job Job) error {
	return nil
}

func (s *JobStore) UpdateJob(id string, job Job) error {
	return nil
}

func (s *JobStore) PublishJob(id string) error {
	return nil
}

func (s *JobStore) ArchiveJob(id string) error {
	return nil
}

func (s *JobStore) DeleteJob(id string) error {
	return nil
}
