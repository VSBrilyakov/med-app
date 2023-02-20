package medapp

import "errors"

type Visit struct {
	DocLogin         string `json:"login" binding:"required" example:"DrHouse"`
	PatientName      string `json:"name" binding:"required" example:"Igor"`
	PatientSurname   string `json:"surname" binding:"required" example:"Vasilev"`
	PatientBirthdate string `json:"birthdate" binding:"required" example:"2000-05-13"`
	VisitDate        string `json:"visitdate" binding:"required" time_format:"2006-01-02" example:"2022-12-26"`
}

type VisitOutput struct {
	DocName        string `json:"docname" binding:"required" example:"Gregory"`
	DocSurname     string `json:"docsurname" binding:"required" example:"House"`
	PatientName    string `json:"patientname" binding:"required" example:"Igor"`
	PatientSurname string `json:"patientsurname" binding:"required" example:"Vasiliev"`
	Date           string `json:"date" binding:"required" example:"2022-12-26"`
}

type UpdVisit struct {
	DocLogin         *string `json:"login" example:"DrBykov"`
	PatientName      *string `json:"name" example:"Mihail"`
	PatientSurname   *string `json:"surname" example:"Kravcov"`
	PatientBirthdate *string `json:"birthdate" example:"2006-01-25"`
	VisitDate        *string `json:"visitdate" time_format:"2006-01-02" example:"2021-12-26"`
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
