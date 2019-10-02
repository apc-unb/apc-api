package submission

import (
	"github.com/apc-unb/apc-api/web/components/student"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// TODO : Check how submissions time are made in Pimenta Judge
// TODO : Decide if veredict gonna be num code or string
type Submission struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Student  student.Student    `json:"student"`
	Veredict string             `json:"veredict"`
	Time     string             `json:"time"`
}

type SubmissionCreate struct {
	Student  student.Student `json:"student"`
	Veredict string          `json:"veredict"`
	Time     string          `json:"time"`
}
