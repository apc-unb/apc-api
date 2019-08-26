package exam

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Exam contains all exam data
type Exam struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	ClassID primitive.ObjectID `bson:"classid,omitempty"`
	Title   string             `json:"title"`
}

// Exam contains all exam data excepted from ID
type ExamCreate struct {
	ClassID primitive.ObjectID `bson:"classid,omitempty"`
	Title   string             `json:"title"`
}
