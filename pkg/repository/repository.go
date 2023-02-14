package repository

import (
	"github.com/jmoiron/sqlx"
	medapp "github.com/mnogohoddovochka/med-app"
)

type Authorisation interface {
	CreateDoctor(doctor medapp.Doctor) (int, error)
	GetDoctor(login, password string) (medapp.Doctor, error)
}

type DoctorList interface {
	GetAll() ([]medapp.Doctor, error)
	GetById(id int) (medapp.Doctor, error)
}

type PatientList interface {
	CreatePatient(input medapp.Patient) (int, error)
	GetAll() ([]medapp.Patient, error)
	GetById(id int) (medapp.Patient, error)
	UpdatePatient(id int, input medapp.UpdPatient) error
}

type VisitList interface {
	CreateVisit(input medapp.Visit) (int, error)
	GetAll() ([]medapp.VisitOutput, error)
	GetById(id int) (medapp.VisitOutput, error)
	UpdateVisit(id int, input medapp.UpdVisit) error
}

type Repository struct {
	Authorisation
	DoctorList
	PatientList
	VisitList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
		DoctorList:    NewDoctorsListPostgres(db),
		PatientList:   NewPatientsListPostgres(db),
		VisitList:     NewVisitsListPostgres(db),
	}
}
