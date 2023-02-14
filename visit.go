package medapp

import "errors"

type Visit struct {
	DocLogin         string `json:"login" binding:"required"`
	PatientName      string `json:"name" binding:"required"`
	PatientSurname   string `json:"surname" binding:"required"`
	PatientBirthdate string `json:"birthdate" binding:"required"`
	VisitDate        string `json:"visitdate" binding:"required" time_format:"2006-01-02"`
}

type VisitOutput struct {
	DocName        string `json:"docname" binding:"required"`
	DocSurname     string `json:"docsurname" binding:"required"`
	PatientName    string `json:"patientname" binding:"required"`
	PatientSurname string `json:"patientsurname" binding:"required"`
	Date           string `json:"date" binding:"required"`
}

type UpdVisit struct {
	DocLogin         *string `json:"login"`
	PatientName      *string `json:"name"`
	PatientSurname   *string `json:"surname"`
	PatientBirthdate *string `json:"birthdate"`
	VisitDate        *string `json:"visitdate" time_format:"2006-01-02"`
}

func (u UpdVisit) Validate() error {
	if u.DocLogin == nil && u.PatientBirthdate == nil && u.PatientName == nil && u.PatientSurname == nil && u.VisitDate == nil {
		return errors.New("update structure has no values")
	}

	if (u.PatientName != nil && (u.PatientSurname == nil || u.PatientBirthdate == nil)) ||
		(u.PatientSurname != nil && (u.PatientName == nil || u.PatientBirthdate == nil)) ||
		(u.PatientBirthdate != nil && (u.PatientName == nil || u.PatientSurname == nil)) {
		return errors.New("all patient information fields must be filled")
	}

	return nil
}
