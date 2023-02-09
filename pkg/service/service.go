package service

import (
	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/mnogohoddovochka/med-app/pkg/repository"
)

type Authorisation interface {
	CreateDoctor(doctor medapp.Doctor) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type DoctorList interface {
	GetAll() ([]medapp.Doctor, error)
	GetById(id int) (medapp.Doctor, error)
}

type PatientList interface {
}

type VisitList interface {
}

type Service struct {
	Authorisation
	DoctorList
	PatientList
	VisitList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repo.Authorisation),
		DoctorList:    NewDoctorsListService(repo.DoctorList),
	}
}
