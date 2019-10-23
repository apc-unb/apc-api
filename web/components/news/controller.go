package news

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

func GetNews(db *mongo.Client, databaseName, collectionName string) ([]News, error) {

	collection := db.Database(databaseName).Collection(collectionName)

	// Here's an array in which you can store the decoded documents
	news := []News{}

	var options options.FindOptions

	options.SetSort(bson.D{{"updatedat", -1}})

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(
		context.TODO(),
		bson.M{},
		&options,
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

func GetNewsClass(db *mongo.Client, classID primitive.ObjectID, databaseName, collectionName string) ([]News, error) {

	collection := db.Database(databaseName).Collection(collectionName)

	// Here's an array in which you can store the decoded documents
	news := []News{}

	var options options.FindOptions

	options.SetSort(bson.D{{"updatedat", -1}})

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(
		context.TODO(),
		bson.M{
			"classid": classID,
		},
		&options,
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

func CreateNews(db *mongo.Client, singleNews NewsCreate, databaseName, collectionName string) error {

	collection := db.Database(databaseName).Collection(collectionName)

	singleNews.CreatedAT = time.Now()
	singleNews.UpdatedAT = time.Now()

	if _, err := collection.InsertOne(context.TODO(), singleNews); err != nil {
		return err
	}

	return nil

}

func UpdateNews(db *mongo.Client, singleNews News, databaseName, collectionName string) error {

	collection := db.Database(databaseName).Collection(collectionName)


	filter := bson.M{
		"_id": singleNews.ID,
	}

	update := bson.M{}

	if !singleNews.ClassID.IsZero() {
		update["classid"] = singleNews.ClassID
	}

	if singleNews.Title != "" {
		update["title"] = singleNews.Title
	}

	if singleNews.Description != "" {
		update["description"] = singleNews.Description
	}

	if len(singleNews.Tags) > 0 {
		update["tags"] = singleNews.Tags
	}

	update["updatedat"] = time.Now()

	updateSet := bson.M{"$set": update}

	if _, err := collection.UpdateOne(context.TODO(), filter, updateSet, nil); err != nil {
		return err
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
