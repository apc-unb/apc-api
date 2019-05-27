package schoolClass

import (
	"plataforma-apc/components/student"
)

type SchoolClass struct {
	ID                 int               `json:"id"`
	ProfessorFirstName string            `json:"professorfirstname"`
	ProfessorLastName  string            `json:"professorlastname"`
	Year               int               `json:"year"`
	Season             int               `json:"season"`
	Students           []student.Student `json:"students"`
}
