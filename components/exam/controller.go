package exam

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateExams receive  a list of exams
// Checks if that list is not null (can't insert null list)
// Insert each exams individually in database
// @param	db				pointer to database
// @param	exams 		list of exams
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	error 			function error
// TODO : Insert all exams at the same time (if possible)
func CreateExams(db *mongo.Client, exams []ExamCreate, databaseName, collectionName string) error {

	if len(exams) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

	for _, exam := range exams {
		if _, err := collection.InsertOne(context.TODO(), exam); err != nil {
			return err
		}
	}

	return nil

}

// GetExamsClass return list of all exams from a current class
// Get class id to filter in mongodb
// @param	db				pointer to database
// @param   classID			class ID
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	[]Exam			list of all exams that match with the given class id
// @return 	error 			function error
func GetExamsClass(db *mongo.Client, classID primitive.ObjectID, databaseName, collectionName string) ([]Exam, error) {

	collection := db.Database(databaseName).Collection(collectionName)

	exams := []Exam{}

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

	for cursor.Next(context.TODO()) {

		var elem Exam

		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		exams = append(exams, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(context.TODO())

	return exams, nil

}

// GetExams return list of all exams from Database
// Get all exams at the same time and store inside cursor
// Decode each exam inside exam class and append into exams array
// @param	db				pointer to database
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	[]Exam		list of all exams
// @return 	error 			function error
func GetExams(db *mongo.Client, databaseName, collectionName string) ([]Exam, error) {

	collection := db.Database(databaseName).Collection(collectionName)

	exams := []Exam{}

	cursor, err := collection.Find(
		context.TODO(),
		bson.M{},
		nil,
	)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {

		var elem Exam

		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		exams = append(exams, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(context.TODO())

	return exams, nil

}

// UpdateExams receive list of exams (updated)
// Checks if exams old password matches with db to update that exams password or email
// @param	db				pointer to database (updated)
// @param	exams 			list of exams
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	error 			function error
func UpdateExams(db *mongo.Client, exams []Exam, databaseName, collectionName string) error {

	if len(exams) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

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

// DeleteExams recieve a list of exams (to be deleted)
// Checks if that list is not null (can't delete null list)
// Delete each exam individually
// @param	db				pointer to database (to be deleted)
// @param	exams	 		list of exams
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	error 			function error
// TODO : Delete all exams at the same time (if possible)
func DeleteExams(db *mongo.Client, exams []Exam, databaseName, collectionName string) error {

	if len(exams) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

	for _, exam := range exams {
		filter := bson.M{"_id": exam.ID}
		if _, err := collection.DeleteOne(context.TODO(), filter); err != nil {
			return err
		}
	}
	return nil

}
