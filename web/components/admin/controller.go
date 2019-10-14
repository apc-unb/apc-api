package admin

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"strings"

	"github.com/apc-unb/apc-api/web/components/user"
	"github.com/apc-unb/apc-api/web/utils"
	"github.com/togatoga/goforces"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateAdmin receive  a list of students
// Checks if that list is not null (can't insert null list)
// Insert each student individually in database
// @param	db				pointer to database
// @param	students 		list of students
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	error 			function error
// TODO : Insert all students at the same time (if possible)
func CreateAdmin(db *mongo.Client, api *goforces.Client, admins []AdminCreate, databaseName, collectionName string) ([]user.UserCredentials, error) {

	var studentsReturn []user.UserCredentials
	var singleAdmin user.UserCredentials
	var mongoReturn *mongo.InsertOneResult
	var err error

	if len(admins) == 0 {
		return nil, nil
	}

	collection := db.Database(databaseName).Collection(collectionName)
	collectionLogin := db.Database(databaseName).Collection(collectionName + "_login")

	for _, admin := range admins {

		pwd := generateRandomPassword()

		if singleAdmin.Password, err = utils.HashAndSalt([]byte(pwd)); err != nil {
			return nil, err
		}

		if mongoReturn, err = collection.InsertOne(context.TODO(), admin); err != nil {
			return studentsReturn, err
		} else {
			singleAdmin.ID = mongoReturn.InsertedID.(primitive.ObjectID)
		}

		singleAdmin.Matricula = admin.Matricula

		if _, err = collectionLogin.InsertOne(context.TODO(), singleAdmin); err != nil {
			return nil, err
		}

		studentsReturn = append(studentsReturn, singleAdmin)

	}

	return studentsReturn, nil
}

// CreateAdminFile receive a csv file in that current format : https://github.com/apc-unb/tree/master/components/student
// calls getAdminFromFile() that return list of AdminCreate and insert into db
// @param	db				pointer to database
// @param   request         all data to be parsed
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	error 			function error
func CreateAdminFile(db *mongo.Client, request, databaseName, collectionName string) ([]user.UserCredentials, error) {

	var admins []AdminCreate
	var studentsReturn []user.UserCredentials
	var singleAdmin user.UserCredentials
	var mongoReturn *mongo.InsertOneResult
	var err error

	if admins, err = getAdminFromFile(db, request); err != nil {
		return nil, err
	}

	if len(admins) == 0 {
		return nil, nil
	}

	collection := db.Database(databaseName).Collection(collectionName)
	collectionLogin := db.Database(databaseName).Collection(collectionName + "_login")

	for _, admin := range admins {

		pwd := generateRandomPassword()

		if singleAdmin.Password, err = utils.HashAndSalt([]byte(pwd)); err != nil {
			return nil, err
		}

		if mongoReturn, err = collection.InsertOne(context.TODO(), admin); err != nil {
			return studentsReturn, err
		} else {
			singleAdmin.ID = mongoReturn.InsertedID.(primitive.ObjectID)
		}

		singleAdmin.Matricula = admin.Matricula

		if _, err = collectionLogin.InsertOne(context.TODO(), singleAdmin); err != nil {
			return nil, err
		}

		studentsReturn = append(studentsReturn, singleAdmin)

	}

	return studentsReturn, nil
}

// GetAdmins return list of all students from Database
// Get all students at the same time and store inside cursor
// Decode each student inside student class and append into students array
// @param	db				pointer to database
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	[]Student		list of all students
// @return 	error 			function error
func GetAdmins(db *mongo.Client, databaseName, collectionName string) ([]AdminInfo, error) {

	collection := db.Database(databaseName).Collection(collectionName)

	admins := []AdminInfo{}

	cursor, err := collection.Find(
		context.TODO(),
		bson.D{{}},
		options.Find(),
	)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {

		var elem AdminInfo

		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		admins = append(admins, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(context.TODO())

	return admins, nil
}

// UpdateAdmins receive admin (updated)
// Checks if admin old password matches with db to update that admin password or email
// @param	db				pointer to database (updated)
// @param	api 			codeforces api
// @param	admin 			list of admins
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	error 			function error
// TODO : Update all students at the same time (if possible)
func UpdateAdmins(db *mongo.Client, api *goforces.Client, admin AdminUpdate, databaseName, collectionName string) error {

	collection := db.Database(databaseName).Collection(collectionName)
	collectionLogin := db.Database(databaseName).Collection(collectionName + "_login")

	var err error
	adminData := user.UserCredentials{}

	filter := bson.M{
		"_id": admin.ID,
	}

	if err := collectionLogin.FindOne(
		context.TODO(),
		filter,
		options.FindOne(),
	).Decode(&adminData); err != nil {
		return err
	}

	if err = utils.ComparePasswords(adminData.Password, admin.Password); err != nil {
		return errors.New("mongo: no documents in result")
	}

	if admin.NewPassword != "" {
		if admin.NewPassword, err = utils.HashAndSalt([]byte(admin.NewPassword)); err != nil {
			return err
		}

		updateSet := bson.M{"$set": bson.M{"password": admin.NewPassword}}

		if _, err := collectionLogin.UpdateOne(context.TODO(), filter, updateSet, nil); err != nil {
			return err
		}

		return nil
	}

	filter = bson.M{
		"_id": admin.ID,
	}

	update := bson.M{}

	if admin.Email != "" {
		update["email"] = admin.Email
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

// UpdateAdminStudent receive update stundets data, receive a student (updated)
// @param	db				pointer to database (updated)
// @param	api				codeforces api
// @param	admin 			student to be updated
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	error 			function error
func UpdateAdminStudent(db *mongo.Client, api *goforces.Client, admin AdminUpdateStudent, databaseName, collectionName string) error {

	collection := db.Database(databaseName).Collection(collectionName)

	filter := bson.M{
		"_id": admin.StudentID,
	}

	update := bson.M{}

	if !admin.ClassID.IsZero() {
		update["classid"] = admin.ClassID
	}

	if admin.FirstName != "" {
		update["firstname"] = admin.FirstName
	}

	if admin.LastName != "" {
		update["lastname"] = admin.LastName
	}

	if admin.Matricula != "" {
		update["matricula"] = admin.Matricula
	}

	if admin.Handles.Codeforces != "" {
		update["handles.codeforces"] = admin.Handles.Codeforces
	}

	if admin.Handles.Uri != "" {
		update["handles.uri"] = admin.Handles.Uri
	}

	if admin.PhotoURL != "" {
		update["photourl"] = admin.PhotoURL
	}

	if admin.Email != "" {
		update["email"] = admin.Email
	}

	if len(admin.Grades.Exams) > 0 {
		update["grades.exams"] = admin.Grades.Exams
	}

	if len(admin.Grades.Lists) > 0 {
		update["grades.lists"] = admin.Grades.Lists
	}

	updateSet := bson.M{"$set": update}

	if _, err := collection.UpdateOne(context.TODO(), filter, updateSet, nil); err != nil {
		return err
	}

	return nil
}

// DeleteAdminStudents recieve a list of students (to be deleted)
// Checks if that list is not null (can't delete null list)
// Delete each student individually
// @param	db				pointer to database (to be deleted)
// @param	admins	 		list of students
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	error 			function error
// TODO : Delete all students at the same time (if possible)
func DeleteAdminStudents(db *mongo.Client, admins []Admin, databaseName, collectionName string) error {

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

// AuthAdmin recieve an admin (to be authenticated)
// Checks if that login and password exist in database
// @param	db				pointer to database
// @param	admin			admin matricula and password
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	AdminInfo		json if exist plus all admin data
// @return 	error 			function error
func AuthAdmin(db *mongo.Client, admin user.UserCredentials, databaseName, collectionName string) (AdminInfo, error) {

	var err error

	collection := db.Database(databaseName).Collection(collectionName)
	collectionLogin := db.Database(databaseName).Collection(collectionName + "_login")

	findAdmin := AdminInfo{}
	adminData := user.UserCredentials{}

	filter := bson.D{
		{"matricula", admin.Matricula},
	}

	if err := collectionLogin.FindOne(
		context.TODO(),
		filter,
		options.FindOne(),
	).Decode(&adminData); err != nil {
		return findAdmin, err
	}

	if err = utils.ComparePasswords(adminData.Password, admin.Password); err != nil {
		return findAdmin, errors.New("mongo: no documents in result")
	}

	if err := collection.FindOne(
		context.TODO(),
		filter,
		options.FindOne(),
	).Decode(&findAdmin); err != nil {
		return findAdmin, err
	}

	return findAdmin, nil
}

// getClassID receive year, season and class name then return id of that current class
// @param	db				pointer to database
// @param	classData		year, season and class name
// @param	databaseName	name of database
// @param	collectionName	name of collection
// @return 	ObjectID		class ID
// @return 	error 			function error
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

// generateRandomPassword generate a random password using Pimenta Judge style
// @return 	string		random password
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

// getAdminFromFile receive a string and return list of admins data
// @param	db				pointer to database
// @param	request			all the csv file to be converted
// @return 	AdminCreate		list of admins
// @return 	error 			function error
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
			//Password:  generateRandomPassword(),
		}

		students = append(students, elem)
	}

	return students, nil
}
