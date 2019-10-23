package admin

import (
	"github.com/apc-unb/apc-api/web/components/student"
	"github.com/apc-unb/apc-api/web/components/user"
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
}

// AdminUpdate contais all data that admin can update
type AdminUpdate struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ClassID     primitive.ObjectID `bson:"classid,omitempty"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	NewPassword string             `json:"newpassword"`
	PhotoURL    string             `json:"photourl"`
	Projects    int32              `json:"projects"`
}

// AdminUpdateStudent all data of a student to be updated
type AdminUpdateStudent struct {
	AdminID 			primitive.ObjectID 	   `bson:"adminid,omitempty"`
	AdminPassword    	string             	   `json:"adminpassword"`
	StudentID			primitive.ObjectID     `bson:"studentid,omitempty"`
	ClassID   			primitive.ObjectID     `bson:"classid,omitempty"`
	FirstName 			string                 `json:"firstname"`
	LastName  			string                 `json:"lastname"`
	Matricula 			string                 `json:"matricula"`
	Handles   			student.StudentHandles `json:"handles"`
	PhotoURL  			string                 `json:"photourl"`
	Email     			string                 `json:"email"`
	Grades    			student.StudentGrades  `json:"grades"`
}

type AdminCreatePage struct {
	Result string                 `json:"result"`
	Admins []user.UserCredentials `json:"admins"`
}
