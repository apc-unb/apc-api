package task

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/plataforma-apc/components/submission"
)

type Task struct {
	ID          primitive.ObjectID      `bson:"_id,omitempty"`
	Statement   string                  `json:"statement"`
	Score       float32                 `json:"score"`
	Tags        []string                `json:"tags"`
	Submissions []submission.Submission `json:"submissions"`
}

type TaskCreate struct {
	Statement   string                  `json:"statement"`
	Score       float32                 `json:"score"`
	Tags        []string                `json:"tags"`
	Submissions []submission.Submission `json:"submissions"`
}
