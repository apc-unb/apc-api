package submission

import (
	"plataforma-apc/components/student"
)

// TODO : Check how submissions time are made in Pimenta Judge
// TODO : Decide if veredict gonna be num code or string
type Submission struct {
	ID       int             `json:"id"`
	Student  student.Student `json:"student"`
	Veredict string          `json:"veredict"`
	Time     string          `json:"time"`
}
