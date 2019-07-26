package task

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTasks(db *mongo.Client, tasks []TaskCreate, database_name, collection_name string) error {

	if len(tasks) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, task := range tasks {
		if _, err := collection.InsertOne(context.TODO(), task); err != nil {
			return err
		}
	}

	return nil

}

func GetTasks(db *mongo.Client, database_name, collection_name string) ([]Task, error) {

	collection := db.Database(database_name).Collection(collection_name)

	// Here's an array in which you can store the decoded documents
	tasks := []Task{}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Task

		// Checks if decoding method didn't return any errors
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		// Push school class inside student array
		tasks = append(tasks, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return tasks, nil

}

func GetTasksClass(db *mongo.Client, examID primitive.ObjectID, database_name, collection_name string) ([]Task, error) {

	collection := db.Database(database_name).Collection(collection_name)

	// Here's an array in which you can store the decoded documents
	tasks := []Task{}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(
		context.TODO(),
		bson.M{
			"examid": examID,
		},
		nil,
	)

	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Task

		// Checks if decoding method didn't return any errors
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		// Push school class inside student array
		tasks = append(tasks, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return tasks, nil

}

func UpdateTasks(db *mongo.Client, tasks []Task, database_name, collection_name string) error {

	if len(tasks) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, task := range tasks {
		filter := bson.M{"_id": task.ID}
		update := bson.M{"$set": task}
		if _, err := collection.UpdateOne(context.TODO(), filter, update, nil); err != nil {
			return err
		}
	}
	return nil

}

func DeleteTasks(db *mongo.Client, tasks []Task, database_name, collection_name string) error {

	if len(tasks) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, task := range tasks {
		filter := bson.M{"_id": task.ID}
		if _, err := collection.DeleteOne(context.TODO(), filter); err != nil {
			return err
		}
	}
	return nil

}
