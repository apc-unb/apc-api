package student

type Student struct {
	ID        int      `json:"id"`
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Matricula string   `json:"matricula"`
	Handles   []string `json:"handles"`
	Password  string   `json:"password"`
	PhotoUrl  string   `json:"photourl"`
	Grade     float64  `json:"grade"`
}
