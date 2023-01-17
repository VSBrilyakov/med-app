package medapp

type Doctor struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Specialization struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
}

type DocSpec struct {
	DocId  int
	SpecId int
}
