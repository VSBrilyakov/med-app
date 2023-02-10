package medapp

type Patient struct {
	Id        int    `json:"-"`
	Name      string `json:"name" db:"name" binding:"required"`
	Surname   string `json:"surname" db:"surname" binding:"required"`
	BirthDate string `json:"birthdate" binding:"required" time_format:"2006-01-02"`
}
