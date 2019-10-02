package test

import (
	"context"
	"log"
	"testing"

	"github.com/components/student"
	"github.com/components/submission"
)

func TestSubmissionDB(t *testing.T) {

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
	collection := db.Database("apc_database_test").Collection("submission_test")

	// Drop all content to start testing
	collection.Drop(context.TODO())

	// Instantiate grades for test
	grades := student.StudentGrades{
		Exams:    []float64{8.98, 2.3, 2.4},
		Projects: []float64{1.6, 3.1, 2.4},
		Lists:    []float64{1.2, 1.2, 1.2},
	}

	// Instantiate some submissions objects
	student1 := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoURL:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     grades,
	}

	student2 := student.Student{
		FirstName: "Vitor",
		LastName:  "Fernandes Dullens",
		Matricula: "160571946",
		Handles:   []string{"vitordullens", "2353251"},
		Password:  "Hgqwge1234",
		PhotoURL:  "https://userpic.codeforces.com/551311/title/95d04d8b95b95302.jpg",
		Grade:     grades,
	}

	student3 := student.Student{
		FirstName: "Giovanni",
		LastName:  "Guidini",
		Matricula: "136246666",
		Handles:   []string{"Gguidini", "11165"},
		Password:  "12rw-1234",
		PhotoURL:  "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Grade:     grades,
	}

	student4 := student.Student{
		FirstName: "Manoel",
		LastName:  "Josias",
		Matricula: "135426666",
		Handles:   []string{"Hehe", "11525"},
		Password:  "121521hh1234",
		PhotoURL:  "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Grade:     grades,
	}

	submission1 := submission.SubmissionCreate{
		Student:  student1,
		Veredict: "WA",
		Time:     "19:03:55",
	}

	submission2 := submission.SubmissionCreate{
		Student:  student2,
		Veredict: "AC",
		Time:     "19:07:55",
	}

	submission3 := submission.SubmissionCreate{
		Student:  student3,
		Veredict: "TLE",
		Time:     "19:23:55",
	}

	submission4 := submission.SubmissionCreate{
		Student:  student3,
		Veredict: "AC",
		Time:     "19:24:00",
	}

	submission5 := submission.SubmissionCreate{
		Student:  student4,
		Veredict: "AC",
		Time:     "19:33:55",
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 								 INSERT SUBMISSION DB TEST   						  	 //
	///////////////////////////////////////////////////////////////////////////////////////////
	// Test if submission class array can be inserted in test database
	// Checks if err variable is not null

	if err := submission.CreateSubmissions(db, []submission.SubmissionCreate{submission1, submission2, submission3, submission4, submission5}, "apc_database_test", "submission_test"); err != nil {
		t.Errorf("Failed to insert contest in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							   GET ALL SUBMISSIONS FROM DB TEST 				      		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	var submissions []submission.Submission

	if submissions, err = submission.GetSubmissions(db, "apc_database_test", "submission_test"); err != nil {
		t.Errorf("Failed to get submissions from Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							UPDATE LIST OF SUBMISSIONS FROM DB TEST   		      		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	submissions[0].Veredict = "CE"
	submissions[0].Student.FirstName, submissions[1].Student.FirstName = submissions[1].Student.FirstName, submissions[0].Student.FirstName

	if err := submission.UpdateSubmissions(db, []submission.Submission{submissions[0], submissions[1]}, "apc_database_test", "submission_test"); err != nil {
		t.Errorf("Failed to update submission in Database : %s", err)
	}

	submissions = nil

	if submissions, err = submission.GetSubmissions(db, "apc_database_test", "submission_test"); err != nil {
		t.Errorf("Failed to get submissions from Database : %s", err)
	}

	if submissions[0].Veredict != "CE" {
		t.Errorf("Invalid submissions[0] veredict, got: %s, want: %s.", submissions[0].Veredict, "CE")
	}

	if submissions[0].Student.FirstName != "Vitor" {
		t.Errorf("Invalid submissions[0] student name, got: %s, want: %s.", submissions[0].Student.FirstName, "Vitor")
	}

	if submissions[1].Student.FirstName != "Thiago" {
		t.Errorf("Invalid submissions[1] student name, got: %s, want: %s.", submissions[1].Student.FirstName, "Thiago")
	}

	if submissions[1].Veredict != "AC" {
		t.Errorf("Invalid submissions[0] veredict, got: %s, want: %s.", submissions[0].Veredict, "AC")
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							DELETE LIST OF SUBMISSIONS FROM DB TEST   		       		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	if err := submission.DeleteSubmissions(db, []submission.Submission{submissions[2], submissions[3]}, "apc_database_test", "submission_test"); err != nil {
		t.Errorf("Failed to delete submissions in Database : %s", err)
	}

	submissions = nil

	if submissions, err = submission.GetSubmissions(db, "apc_database_test", "submission_test"); err != nil {
		t.Errorf("Failed to get submissions from Database : %s", err)
	}

	if len(submissions) != 3 {
		t.Errorf("Invalid submissions size, got: %d, want: %d.", len(submissions), 3)
	}

}
