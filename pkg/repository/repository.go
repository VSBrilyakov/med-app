package repository

type Authorisation interface {
}

type DoctorList interface {
}

type PatientList interface {
}

type Repository struct {
	Authorisation
	DoctorList
	PatientList
}

func NewRepository() *Repository {
	return &Repository{}
}
