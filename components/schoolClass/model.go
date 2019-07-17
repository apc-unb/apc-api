package schoolClass

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/plataforma-apc/components/student"
)

type SchoolClass struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	ProfessorFirstName string             `json:"professorfirstname"`
	ProfessorLastName  string             `json:"professorlastname"`
	Year               int                `json:"year"`
	Season             int                `json:"season"`
	Students           []student.Student  `json:"students"`
}

type SchoolClassCreate struct {
	ProfessorFirstName string            `json:"professorfirstname"`
	ProfessorLastName  string            `json:"professorlastname"`
	Year               int               `json:"year"`
	Season             int               `json:"season"`
	Students           []student.Student `json:"students"`
}
