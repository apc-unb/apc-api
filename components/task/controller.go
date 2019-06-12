package task

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func CreateTasks(db *mongo.Client, task []Task, database_name, collection_name string) error {
	return errors.New("Function not implemented")
}

func GetTasks(db *mongo.Client, database_name, collection_name string) ([]Task, error) {
	return nil, errors.New("Function not implemented")
}

func UpdateTasks(db *mongo.Client, task []Task, database_name, collection_name string) error {
	return errors.New("Function not implemented")
}

func DeleteTasks(db *mongo.Client, task []Task, database_name, collection_name string) error {
	return errors.New("Function not implemented")
}
