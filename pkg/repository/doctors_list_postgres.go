package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	medapp "github.com/mnogohoddovochka/med-app"
)

type DoctorsListPostgres struct {
	db *sqlx.DB
}

func NewDoctorsListPostgres(db *sqlx.DB) *DoctorsListPostgres {
	return &DoctorsListPostgres{db: db}
}

func (r *DoctorsListPostgres) GetAll() ([]medapp.Doctor, error) {
	var doctors []medapp.Doctor

	query := fmt.Sprintf("SELECT id, name, surname, login FROM %s", doctorsTable)
	err := r.db.Select(&doctors, query)

	return doctors, err
}
