package service

import (
	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/mnogohoddovochka/med-app/pkg/repository"
)

type DoctorsListService struct {
	repo repository.DoctorList
}

func NewDoctorsListService(repo repository.DoctorList) *DoctorsListService {
	return &DoctorsListService{repo: repo}
}

func (s *DoctorsListService) GetAll() ([]medapp.Doctor, error) {
	return s.repo.GetAll()
}
