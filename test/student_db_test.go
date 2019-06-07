package test

import (
	"context"
	"log"
	"plataforma-apc/components/student"
	"testing"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func GetMongoDB(host, port string) (*mongo.Client, error) {

	db, err := mongo.Connect(context.TODO(), "mongodb://"+host+":"+port)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestStudentDB(t *testing.T) {

	//	Get conection with database
	//	Use config mongo function

	db, err := GetMongoDB("localhost", "27017")

	// Close conection in the end

	defer db.Disconnect(context.TODO())

	// Checks if creating conection with mongo db
	// doesn't return any errors

	if err != nil {
		log.Fatal(err)
	}

	// Get test collection of student

	collection := db.Database("apc_database_test").Collection("student_test")

	// Drop all content to start testing

	collection.Drop(context.TODO())

	// Instantiate some students objects

	studentClass1 := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoUrl:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     8.98,
	}

	studentClass2 := student.Student{
		FirstName: "Vitor",
		LastName:  "Fernandes Dullens",
		Matricula: "160571946",
		Handles:   []string{"vitordullens", "2353251"},
		Password:  "Hgqwge1234",
		PhotoUrl:  "https://userpic.codeforces.com/551311/title/95d04d8b95b95302.jpg",
		Grade:     9.08,
	}

	studentClass3 := student.Student{
		FirstName: "Giovanni",
		LastName:  "Guidini",
		Matricula: "136246666",
		Handles:   []string{"Gguidini", "11165"},
		Password:  "12rw-1234",
		PhotoUrl:  "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Grade:     9.98,
	}

	// 					INSERT TEST
	//
	// Test if student class array can be inserted in test database
	// Checks if err variable is not null, then existis an error

	if err := student.CreateStudent(db, []student.Student{studentClass1, studentClass2, studentClass3}, "apc_database_test", "student_test"); err != nil {
		t.Errorf("Failed to insert students in Database : %s", err)
	}

	// 					SELECT TEST
	//
	// 1º
	//
	// Test if can get all students from database
	// Checks if err variable is not null, then existis an error
	//
	// 2º
	//
	// Test if students array len equals to 3 because INSERT TEST only insert 3 students
	//
	// 3º, 4º, 5º
	//
	// Test if first name of each respective student are correct
	// It is expected that the output is in the same order as the input

	var students []student.Student

	if students, err = student.GetStudents(db, "apc_database_test", "student_test"); err != nil {
		t.Errorf("Failed to get students from Database : %s", err)
	}

	if len(students) != 3 {
		t.Errorf("Invalid students size, got: %d, want: %d.", len(students), 3)
	}

	if students[0].FirstName != "Thiago" {
		t.Errorf("Invalid students[0] first name, got: %s, want: %s.", students[0].FirstName, "Thiago")
	}

	if students[1].FirstName != "Vitor" {
		t.Errorf("Invalid students[1] first name, got: %s, want: %s.", students[1].FirstName, "Vitor")
	}

	if students[2].FirstName != "Giovanni" {
		t.Errorf("Invalid students[2] first name, got: %s, want: %s.", students[2].FirstName, "Giovanni")
	}

}
