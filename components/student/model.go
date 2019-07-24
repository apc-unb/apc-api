package student

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Student struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ClassID   primitive.ObjectID `bson:"classID,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	Handles   []string           `json:"handles"`
	Password  string             `json:"password"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
	Grade     StudentGrades      `json:"grade"`
}

type StudentCreate struct {
	ClassID   primitive.ObjectID `bson:"classID,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	Handles   []string           `json:"handles"`
	Password  string             `json:"password"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
	Grades    StudentGrades      `json:"grades"`
}

type StudentInfo struct {
	ClassID   primitive.ObjectID `bson:"classID,omitempty"`
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	Handles   []string           `json:"handles"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
	Grades    StudentGrades      `json:"grades"`
}

type StudentUpdate struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	NewPassword string             `json:"newpassword"`
}

type StudentLogin struct {
	Matricula string `json:"matricula"`
	Password  string `json:"password"`
}

type StudentGrades struct {
	Exams    []float64 `json:"exams"`
	Projects []float64 `json:"projects"`
	Lists    []float64 `json:"lists"`
}
