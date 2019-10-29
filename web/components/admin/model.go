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
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
	Projects  int32              `json:"projects"`
	Teacher   bool               `json:"teacher"`
}

// AdminCreate contais all admin data except from ID
type AdminCreate struct {
	ClassID   primitive.ObjectID `bson:"classid,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
	Projects  int32              `json:"projects"`
	Teacher   bool               `json:"teacher"`
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
	Projects  int32              `json:"projects"`
	Teacher   bool               `json:"teacher"`
}

// AdminUpdate contais all data that admin can update
type AdminUpdate struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ClassID   primitive.ObjectID `bson:"classid,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Matricula string             `json:"matricula"`
	PhotoURL  string             `json:"photourl"`
	Email     string             `json:"email"`
	Projects  int32              `json:"projects"`
	Password    string           `json:"password"`
	NewPassword string           `json:"newpassword"`
}

// AdminUpdateStudent all data of a student to be updated
type AdminUpdateStudent struct {
	AdminID 			primitive.ObjectID 	   `bson:"adminid,omitempty"`
	StudentID			primitive.ObjectID     `bson:"studentid,omitempty"`
	ClassID   			primitive.ObjectID     `bson:"classid,omitempty"`
	AdminPassword    	string             	   `json:"adminpassword"`
	FirstName 			string                 `json:"firstname"`
	LastName  			string                 `json:"lastname"`
	Matricula 			string                 `json:"matricula"`
	Handles   			student.StudentHandles `json:"handles"`
	PhotoURL  			string                 `json:"photourl"`
	Email     			string                 `json:"email"`
	Grades    			student.StudentGrades  `json:"grades"`
}