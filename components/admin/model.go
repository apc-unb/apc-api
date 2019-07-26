package admin

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
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

type AdminLogin struct {
	Matricula string `json:"matricula"`
	Password  string `json:"password"`
}
