package user

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type UserCredentials struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Matricula string             `json:"matricula"`
	Password  string             `json:"password"`
}
