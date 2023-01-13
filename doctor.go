package medapp

type Doctor struct {
	Id      int    `json:"-"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Specialization struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
}

type DocSpec struct {
	DocId  int
	SpecId int
}
