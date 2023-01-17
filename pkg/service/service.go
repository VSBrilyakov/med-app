package service

import (
	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/mnogohoddovochka/med-app/pkg/repository"
)

type Authorisation interface {
	CreateDoctor(doctor medapp.Doctor) (int, error)
}

type DoctorList interface {
}

type PatientList interface {
}

type Service struct {
	Authorisation
	DoctorList
	PatientList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repo.Authorisation),
	}
}
