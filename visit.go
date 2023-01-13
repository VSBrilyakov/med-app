package medapp

import "time"

type Visit struct {
	Id        int       `json:"-"`
	DoctorId  int       `json:"doctorId"`
	PatientId int       `json:"patientId"`
	Date      time.Time `json:"date"`
}
