package medapp

type Doctor struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required" example:"Gregory"`
	Surname  string `json:"surname" binding:"required" example:"House"`
	Login    string `json:"login" binding:"required" example:"DrHouse"`
	Password string `json:"password" binding:"required" example:"ilovemedicine777"`
}

type Specialization struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
}

type DocSpec struct {
	DocId  int
	SpecId int
}
