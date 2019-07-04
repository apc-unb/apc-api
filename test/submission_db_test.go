package test

import (
	"context"
	"log"
	"plataforma-apc/components/submission"
	"plataforma-apc/components/student"
	"testing"
	"gopkg.in/mgo.v2/bson"
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

	// Instantiate some submissions objects
	
	student_1 := student.Student{
		ID			: bson.NewObjectId(),
		FirstName	: "Thiago",
		LastName	: "Veras Machado",
		Matricula	: "160156666",
		Handles		: []string{"Veras", "113065"},
		Password	: "HQFnf-1234",
		PhotoUrl	: "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade		: 8.98,
	}

	student_2 := student.Student{
		ID			: bson.NewObjectId(),
		FirstName	: "Vitor",
		LastName	: "Fernandes Dullens",
		Matricula	: "160571946",
		Handles		: []string{"vitordullens", "2353251"},
		Password	: "Hgqwge1234",
		PhotoUrl	: "https://userpic.codeforces.com/551311/title/95d04d8b95b95302.jpg",
		Grade		: 9.08,
	}

	student_3 := student.Student{
		ID			: bson.NewObjectId(),
		FirstName	: "Giovanni",
		LastName	: "Guidini",
		Matricula	: "136246666",
		Handles		: []string{"Gguidini", "11165"},
		Password	: "12rw-1234",
		PhotoUrl	: "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Grade		: 9.98,
	}

	student_4 := student.Student{
		ID			: bson.NewObjectId(),
		FirstName	: "Manoel",
		LastName	: "Josias",
		Matricula	: "135426666",
		Handles		: []string{"Hehe", "11525"},
		Password	: "121521hh1234",
		PhotoUrl	: "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Grade		: 6.27,
	}

	submission_1 := submission.Submission{
		ID			: bson.NewObjectId(),
		Student  	: student_1,
		Veredict 	: "WA",
		Time     	: "19:03:55",
	}

	submission_2 := submission.Submission{
		ID			: bson.NewObjectId(),
		Student  	: student_2,
		Veredict 	: "AC",
		Time     	: "19:07:55",
	}

	submission_3 := submission.Submission{
		ID			: bson.NewObjectId(),
		Student  	: student_3,
		Veredict 	: "TLE",
		Time     	: "19:23:55",
	}

	submission_4 := submission.Submission{
		ID			: bson.NewObjectId(),
		Student  	: student_3,
		Veredict 	: "AC",
		Time     	: "19:24:00",
	}

	submission_5 := submission.Submission{
		ID			: bson.NewObjectId(),
		Student  	: student_4,
		Veredict 	: "AC",
		Time     	: "19:33:55",
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 								 INSERT SUBMISSION DB TEST   						  	 //
	///////////////////////////////////////////////////////////////////////////////////////////
	// Test if submission class array can be inserted in test database
	// Checks if err variable is not null

	if err := submission.CreateSubmissions(db, []submission.Submission{submission_1, submission_2, submission_3, submission_4, submission_5}, "apc_database_test", "submission_test"); err != nil {
		t.Errorf("Failed to insert contest in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							UPDATE LIST OF SUBMISSIONS FROM DB TEST   		      		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	submission_1.Veredict = "CE"
	submission_1.Student = student_2

	submission_2.Student = student_1


	if err := submission.UpdateSubmissions(db, []submission.Submission{submission_1, submission_2}, "apc_database_test", "submission_test"); err != nil {
		t.Errorf("Failed to update submission in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							DELETE LIST OF SUBMISSIONS FROM DB TEST   		       		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	if err := submission.DeleteSubmissions(db, []submission.Submission{submission_3,submission_4}, "apc_database_test", "submission_test"); err != nil {
		t.Errorf("Failed to delete submissions in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							   GET ALL SUBMISSIONS FROM DB TEST 				      		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	var submissions []submission.Submission

	if submissions, err = submission.GetSubmissions(db, "apc_database_test", "submission_test"); err != nil {
		t.Errorf("Failed to get submissions from Database : %s", err)
	}

	if len(submissions) != 3 {
		t.Errorf("Invalid submissions size, got: %d, want: %d.", len(submissions), 3)
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
	
}
