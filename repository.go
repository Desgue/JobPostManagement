package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(conn string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	postgresStore := &PostgresStore{db: db}
	postgresStore.mustPing()
	return postgresStore, nil
}

func (s *PostgresStore) Close() error {
	return s.db.Close()
}

func (s *PostgresStore) mustPing() {
	if err := s.db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to the database")
}

// CompanyStore is responsible for interacting with the postgress database and handling the company data
type CompanyStore struct {
	postgres *PostgresStore
}

func NewCompanyStore(db *PostgresStore) *CompanyStore {
	return &CompanyStore{postgres: db}
}

func (s *CompanyStore) GetCompanies() ([]Company, error) {
	var companies []Company
	rows, err := s.postgres.db.Query("SELECT * FROM companies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var company Company
		rows.Scan(&company.Id, &company.Name, &company.CreatedAt, &company.UpdatedAt)
		companies = append(companies, company)
	}
	return companies, nil
}

func (s *CompanyStore) GetCompany(id string) (Company, error) {
	var company Company
	err := s.postgres.db.QueryRow("SELECT * FROM companies WHERE id = $1", id).Scan(&company.Id, &company.Name, &company.CreatedAt, &company.UpdatedAt)
	if err != nil {
		return Company{}, err
	}
	return company, nil
}

// JobStore is responsible for interacting with the postgress database and handling the job data
type JobStore struct {
	postgres *PostgresStore
}

func NewJobStore(db *PostgresStore) *JobStore {
	return &JobStore{postgres: db}
}

// GetFeed queries a S3 bucket and return a json file with all the jobs that are published
func (s *JobStore) GetFeed() ([]Job, error) {
	var jobs []Job
	rows, err := s.postgres.db.Query("SELECT * FROM jobs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var job Job
		rows.Scan(&job.Id, &job.CompanyId, &job.Title, &job.Description, &job.Location, &job.Notes, &job.Status, &job.CreatedAt, &job.UpdatedAt)
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (s *JobStore) GetJob(id string) (Job, error) {
	var job Job
	err := s.postgres.db.QueryRow("SELECT * FROM jobs WHERE id = $1", id).Scan(&job.Id,
		&job.CompanyId,
		&job.Title,
		&job.Description,
		&job.Location,
		&job.Notes,
		&job.Status,
		&job.CreatedAt,
		&job.UpdatedAt,
	)
	if err != nil {
		return Job{}, err
	}
	return job, nil
}

// All the other methods will query the database directly

func (s *JobStore) CreateJob(job *JobRequest) error {
	_, err := s.postgres.db.Exec(
		"INSERT INTO jobs (company_id, title, description, location,notes, status) VALUES ($1, $2, $3, $4, $5, $6)",
		job.CompanyId,
		job.Title,
		job.Description,
		job.Location,
		job.Notes,
		job.Status,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *JobStore) UpdateJob(id string, job JobRequest) error {
	return nil
}

func (s *JobStore) PublishJob(id string) error {
	_, err := s.postgres.db.Exec("UPDATE jobs SET status = 'published', updated_at = now() WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *JobStore) ArchiveJob(id string) error {
	_, err := s.postgres.db.Exec("UPDATE jobs SET status = 'archived', updated_at = now() WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *JobStore) DeleteJob(id string) error {
	_, err := s.postgres.db.Exec("DELETE FROM jobs WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
