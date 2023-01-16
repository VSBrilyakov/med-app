package repository

import "github.com/jmoiron/sqlx"

type Authorisation interface {
}

type DoctorList interface {
}

type PatientList interface {
}

type Repository struct {
	Authorisation
	DoctorList
	PatientList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
