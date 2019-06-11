package student

import "gopkg.in/mgo.v2/bson"

type Student struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
	Matricula string        `json:"matricula"`
	Handles   []string      `json:"handles"`
	Password  string        `json:"password"`
	PhotoUrl  string        `json:"photourl"`
	Grade     float64       `json:"grade"`
}
