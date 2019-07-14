package submission

import (
	"github.com/plataforma-apc/components/student"
	"gopkg.in/mgo.v2/bson"
)

// TODO : Check how submissions time are made in Pimenta Judge
// TODO : Decide if veredict gonna be num code or string
type Submission struct {
	ID       bson.ObjectId   `json:"id" bson:"_id"`
	Student  student.Student `json:"student"`
	Veredict string          `json:"veredict"`
	Time     string          `json:"time"`
}
