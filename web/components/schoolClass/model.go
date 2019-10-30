package schoolClass

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type SchoolClass struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	ProfessorID        primitive.ObjectID `bson:"professorid,omitempty"`
	ProfessorFirstName string             `json:"professorfirstname"`
	ProfessorLastName  string             `json:"professorlastname"`
	ClassName          string             `json:"classname"`
	Address            string             `json:"address"`
	Year               int                `json:"year"`
	Season             int                `json:"season"`
	ContestsIDs        []int 			  `json:"contestsids"`
	GroupID 		   string   	      `json:"groupid"`
}

type SchoolClassCreate struct {
	ProfessorID        primitive.ObjectID `bson:"professorid,omitempty"`
	ProfessorFirstName 		string		   `json:"professorfirstname"`
	ProfessorLastName  		string		   `json:"professorlastname"`
	ClassName          		string		   `json:"classname"`
	Address            		string		   `json:"address"`
	Year             		int			   `json:"year"`
	Season           		int 		   `json:"season"`
	ContestsIDs        		[]int		   `json:"contestsids"`
	GroupID 		  		string   	   `json:"groupid"`
}