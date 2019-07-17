package contest

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/plataforma-apc/components/schoolClass"
	"github.com/plataforma-apc/components/task"
)

type Contest struct {
	ID    primitive.ObjectID      `bson:"_id,omitempty"`
	Date  string                  `json:"date"`
	Class schoolClass.SchoolClass `json:"class"`
	Tasks []task.Task             `json:"tasks"`
}

type ContestCreate struct {
	Date  string                  `json:"date"`
	Class schoolClass.SchoolClass `json:"class"`
	Tasks []task.Task             `json:"tasks"`
}
