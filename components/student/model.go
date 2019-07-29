package student

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Student struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ClassID   primitive.ObjectID `bson:"classid,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	Handles   StudentHandles     `json:"handles"`
	Password  string             `json:"password"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
	Grades    StudentGrades      `json:"grades"`
}

type StudentCreate struct {
	ClassID   primitive.ObjectID `bson:"classid,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	Handles   StudentHandles     `json:"handles"`
	Password  string             `json:"password"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
	Grades    StudentGrades      `json:"grades"`
}

type StudentInfo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ClassID   primitive.ObjectID `bson:"classid,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	Handles   StudentHandles     `json:"handles"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
	Grades    StudentGrades      `json:"grades"`
}

type StudentUpdate struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	NewPassword string             `json:"newpassword"`
	Handles     StudentHandles     `json:"handles"`
	PhotoURL    string             `json:"photourl"`
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

type StudentHandles struct {
	Codeforces string `json:"codeforces"`
	Uri        string `json:"uri"`
}
