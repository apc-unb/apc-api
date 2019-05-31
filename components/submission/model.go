package submission

import (
	"plataforma-apc/components/student"
)

type Submission struct {
	ID       int             `json:"id"`
	Student  student.Student `json:"student"`
	Veredict string          `json:"veredict"`
	Time     string          `json:"time"`
}
