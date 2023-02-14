package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/sirupsen/logrus"
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

func (r *PatientsListPostgres) GetAll() ([]medapp.Patient, error) {
	var patients []medapp.Patient

	query := fmt.Sprintf("SELECT id, name, surname, birthdate FROM %s", patientsTable)
	err := r.db.Select(&patients, query)

	return patients, err
}

func (r *PatientsListPostgres) GetById(id int) (medapp.Patient, error) {
	var patient medapp.Patient

	query := fmt.Sprintf("SELECT id, name, surname, birthdate FROM %s WHERE id = $1", patientsTable)
	err := r.db.Get(&patient, query, id)

	return patient, err
}

func (r *PatientsListPostgres) UpdatePatient(id int, input medapp.UpdPatient) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *input.Surname)
		argId++
	}

	if input.BirthDate != nil {
		setValues = append(setValues, fmt.Sprintf("birthdate=$%d", argId))
		args = append(args, *input.BirthDate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d", patientsTable, setQuery, argId)
	args = append(args, id)

	logrus.Debug("updateQuery: %s", query)
	logrus.Debug("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
