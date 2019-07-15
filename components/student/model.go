package student

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/plataforma-apc/components/news"
)

type Student struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	Handles   []string           `json:"handles"`
	Password  string             `json:"password"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
	Grade     float64            `json:"grade"`
}

type StudentCreate struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Matricula string   `json:"matricula"`
	Handles   []string `json:"handles"`
	Password  string   `json:"password"`
	PhotoURL  string   `json:"photourl"`
	Email     string   `json:"email"`
}

type StudentInfo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	Handles   []string           `json:"handles"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
	Grade     float64            `json:"grade"`
}

type StudentUpdate struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `json:"firstname"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
}

type StudentLogin struct {
	Matricula string `json:"matricula"`
	Password  string `json:"password"`
}

type StudentPage struct {
	UserExist bool        `json:"userexist"`
	User      StudentInfo `json:"student"`
	News      []news.News `json:"news"`
}
