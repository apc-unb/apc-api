package schoolClass

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateClasses(db *mongo.Client, schoolClass []SchoolClassCreate, database_name, collection_name string) error {

	if len(schoolClass) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, class := range schoolClass {
		if _, err := collection.InsertOne(context.TODO(), class); err != nil {
			return err
		}
	}

	return nil

}

func GetClass(db *mongo.Client, classID primitive.ObjectID, databaseName, collectionName string) (SchoolClass, error) {

	collection := db.Database(databaseName).Collection(collectionName)

	findClass := SchoolClass{}

	filter := bson.D{
		{"_id", classID},
	}

	if err := collection.FindOne(
		context.TODO(),
		filter,
	).Decode(&findClass); err != nil {
		return findClass, err
	}

	return findClass, nil
}

func GetClasses(db *mongo.Client, database_name, collection_name string) ([]SchoolClass, error) {

	collection := db.Database(database_name).Collection(collection_name)

	// Here's an array in which you can store the decoded documents
	classes := []SchoolClass{}

	var options options.FindOptions

	options.SetSort(bson.D{{"year", -1}, {"season", -1}, {"classname", 1}})

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(
		context.TODO(),
		bson.D{{}},
		&options,
	)
	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem SchoolClass

		// Checks if decoding method didn't return any errors
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		// Push school class inside student array
		classes = append(classes, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return classes, nil

}

func UpdateClasses(db *mongo.Client, schoolClass []SchoolClass, database_name, collection_name string) error {

	if len(schoolClass) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, schoolClass := range schoolClass {

		filter := bson.M{
			"_id": schoolClass.ID,
		}

		update := bson.M{}

		if schoolClass.ProfessorFirstName != "" {
			update["professorfirstname"] = schoolClass.ProfessorFirstName
		}

		if schoolClass.ProfessorLastName != "" {
			update["professorlastname"] = schoolClass.ProfessorLastName
		}

		if schoolClass.ClassName != "" {
			update["classname"] = schoolClass.ClassName
		}

		if schoolClass.GroupID != "" {
			update["groupid"] = schoolClass.GroupID
		}

		if schoolClass.Address != "" {
			update["address"] = schoolClass.Address
		}

		if schoolClass.Year != 0 {
			update["year"] = schoolClass.Year
		}

		if schoolClass.Season != 0 {
			update["season"] = schoolClass.Season
		}

		updateSet := bson.M{"$set": update}

		if _, err := collection.UpdateOne(context.TODO(), filter, updateSet, nil); err != nil {
			return err
		}
	}
	return nil
}

func DeleteClasses(db *mongo.Client, schoolClass []SchoolClass, database_name, collection_name string) error {

	if len(schoolClass) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, schoolClass := range schoolClass {
		filter := bson.M{"_id": schoolClass.ID}
		if _, err := collection.DeleteOne(context.TODO(), filter); err != nil {
			return err
		}
	}
	return nil
}
