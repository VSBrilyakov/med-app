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
	query := fmt.Sprintf("INSERT INTO %s (name, surname, login, password_hash) VALUES ($1, $2, $3, $4) RETURNING id",
		doctorsTable)

	row := r.db.QueryRow(query, doctor.Name, doctor.Surname, doctor.Login, doctor.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetDoctor(login, password string) (medapp.Doctor, error) {
	var doctor medapp.Doctor
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", doctorsTable)
	err := r.db.Get(&doctor, query, login, password)

	return doctor, err
}
