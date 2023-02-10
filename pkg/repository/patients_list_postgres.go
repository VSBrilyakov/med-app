package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	medapp "github.com/mnogohoddovochka/med-app"
)

type PatientsListPostgres struct {
	db *sqlx.DB
}

func NewPatientsListPostgres(db *sqlx.DB) *PatientsListPostgres {
	return &PatientsListPostgres{db: db}
}

func (r *PatientsListPostgres) CreatePatient(input medapp.Patient) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, birthdate) VALUES ($1, $2, $3) RETURNING id",
		patientsTable)

	row := r.db.QueryRow(query, input.Name, input.Surname, input.BirthDate)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
