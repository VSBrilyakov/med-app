package medapp

import "time"

type Patient struct {
	Id        int       `json:"-"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	BirthDate time.Time `json:"birthdate"`
}
