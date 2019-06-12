package schoolClass

import (
	"plataforma-apc/components/student"

	"gopkg.in/mgo.v2/bson"
)

type SchoolClass struct {
	ID                 bson.ObjectId     `json:"id" bson:"_id"`
	ProfessorFirstName string            `json:"professorfirstname"`
	ProfessorLastName  string            `json:"professorlastname"`
	Year               int               `json:"year"`
	Season             int               `json:"season"`
	Students           []student.Student `json:"students"`
}
