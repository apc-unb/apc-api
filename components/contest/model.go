package contest

import (
	"plataforma-apc/components/schoolClass"
	"plataforma-apc/components/task"
	"gopkg.in/mgo.v2/bson"
)

type Contest struct {
	ID    bson.ObjectId     	  `json:"id" bson:"_id"`
	Date  string                  `json:"date"`
	Class schoolClass.SchoolClass `json:"class"`
	Tasks []task.Task             `json:"tasks"`
}

