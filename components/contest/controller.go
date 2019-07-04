package contest

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateContests(db *mongo.Client, contests []Contest, database_name, collection_name string) error {
	
	if len(contests) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, contest := range contests {
		if _, err := collection.InsertOne(context.TODO(), contest); err != nil {
			return err
		}
	}

	return nil
	
}

func GetContests(db *mongo.Client, database_name, collection_name string) ([]Contest, error) {
	
	collection := db.Database(database_name).Collection(collection_name)

	// Here's an array in which you can store the decoded documents
	contests := []Contest{}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Contest

		// Checks if decoding method didn't return any errors
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		// Push school class inside student array
		contests = append(contests, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return contests, nil

}

func UpdateContests(db *mongo.Client, contests []Contest, database_name, collection_name string) error {
	
	if len(contests) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, contest := range contests {
		filter := bson.M{"_id": contest.ID}
		update := bson.M{"$set": contest}
		if _, err := collection.UpdateOne(context.TODO(), filter, update, nil); err != nil {
			return err
		}
	}
	return nil

}

func DeleteContests(db *mongo.Client, contests []Contest, database_name, collection_name string) error {
	
	if len(contests) == 0 {
		return nil
	}

	collection := db.Database(database_name).Collection(collection_name)

	for _, contest := range contests {
		filter := bson.M{"_id": contest.ID}
		if _, err := collection.DeleteOne(context.TODO(), filter); err != nil {
			return err
		}
	}
	return nil
	
}
