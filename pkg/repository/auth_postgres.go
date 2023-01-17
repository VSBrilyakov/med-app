package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	medapp "github.com/mnogohoddovochka/med-app"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateDoctor(doctor medapp.Doctor) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, password_hash) VALUES ($1, $2, $3) RETURNING id",
		doctorsTable)

	row := r.db.QueryRow(query, doctor.Name, doctor.Surname, doctor.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return 0, nil
}
