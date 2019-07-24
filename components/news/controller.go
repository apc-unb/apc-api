package news

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateNews(db *mongo.Client, news []NewsCreate, databaseName, collectionName string) error {

	if len(news) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

	for _, singleNews := range news {
		if _, err := collection.InsertOne(context.TODO(), singleNews); err != nil {
			return err
		}
	}

	return nil

}

func GetNews(db *mongo.Client, classID primitive.ObjectID, databaseName, collectionName string) ([]News, error) {

	collection := db.Database(databaseName).Collection(collectionName)

	// Here's an array in which you can store the decoded documents
	news := []News{}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(
		context.TODO(),
		bson.M{
			"classID": classID,
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
		var elem News

		// Checks if decoding method didn't return any errors
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		// Push student inside student array
		news = append(news, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return news, nil
}

func UpdateNews(db *mongo.Client, news []News, databaseName, collectionName string) error {

	if len(news) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

	for _, singleNews := range news {
		filter := bson.M{"_id": singleNews.ID}
		update := bson.M{"$set": singleNews}
		if _, err := collection.UpdateOne(context.TODO(), filter, update, nil); err != nil {
			return err
		}
	}
	return nil

}

func DeleteNews(db *mongo.Client, news []News, databaseName, collectionName string) error {

	if len(news) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

	for _, singleNews := range news {
		filter := bson.M{"_id": singleNews.ID}
		if _, err := collection.DeleteOne(context.TODO(), filter); err != nil {
			return err
		}
	}
	return nil

}
