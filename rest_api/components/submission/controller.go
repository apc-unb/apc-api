package submission

import (
	"context"
	"reflect"

	"github.com/VerasThiago/plataforma-apc/api/components/student"

	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateSubmissions(db *mongo.Client, submissions []SubmissionCreate, database_name, collection_name string) error {

	if len(submissions) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, submission := range submissions {
		if _, err := collection.InsertOne(context.TODO(), submission); err != nil {
			return err
		}
	}

	return nil
}

func GetSubmissions(db *mongo.Client, database_name, collection_name string) ([]Submission, error) {

	collection := db.Database(database_name).Collection(collection_name)

	// Here's an array in which you can store the decoded documents
	submissions := []Submission{}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Submission

		// Checks if decoding method didn't return any errors
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		// Push school class inside student array
		submissions = append(submissions, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return submissions, nil

}

func UpdateSubmissions(db *mongo.Client, submissions []Submission, database_name, collection_name string) error {

	if len(submissions) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, submission := range submissions {

		filter := bson.M{
			"_id": submission.ID,
		}

		update := bson.M{}

		if reflect.DeepEqual(submission.Student, student.Student{}) {
			update["student"] = submission.Student
		}

		if submission.Time != "" {
			update["time"] = submission.Time
		}

		if submission.Veredict != "" {
			update["veredict"] = submission.Veredict
		}

		updateSet := bson.M{"$set": update}

		if _, err := collection.UpdateOne(context.TODO(), filter, updateSet, nil); err != nil {
			return err
		}
	}
	return nil

}

func DeleteSubmissions(db *mongo.Client, submissions []Submission, database_name, collection_name string) error {

	if len(submissions) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, submission := range submissions {
		filter := bson.M{"_id": submission.ID}
		if _, err := collection.DeleteOne(context.TODO(), filter); err != nil {
			return err
		}
	}
	return nil

}
