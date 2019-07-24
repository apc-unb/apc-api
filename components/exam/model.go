package exam

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/plataforma-apc/components/task"
)

type Exam struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	ClassID primitive.ObjectID `bson:"classID,omitempty"`
	Title   string             `json:"title"`
	Tasks   []task.Task        `json:"tasks"`
}

type ExamCreate struct {
	ClassID primitive.ObjectID `bson:"classID,omitempty"`
	Title   string             `json:"title"`
	Tasks   []task.Task        `json:"tasks"`
}
