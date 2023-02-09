package medapp

type Doctor struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name" binding:"required"`
	Surname  string `json:"surname" db:"surname" binding:"required"`
	Login    string `json:"login" db:"login" binding:"required"`
	Password string `json:"-" binding:"required"`
}

type Specialization struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
}

type DocSpec struct {
	DocId  int
	SpecId int
}
