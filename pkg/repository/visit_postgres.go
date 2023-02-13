package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	medapp "github.com/mnogohoddovochka/med-app"
)

type VisitsListPostgres struct {
	db *sqlx.DB
}

func NewVisitsListPostgres(db *sqlx.DB) *VisitsListPostgres {
	return &VisitsListPostgres{db: db}
}

func (r *VisitsListPostgres) CreateVisit(input medapp.Visit) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var docId int
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1", doctorsTable)
	row := tx.QueryRow(query, input.DocLogin)
	if err := row.Scan(&docId); err != nil {
		tx.Rollback()
		return 0, err
	}

	var patId int
	query = fmt.Sprintf("SELECT id FROM %s WHERE name=$1 AND surname=$2 AND birthdate=$3", patientsTable)
	row = tx.QueryRow(query, input.PatientName, input.PatientSurname, input.PatientBirthdate)
	if err := row.Scan(&patId); err != nil {
		tx.Rollback()
		return 0, err
	}

	var id int
	query = fmt.Sprintf("INSERT INTO %s (docId, patientId, date) VALUES ($1, $2, $3) RETURNING id",
		visitsTable)
	row = tx.QueryRow(query, docId, patId, input.VisitDate)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *VisitsListPostgres) GetAll() ([]medapp.VisitOutput, error) {
	var visits []medapp.VisitOutput

	query := fmt.Sprintf("SELECT doctors.name AS docname, doctors.surname AS docsurname, patients.name AS patientname, patients.surname AS patientsurname, visits.date FROM %s LEFT JOIN %s ON visits.docId = doctors.id LEFT JOIN %s ON visits.patientId = patients.id",
		visitsTable, doctorsTable, patientsTable)
	err := r.db.Select(&visits, query)

	return visits, err
}

func (r *VisitsListPostgres) GetById(id int) (medapp.VisitOutput, error) {
	var visit medapp.VisitOutput

	query := fmt.Sprintf("SELECT doctors.name AS docname, doctors.surname AS docsurname, patients.name AS patientname, patients.surname AS patientsurname, visits.date FROM %s LEFT JOIN %s ON visits.docId = doctors.id LEFT JOIN %s ON visits.patientId = patients.id WHERE visits.id = $1",
		visitsTable, doctorsTable, patientsTable)
	err := r.db.Get(&visit, query, id)

	return visit, err
}
