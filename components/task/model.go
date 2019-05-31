package task

import "plataforma-apc/components/submission"

type Task struct {
	ID          int                     `json:"id"`
	Statement   string                  `json:"statement"`
	Score       float32                 `json:"score"`
	Tags        []string                `json:"tags"`
	Submissions []submission.Submission `json:"submissions"`
}
