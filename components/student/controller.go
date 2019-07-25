package student

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
func CreateStudents(db *mongo.Client, api *goforces.Client, students []StudentCreate, databaseName, collectionName string) error {

	if len(students) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

	for _, student := range students {

		student.PhotoURL = getCodeforcesAvatarURL(student.Handles.Codeforces, api)

		if _, err := collection.InsertOne(context.TODO(), student); err != nil {
			return err
		}
	}

	return nil
}

func CreateStudentsFile(db *mongo.Client, request string, databaseName, collectionName string) error {

	var students []StudentCreate
	var err error

	if students, err = getStudentsFromFile(db, request); err != nil {
		return err
	}

	if len(students) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

	for _, student := range students {
		if _, err := collection.InsertOne(context.TODO(), student); err != nil {
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
func GetStudents(db *mongo.Client, databaseName, collectionName string) ([]StudentInfo, error) {

	collection := db.Database(databaseName).Collection(collectionName)

	// Here's an array in which you can store the decoded documents
	students := []StudentInfo{}

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
		var elem StudentInfo

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

// UpdateStudents recieve student (updated)
// Checks if student old password matches with db to update that student password or email
// @param	db				pointer to database (updated)
// @param	students 		list of students
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	StudentUpdate	student new data
// @return 	error 			function error
// TODO : Update all students at the same time (if possible)
func UpdateStudents(db *mongo.Client, student StudentUpdate, databaseName, collectionName string) error {

	collection := db.Database(databaseName).Collection(collectionName)

	currentStudent := Student{}

	filter := bson.M{
		"_id":      student.ID,
		"password": student.Password,
	}

	projection := bson.M{
		"_id": 1,
	}

	if err := collection.FindOne(
		context.TODO(),
		filter,
		options.FindOne().SetProjection(projection),
	).Decode(&currentStudent); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("Invalid password")
		}
		return err
	}

	filter = bson.M{
		"_id": student.ID,
	}

	update := bson.M{}

	if student.Email != "" {
		update["email"] = student.Email
	}

	if student.NewPassword != "" {
		update["password"] = student.NewPassword
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
func DeleteStudents(db *mongo.Client, students []Student, databaseName, collectionName string) error {

	if len(students) == 0 {
		return nil
	}

	collection := db.Database(databaseName).Collection(collectionName)

	for _, student := range students {
		filter := bson.M{"_id": student.ID}
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
func AuthStudent(db *mongo.Client, student StudentLogin, databaseName, collectionName string) (StudentInfo, error) {

	collection := db.Database(databaseName).Collection(collectionName)

	findStudent := StudentInfo{}

	filter := bson.D{
		{"matricula", student.Matricula},
		{"password", student.Password},
	}

	projection := bson.D{
		{"password", 0},
	}

	if err := collection.FindOne(
		context.TODO(),
		filter,
		options.FindOne().SetProjection(projection),
	).Decode(&findStudent); err != nil {
		return findStudent, err
	}

	return findStudent, nil
}

// getCodeforcesAvatarURL recieve handle string
// Return handle avatar url if exist
// @param	handle			student codeforces handle
// @param	api				pointer to goforces client
// @return 	string 			avatar url
func getCodeforcesAvatarURL(handle string, api *goforces.Client) string {

	ctx := context.Background()

	var userAvatarURL string

	if handlesArray, err := api.GetUserInfo(ctx, []string{handle}); err == nil {
		userAvatarURL = "https:" + handlesArray[0].Avatar
	}

	return userAvatarURL
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

func getStudentsFromFile(db *mongo.Client, request string) ([]StudentCreate, error) {

	var total []string
	var students []StudentCreate
	var classID primitive.ObjectID
	var err error

	partial := strings.Split(request, ",")
	total = append(total, strings.Split(partial[1], "\n")[1])
	classData := strings.Split(strings.Split(partial[1], "\n")[0], "/")

	if classID, err = getClassID(db, classData, "apc_database", "schoolClass"); err != nil {
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

		elem := StudentCreate{

			FirstName: names[0],
			LastName:  names[1],
			Matricula: strings.Trim(total[i], "\""),
			ClassID:   classID,
			Password:  generateRandomPassword(),
		}

		students = append(students, elem)
	}

	return students, nil
}
