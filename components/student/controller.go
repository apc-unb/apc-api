package student

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateStudent(db *mongo.Client, students []Student, database_name, collection_name string) error {

	if len(students) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, student := range students {
		_, err := collection.InsertOne(context.TODO(), student)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetStudents(db *mongo.Client, database_name, collection_name string) ([]Student, error) {

	collection := db.Database(database_name).Collection(collection_name)

	// Here's an array in which you can store the decoded documents
	students := []Student{}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Student

		// Checks if decoding method didn't return any errors
		if err := cur.Decode(&elem); err != nil {
			return nil, err
		}

		// Push student inside student array
		students = append(students, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return students, nil
}
