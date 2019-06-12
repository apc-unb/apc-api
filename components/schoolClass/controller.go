package schoolClass

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func CreateClasses(db *mongo.Client, schoolClass []SchoolClass, database_name, collection_name string) error {
	return errors.New("Function not implemented")
}

func GetClasses(db *mongo.Client, database_name, collection_name string) ([]SchoolClass, error) {
	return nil, errors.New("Function not implemented")
}

func UpdateClasses(db *mongo.Client, schoolClass []SchoolClass, database_name, collection_name string) error {
	return errors.New("Function not implemented")
}

func DeleteClasses(db *mongo.Client, schoolClass []SchoolClass, database_name, collection_name string) error {
	return errors.New("Function not implemented")
}
