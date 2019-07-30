package admin

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/plataforma-apc/components/student"
)

type Admin struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ClassID   primitive.ObjectID `bson:"classid,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	Password  string             `json:"password"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
}

type AdminCreate struct {
	ClassID   primitive.ObjectID `bson:"classid,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	Password  string             `json:"password"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
}

type AdminInfo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ClassID   primitive.ObjectID `bson:"classid,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
}

type AdminUpdate struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ClassID     primitive.ObjectID `bson:"classid,omitempty"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	NewPassword string             `json:"newpassword"`
	PhotoURL    string             `json:"photourl"`
}

type AdminUpdateStudent struct {
	StudentID primitive.ObjectID     `bson:"_id,omitempty"`
	ClassID   primitive.ObjectID     `bson:"classid,omitempty"`
	FirstName string                 `json:"firstname"`
	LastName  string                 `json:"lastname"`
	Matricula string                 `json:"matricula"`
	Handles   student.StudentHandles `json:"handles"`
	PhotoURL  string                 `json:"photourl"`
	Email     string                 `json:"email"`
	Grades    student.StudentGrades  `json:"grades"`
}

type AdminLogin struct {
	Matricula string `json:"matricula"`
	Password  string `json:"password"`
}
