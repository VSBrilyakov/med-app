package service

import (
	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/mnogohoddovochka/med-app/pkg/repository"
)

type VisitsListPostgres struct {
	repo repository.VisitList
}

func NewVisitsListPostgres(repo repository.VisitList) *VisitsListPostgres {
	return &VisitsListPostgres{repo: repo}
}

func (s *VisitsListPostgres) CreateVisit(visit medapp.Visit) (int, error) {
	return s.repo.CreateVisit(visit)
}

func (s *VisitsListPostgres) GetAll() ([]medapp.VisitOutput, error) {
	return s.repo.GetAll()
}

func (s *VisitsListPostgres) GetById(id int) (medapp.VisitOutput, error) {
	return s.repo.GetById(id)
}

func (s *VisitsListPostgres) UpdateVisit(id int, input medapp.UpdVisit) error {
	return s.repo.UpdateVisit(id, input)
}

func (s *VisitsListPostgres) DeleteVisit(id int) error {
	return s.repo.DeleteVisit(id)
}
