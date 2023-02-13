package medapp

type Visit struct {
	DocLogin         string `json:"login" binding:"required"`
	PatientName      string `json:"name" binding:"required"`
	PatientSurname   string `json:"surname" binding:"required"`
	PatientBirthdate string `json:"birthdate" binding:"required"`
	VisitDate        string `json:"visitdate" binding:"required" time_format:"2006-01-02"`
}
