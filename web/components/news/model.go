package news

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type News struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ClassID     primitive.ObjectID `bson:"classid,omitempty"`
	AuthorID	primitive.ObjectID `bson:"authorID,omitempty"`
	AuthorName  string             `json:"authorName,omitempty"`
	Admin       bool 			   `bson:"admin,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Tags        []string           `json:"tags"`
	CreatedAT   time.Time          `json:"createdat"`
	UpdatedAT   time.Time          `json:"updatedat"`
}

type NewsCreate struct {
	ClassID     primitive.ObjectID `bson:"classid,omitempty"`
	AuthorID	primitive.ObjectID `bson:"authorID,omitempty"`
	AuthorName  string             `json:"authorName,omitempty"`
	Admin       bool 			   `bson:"admin,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Tags        []string           `json:"tags"`
	CreatedAT   time.Time          `json:"createdat"`
	UpdatedAT   time.Time          `json:"updatedat"`
}
