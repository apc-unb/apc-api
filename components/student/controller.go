package student

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

// TODO : Insert all students at the same time (if possible)
// Recieve a list of students
// Checks if that list is not null (can't insert null list)
// Insert each student individually in database
// @param	db				pointer to database
// @param	students 		list of students
// @param	database_name	name of database
// @param	collection_name	name of collection
// @return 	error 			function error

func CreateStudents(db *mongo.Client, students []Student, database_name, collection_name string) error {

	if len(students) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, student := range students {
		if _, err := collection.InsertOne(context.TODO(), student); err != nil {
			return err
		}
	}

	return nil
}

// Return list of all students from Database
// Get all students at the same time and store inside cursor
// Decode each student inside student class and append into students array
// @param	db				pointer to database
// @param	database_name	name of database
// @param	collection_name	name of collection
// @return 	[]Student		list of all students
// @return 	error 			function error

func GetStudents(db *mongo.Client, database_name, collection_name string) ([]Student, error) {

	collection := db.Database(database_name).Collection(collection_name)

	// Here's an array in which you can store the decoded documents
	students := []Student{}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Student

		// Checks if decoding method didn't return any errors
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		// Push student inside student array
		students = append(students, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return students, nil
}

// TODO : Update all students at the same time (if possible)
// Recieve a list of students (updated)
// Checks if that list is not null (can't update null list)
// Update each student individually
// @param	db				pointer to database (updated)
// @param	students 		list of students
// @param	database_name	name of database
// @param	collection_name	name of collection
// @return 	[]Student		list of all students
// @return 	error 			function error

func UpdateStudents(db *mongo.Client, students []Student, database_name, collection_name string) error {

	if len(students) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, student := range students {
		filter := bson.M{"_id": student.ID}
		update := bson.M{"$set": student}
		if _, err := collection.UpdateOne(context.TODO(), filter, update, nil); err != nil {
			return err
		}
	}
	return nil
}

// TODO : Delete all students at the same time (if possible)
// Recieve a list of students (to be deleted)
// Checks if that list is not null (can't delete null list)
// Delete each student individually
// @param	db				pointer to database (to be deleted)
// @param	students 		list of students
// @param	database_name	name of database
// @param	collection_name	name of collection
// @return 	[]Student		list of all students
// @return 	error 			function error

func DeleteStudents(db *mongo.Client, students []Student, database_name, collection_name string) error {

	if len(students) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, student := range students {
		filter := bson.M{"_id": student.ID}
		if _, err := collection.DeleteOne(context.TODO(), filter); err != nil {
			return err
		}
	}
	return nil

}
