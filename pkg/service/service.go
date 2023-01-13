package service

import "github.com/mnogohoddovochka/med-app/pkg/repository"

type Authorisation interface {
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
	return &Service{}
}
