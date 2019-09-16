package task

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ExamID    primitive.ObjectID `bson:"examid,omitempty"`
	Title     string             `json:"title"`
	Statement string             `json:"statement"`
	Score     float32            `json:"score"`
	Tags      []string           `json:"tags"`
}

type TaskCreate struct {
	ExamID    primitive.ObjectID `bson:"examid,omitempty"`
	Title     string             `json:"title"`
	Statement string             `json:"statement"`
	Score     float32            `json:"score"`
	Tags      []string           `json:"tags"`
}
