package project

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type ProjectType struct {
	ID       *primitive.ObjectID `bson:"_id,omitempty"`
	Name     string              `json:"name"`
	Order    int                 `json:"order"`
	DeadLine time.Time           `json:"deadline"`
	Score    float64             `json:"score"`
}

type Project struct {
	ID            *primitive.ObjectID `bson:"_id,omitempty"`
	StudentID     primitive.ObjectID  `bson:"studentID,omitempty"`
	ProjectTypeID primitive.ObjectID  `bson:"projectypeID,omitempty"`
	MonitorID     primitive.ObjectID  `bson:"monitorID,omitempty"`
	ClassID       primitive.ObjectID  `bson:"classID,omitempty"`
	SendTime      time.Time           `json:"time,omitempty"`
	FileName      string              `json:"filename,omitempty"`
	Status        string              `json:"status,omitempty"`
	Score         float64             `json:"score,omitempty"`
}
