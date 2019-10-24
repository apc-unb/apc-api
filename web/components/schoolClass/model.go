package schoolClass

import (
	"github.com/apc-unb/apc-api/web/components/admin"
	"github.com/apc-unb/apc-api/web/components/news"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type SchoolClass struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	ProfessorFirstName string             `json:"professorfirstname"`
	ProfessorLastName  string             `json:"professorlastname"`
	ClassName          string             `json:"classname"`
	Address            string             `json:"address"`
	Year               int                `json:"year"`
	Season             int                `json:"season"`
	ContestsIDs        []int 			  `json:"contestsids"`
	GroupID 		  		string   	   `json:"groupid"`
}

type SchoolClassCreate struct {
	ProfessorFirstName 		string		   `json:"professorfirstname"`
	ProfessorLastName  		string		   `json:"professorlastname"`
	ClassName          		string		   `json:"classname"`
	Address            		string		   `json:"address"`
	Year             		int			   `json:"year"`
	Season           		int 		   `json:"season"`
	ContestsIDs        		[]int		   `json:"contestsids"`
	GroupID 		  		string   	   `json:"groupid"`
}

type AdminPage struct {
	UserExist bool            `json:"userexist"`
	Admin     admin.AdminInfo `json:"admin"`
	Class     SchoolClass     `json:"class"`
	News      []news.News     `json:"news"`
}
