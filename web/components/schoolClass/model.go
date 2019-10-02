package schoolClass

import (
	"github.com/apc-unb/apc-api/web/components/admin"
	"github.com/apc-unb/apc-api/web/components/news"
	"github.com/apc-unb/apc-api/web/components/student"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
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
	Result    string              `json:"result"`
	Student   student.StudentInfo `json:"student"`
	Class     SchoolClass         `json:"class"`
	News      []news.News         `json:"news"`
}

type AdminPage struct {
	UserExist bool            `json:"userexist"`
	Admin     admin.AdminInfo `json:"admin"`
	Class     SchoolClass     `json:"class"`
	News      []news.News     `json:"news"`
}
