package exam

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateExams(db *mongo.Client, exams []ExamCreate, database_name, collection_name string) error {

	if len(exams) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, exam := range exams {
		if _, err := collection.InsertOne(context.TODO(), exam); err != nil {
			return err
		}
	}

	return nil

}

func GetExamsClass(db *mongo.Client, classID primitive.ObjectID, database_name, collection_name string) ([]Exam, error) {

	collection := db.Database(database_name).Collection(collection_name)

	// Here's an array in which you can store the decoded documents
	exams := []Exam{}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(
		context.TODO(),
		bson.M{
			"classid": classID,
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
		var elem Exam

		// Checks if decoding method didn't return any errors
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		// Push school class inside student array
		exams = append(exams, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return exams, nil

}

func GetExams(db *mongo.Client, database_name, collection_name string) ([]Exam, error) {

	collection := db.Database(database_name).Collection(collection_name)

	// Here's an array in which you can store the decoded documents
	exams := []Exam{}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(
		context.TODO(),
		bson.M{},
		nil,
	)

	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Exam

		// Checks if decoding method didn't return any errors
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		// Push school class inside student array
		exams = append(exams, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return exams, nil

}

func UpdateExams(db *mongo.Client, exams []Exam, database_name, collection_name string) error {

	if len(exams) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, exam := range exams {

		filter := bson.M{
			"_id": exam.ID,
		}

		update := bson.M{}

		if exam.Title != "" {
			update["title"] = exam.Title
		}

		if !exam.ClassID.IsZero() {
			update["classid"] = exam.ClassID
		}

		updateSet := bson.M{"$set": update}

		if _, err := collection.UpdateOne(context.TODO(), filter, updateSet, nil); err != nil {
			return err
		}
	}
	return nil

}

func DeleteExams(db *mongo.Client, exams []Exam, database_name, collection_name string) error {

	if len(exams) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, exam := range exams {
		filter := bson.M{"_id": exam.ID}
		if _, err := collection.DeleteOne(context.TODO(), filter); err != nil {
			return err
		}
	}
	return nil

}
