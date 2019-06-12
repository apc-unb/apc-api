package submission

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func CreateSubmissions(db *mongo.Client, submission []Submission, database_name, collection_name string) error {
	return errors.New("Function not implemented")
}

func GetSubmissions(db *mongo.Client, database_name, collection_name string) ([]Submission, error) {
	return nil, errors.New("Function not implemented")
}

func UpdateSubmissions(db *mongo.Client, submission []Submission, database_name, collection_name string) error {
	return errors.New("Function not implemented")
}

func DeleteSubmissions(db *mongo.Client, submission []Submission, database_name, collection_name string) error {
	return errors.New("Function not implemented")
}
