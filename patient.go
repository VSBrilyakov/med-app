package medapp

import "errors"

type Patient struct {
	Id        int    `json:"-"`
	Name      string `json:"name" db:"name" binding:"required"`
	Surname   string `json:"surname" db:"surname" binding:"required"`
	BirthDate string `json:"birthdate" binding:"required" time_format:"2006-01-02"`
}

type UpdPatient struct {
	Name      *string `json:"name"`
	Surname   *string `json:"surname"`
	BirthDate *string `json:"birthdate" time_format:"2006-01-02"`
}

func (u UpdPatient) Validate() error {
	if u.Name == nil && u.Surname == nil && u.BirthDate == nil {
		errors.New("update structure has no values")
	}

	return nil
}
