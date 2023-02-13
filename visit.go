package medapp

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
