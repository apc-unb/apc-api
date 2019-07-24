package schoolClass

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/plataforma-apc/components/exam"
	"github.com/plataforma-apc/components/news"
	"github.com/plataforma-apc/components/student"
)

type SchoolClass struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	ProfessorFirstName string             `json:"professorfirstname"`
	ProfessorLastName  string             `json:"professorlastname"`
	ClassName          string             `json:"classname"`
	Address            string             `json:"address"`
	Year               int                `json:"year"`
	Season             int                `json:"season"`
}

type SchoolClassCreate struct {
	ProfessorFirstName string `json:"professorfirstname"`
	ProfessorLastName  string `json:"professorlastname"`
	ClassName          string `json:"classname"`
	Address            string `json:"address"`
	Year               int    `json:"year"`
	Season             int    `json:"season"`
}

type StudentPage struct {
	UserExist bool                `json:"userexist"`
	User      student.StudentInfo `json:"student"`
	Class     SchoolClass         `json:"class"`
	News      []news.News         `json:"news"`
	Exams     []exam.Exam         `json:"exams"`
}
