package repository

import (
	"github.com/jmoiron/sqlx"
	medapp "github.com/mnogohoddovochka/med-app"
)

type Authorisation interface {
	CreateDoctor(doctor medapp.Doctor) (int, error)
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
	return &Repository{
		Authorisation: NewAuthPostgres(db),
	}
}
