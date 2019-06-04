package student

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func CreateStudent(db *mongo.Client, students []Student) error {

	if len(students) == 0 {
		return nil
	}

	collection := db.Database("apc_database").Collection("students")

	for _, student := range students {
		insertResult, err := collection.InsertOne(context.TODO(), student)
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}

	return nil
}
