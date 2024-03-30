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
	db *PostgresStore
}

func NewCompanyStore(db *PostgresStore) *CompanyStore {
	return &CompanyStore{db: db}
}

func (s *CompanyStore) GetCompanies() ([]Company, error) {
	return []Company{}, nil
}

func (s *CompanyStore) GetCompany(id string) (Company, error) {
	return Company{}, nil
}

// JobStore is responsible for interacting with the postgress database and handling the job data
type JobStore struct {
	db *PostgresStore
}

func NewJobStore(db *PostgresStore) *JobStore {
	return &JobStore{db: db}
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
