package admin

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"strings"

	"github.com/togatoga/goforces"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateStudents recieve a list of students
// Checks if that list is not null (can't insert null list)
// Insert each student individually in database
// @param	db				pointer to database
// @param	students 		list of students
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	error 			function error
// TODO : Insert all students at the same time (if possible)
func CreateAdmin(db *mongo.Client, api *goforces.Client, admins []AdminCreate, databaseName, collectionName string) error {

	if len(admins) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

	for _, admin := range admins {

		admin.Password = generateRandomPassword()

		if _, err := collection.InsertOne(context.TODO(), admin); err != nil {
			return err
		}
	}

	return nil
}

func CreateAdminFile(db *mongo.Client, request string, databaseName, collectionName string) error {

	var admins []AdminCreate
	var err error

	if admins, err = getAdminFromFile(db, request); err != nil {
		return err
	}

	if len(admins) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

	for _, admin := range admins {
		if _, err := collection.InsertOne(context.TODO(), admin); err != nil {
			return err
		}
	}

	return nil
}

// GetStudents return list of all students from Database
// Get all students at the same time and store inside cursor
// Decode each student inside student class and append into students array
// @param	db				pointer to database
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	[]Student		list of all students
// @return 	error 			function error
func GetAdmins(db *mongo.Client, databaseName, collectionName string) ([]AdminInfo, error) {

	collection := db.Database(databaseName).Collection(collectionName)

	// Here's an array in which you can store the decoded documents
	admins := []AdminInfo{}

	//
	projection := bson.D{
		{"password", 0},
	}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(
		context.TODO(),
		bson.D{{}},
		options.Find().SetProjection(projection),
	)

	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem AdminInfo

		// Checks if decoding method didn't return any errors
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		// Push student inside student array
		admins = append(admins, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return admins, nil
}

// UpdateStudents recieve student (updated)
// Checks if student old password matches with db to update that student password or email
// @param	db				pointer to database (updated)
// @param	students 		list of students
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	StudentUpdate	student new data
// @return 	error 			function error
// TODO : Update all students at the same time (if possible)
func UpdateAdmins(db *mongo.Client, api *goforces.Client, admin AdminUpdate, databaseName, collectionName string) error {

	collection := db.Database(databaseName).Collection(collectionName)

	currentAdmin := AdminUpdate{}

	filter := bson.M{
		"_id":      admin.ID,
		"password": admin.Password,
	}

	projection := bson.M{
		"_id": 1,
	}

	if err := collection.FindOne(
		context.TODO(),
		filter,
		options.FindOne().SetProjection(projection),
	).Decode(&currentAdmin); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("Invalid password")
		}
		return err
	}

	filter = bson.M{
		"_id": admin.ID,
	}

	update := bson.M{}

	if admin.Email != "" {
		update["email"] = admin.Email
	}

	if admin.NewPassword != "" {
		update["password"] = admin.NewPassword
	}

	if admin.PhotoURL != "" {
		update["photourl"] = admin.PhotoURL
	}

	if admin.ClassID.String() != "" {
		update["classid"] = admin.ClassID
	}

	updateSet := bson.M{"$set": update}

	if _, err := collection.UpdateOne(context.TODO(), filter, updateSet, nil); err != nil {
		return err
	}

	return nil
}

// DeleteStudents recieve a list of students (to be deleted)
// Checks if that list is not null (can't delete null list)
// Delete each student individually
// @param	db				pointer to database (to be deleted)
// @param	students 		list of students
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	[]Student		list of all students
// @return 	error 			function error
// TODO : Delete all students at the same time (if possible)
func DeleteStudents(db *mongo.Client, admins []Admin, databaseName, collectionName string) error {

	if len(admins) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

	for _, admin := range admins {
		filter := bson.M{"_id": admin.ID}
		if _, err := collection.DeleteOne(context.TODO(), filter); err != nil {
			return err
		}
	}
	return nil

}

// AuthStudent recieve a student (to be authenticated)
// Checks if that date exist in databse
// Return true if exist
// @param	db				pointer to database (to be deleted)
// @param	student			student matricula and password
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	[]bool			user exist veredict
// @return 	error 			function error
func AuthAdmin(db *mongo.Client, admin AdminLogin, databaseName, collectionName string) (AdminInfo, error) {

	collection := db.Database(databaseName).Collection(collectionName)

	findAdmin := AdminInfo{}

	filter := bson.D{
		{"matricula", admin.Matricula},
		{"password", admin.Password},
	}

	projection := bson.D{
		{"password", 0},
	}

	if err := collection.FindOne(
		context.TODO(),
		filter,
		options.FindOne().SetProjection(projection),
	).Decode(&findAdmin); err != nil {
		return findAdmin, err
	}

	return findAdmin, nil
}

func getClassID(db *mongo.Client, classData []string, databaseName, collectionName string) (primitive.ObjectID, error) {

	type teste struct {
		ID primitive.ObjectID `bson:"_id,omitempty"`
	}

	var classID teste

	if len(classData) != 3 {
		return classID.ID, errors.New("YEAR/SEASON/CLASSNAME header error")
	}

	year, _ := strconv.Atoi(strings.Trim(classData[0], "\""))
	season, _ := strconv.Atoi(classData[1])
	classname := strings.Trim(classData[2], "\"")

	collection := db.Database(databaseName).Collection(collectionName)

	filter := bson.D{
		{"year", year},
		{"season", season},
		{"classname", classname},
	}

	projection := bson.D{
		{"_id", 1},
	}

	if err := collection.FindOne(
		context.TODO(),
		filter,
		options.FindOne().SetProjection(projection),
	).Decode(&classID); err != nil {
		return classID.ID, err
	}

	return classID.ID, nil
}

func generateRandomPassword() string {

	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	var password string

	for i := 0; i < 3; i++ {
		password += string(letters[rand.Intn(25)])
	}
	for i := 0; i < 3; i++ {
		password += string(numbers[rand.Intn(9)])
	}

	return password

}

func getAdminFromFile(db *mongo.Client, request string) ([]AdminCreate, error) {

	var total []string
	var students []AdminCreate
	var classID primitive.ObjectID
	var err error

	partial := strings.Split(request, ",")
	total = append(total, strings.Split(partial[1], "\n")[1])
	classData := strings.Split(strings.Split(partial[1], "\n")[0], "/")

	if classID, err = getClassID(db, classData, "apc_database", "admin"); err != nil {
		return students, err
	}

	for i := 2; i < len(partial); i++ {
		aux := strings.Split(partial[i], "\n")
		total = append(total, aux[0])
		total = append(total, aux[1])
	}

	total = total[:len(total)-1]

	for i := 0; i < len(total); i += 2 {

		names := strings.SplitAfterN(total[i+1], " ", 2)

		elem := AdminCreate{

			FirstName: strings.Trim(names[0], "\""),
			LastName:  strings.Trim(names[1], "\""),
			Matricula: strings.Trim(total[i], "\""),
			ClassID:   classID,
			Password:  generateRandomPassword(),
		}

		students = append(students, elem)
	}

	return students, nil
}
