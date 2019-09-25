package test

import (
	"context"
	"log"
	"testing"

	"github.com/VerasThiago/components/student"
)

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

	// Instantiate grades for test
	grades := student.StudentGrades{
		Exams:    []float64{1.4, 2.3, 2.4},
		Projects: []float64{1.6, 3.1, 2.4},
		Lists:    []float64{1.2, 1.2, 1.2},
	}

	// Instantiate some students objects
	student1 := student.StudentCreate{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoURL:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Email:     "teste@gmail.com",
		Grades:    grades,
	}

	student2 := student.StudentCreate{
		FirstName: "Vitor",
		LastName:  "Fernandes Dullens",
		Matricula: "160571946",
		Handles:   []string{"vitordullens", "2353251"},
		Password:  "Hgqwge1234",
		PhotoURL:  "https://userpic.codeforces.com/551311/title/95d04d8b95b95302.jpg",
		Email:     "teste@gmail.com",
		Grades:    grades,
	}

	student3 := student.StudentCreate{
		FirstName: "Giovanni",
		LastName:  "Guidini",
		Matricula: "136246666",
		Handles:   []string{"Gguidini", "11165"},
		Password:  "12rw-1234",
		PhotoURL:  "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Email:     "teste@gmail.com",
		Grades:    grades,
	}
	///////////////////////////////////////////////////////////////////////////////////////////
	// 								 INSERT STUDENTS DB TEST 								 //
	///////////////////////////////////////////////////////////////////////////////////////////
	// Test if student class array can be inserted in test database
	// Checks if err variable is not null

	if err := student.CreateStudents(db, []student.StudentCreate{student1, student2, student3}, "apc_database_test", "student_test"); err != nil {
		t.Errorf("Failed to insert students in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							   GET ALL STUDENTS FROM DB TEST 				      		 //
	///////////////////////////////////////////////////////////////////////////////////////////
	//
	// Test if can get all students from database
	// Checks if err variable is not null
	//

	var students []student.Student

	if students, err = student.GetStudents(db, "apc_database_test", "student_test"); err != nil {
		t.Errorf("Failed to get students from Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							UPDATE LIST OF STUDENTS FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////
	//
	// Test if student class array can be updated in test database
	// Change some students name, then update that class on DB
	// Checks if err variable is not null

	students[0].FirstName = "Guilherme"
	students[0].LastName = "Carvalho"

	students[2].FirstName = "Henrique"
	students[2].LastName = "Machado"

	if err := student.UpdateStudents(db, []student.Student{students[0], students[2]}, "apc_database_test", "student_test"); err != nil {
		t.Errorf("Failed to update students in Database : %s", err)
	}

	students = nil

	if students, err = student.GetStudents(db, "apc_database_test", "student_test"); err != nil {
		t.Errorf("Failed to get students from Database : %s", err)
	}

	if students[0].FirstName != "Guilherme" {
		t.Errorf("Invalid students[0] first name, got: %s, want: %s.", students[0].FirstName, "Thiago")
	}

	if students[2].LastName != "Machado" {
		t.Errorf("Invalid students[2] last name, got: %s, want: %s.", students[2].LastName, "Machado")
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							DELETE LIST OF STUDENTS FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////
	//
	// Test if student class array can be deleted in test database
	// Checks if err variable is not null

	if err := student.DeleteStudents(db, []student.Student{students[0]}, "apc_database_test", "student_test"); err != nil {
		t.Errorf("Failed to delete students in Database : %s", err)
	}

	students = nil

	if students, err = student.GetStudents(db, "apc_database_test", "student_test"); err != nil {
		t.Errorf("Failed to get students from Database : %s", err)
	}

	if len(students) != 2 {
		t.Errorf("Invalid students size, got: %d, want: %d.", len(students), 2)
	}

}
