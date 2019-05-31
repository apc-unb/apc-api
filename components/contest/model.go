package contest

import (
	"plataforma-apc/components/schoolClass"
	"plataforma-apc/components/task"
)

type Contest struct {
	ID    int                     `json:"id"`
	Date  string                  `json:"date"`
	Class schoolClass.SchoolClass `json:"class"`
	Tasks []task.Task             `json:"tasks"`
}
