package news

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type News struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ClassID     primitive.ObjectID `bson:"classID,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Tags        []string           `json:"tags"`
}

type NewsCreate struct {
	ClassID     primitive.ObjectID `bson:"classID,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Tags        []string           `json:"tags"`
}
