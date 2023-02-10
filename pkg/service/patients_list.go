package service

import (
	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/mnogohoddovochka/med-app/pkg/repository"
)

type PatientsListService struct {
	repo repository.PatientList
}

func NewPatientsListService(repo repository.PatientList) *PatientsListService {
	return &PatientsListService{repo: repo}
}

func (s *PatientsListService) CreatePatient(input medapp.Patient) (int, error) {
	return s.repo.CreatePatient(input)
}

func (s *PatientsListService) GetAll() ([]medapp.Patient, error) {
	return s.repo.GetAll()
}

func (s *PatientsListService) GetById(id int) (medapp.Patient, error) {
	return s.repo.GetById(id)
}
