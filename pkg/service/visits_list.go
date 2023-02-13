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
