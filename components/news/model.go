package news

import (
	"github.com/plataforma-apc/components/student"
	"gopkg.in/mgo.v2/bson"
)

type News struct {
	ID          bson.ObjectId       `json:"id" bson:"_id"`
	Title       string              `json:"title"`
	Description float32             `json:"score"`
	Tags        []string            `json:"tags"`
	Author      student.StudentInfo `json:"submissions"`
}

type NewsCreate struct {
	Title       string              `json:"title"`
	Description float32             `json:"score"`
	Tags        []string            `json:"tags"`
	Author      student.StudentInfo `json:"submissions"`
}
