package project

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)


const (
	Created = "Created"
	Updated = "Updated"
	Pending = "Pending"
	Received = "Received"
	Confirmed = "Confirmed"
)

type ProjectType struct {
	ID       	*primitive.ObjectID `bson:"_id,omitempty"`
	Name     	string              `json:"name"`
	Description string              `json:"descripton"`
	ClassID  	primitive.ObjectID  `bson:"classid,omitempty"`
	Start 		time.Time           `json:"start"`
	End 		time.Time           `json:"end"`
	Score    	float64             `json:"score"`
}

type Project struct {
	ID            *primitive.ObjectID `bson:"_id,omitempty"`
	StudentID     primitive.ObjectID  `bson:"studentid,omitempty"`
	ProjectTypeID primitive.ObjectID  `bson:"projectypeid,omitempty"`
	ClassID  	primitive.ObjectID  `bson:"classid,omitempty"`
	MonitorID     primitive.ObjectID  `bson:"monitorid,omitempty"`
	CreatedAT     time.Time           `json:"createdat,omitempty"`
	UpdatedAT     time.Time           `json:"updatedat,omitempty"`
	FileName      string              `json:"filename,omitempty"`
	Status        string              `json:"status,omitempty"`
	Score         float64             `json:"score,omitempty"`
}