package task

import (
	"github.com/plataforma-apc/components/submission"
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	ID          bson.ObjectId           `json:"id" bson:"_id"`
	Statement   string                  `json:"statement"`
	Score       float32                 `json:"score"`
	Tags        []string                `json:"tags"`
	Submissions []submission.Submission `json:"submissions"`
}
