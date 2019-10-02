package admin

import (
	"github.com/apc-unb/apc-api/web/components/student"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Admin contains all admin data
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

// AdminCreate contais all admin data except from ID
type AdminCreate struct {
	ClassID   primitive.ObjectID `bson:"classid,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	Password  string             `json:"password"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
}

// AdminInfo contais all admin data except from Password
type AdminInfo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ClassID   primitive.ObjectID `bson:"classid,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
}

// AdminUpdate contais all data that admin can update
type AdminUpdate struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ClassID     primitive.ObjectID `bson:"classid,omitempty"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	NewPassword string             `json:"newpassword"`
	PhotoURL    string             `json:"photourl"`
}

// AdminUpdateStudent all data of a student to be updated
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

// AdminLogin login data
type AdminLogin struct {
	Matricula string `json:"matricula"`
	Password  string `json:"password"`
}
