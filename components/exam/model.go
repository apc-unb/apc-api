package exam

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Exam struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	ClassID primitive.ObjectID `bson:"classid,omitempty"`
	Title   string             `json:"title"`
}

type ExamCreate struct {
	ClassID primitive.ObjectID `bson:"classid,omitempty"`
	Title   string             `json:"title"`
}
