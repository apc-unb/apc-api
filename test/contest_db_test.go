package test

import (
	"context"
	"log"
	"testing"

	"github.com/plataforma-apc/components/contest"
	"github.com/plataforma-apc/components/schoolClass"
	"github.com/plataforma-apc/components/student"
	"github.com/plataforma-apc/components/submission"
	"github.com/plataforma-apc/components/task"

	"gopkg.in/mgo.v2/bson"
)

func TestContestDB(t *testing.T) {

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

	collection := db.Database("apc_database_test").Collection("contest_test")

	// Drop all content to start testing

	collection.Drop(context.TODO())

	// Instantiate some students objects

	task_1 := task.Task{
		Statement:   "Some 2 números inteiros",
		Score:       2.5,
		Tags:        []string{"String", "Matrix", "Array"},
		Submissions: []submission.Submission{},
	}

	task_2 := task.Task{
		Statement:   "Some 3 números inteiros",
		Score:       4.5,
		Tags:        []string{"Dp", "Array"},
		Submissions: []submission.Submission{},
	}

	task_3 := task.Task{
		Statement:   "Some 5 números inteiros",
		Score:       5.5,
		Tags:        []string{"Sement Tree", "Trie"},
		Submissions: []submission.Submission{},
	}

	task_4 := task.Task{
		Statement:   "Some",
		Score:       1.5,
		Tags:        []string{"Ad Hoc"},
		Submissions: []submission.Submission{},
	}

	student_1 := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoURL:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     8.98,
	}

	student_2 := student.Student{
		FirstName: "Vitor",
		LastName:  "Fernandes Dullens",
		Matricula: "160571946",
		Handles:   []string{"vitordullens", "2353251"},
		Password:  "Hgqwge1234",
		PhotoURL:  "https://userpic.codeforces.com/551311/title/95d04d8b95b95302.jpg",
		Grade:     9.08,
	}

	class_1 := schoolClass.SchoolClass{
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		Year:               2019,
		Season:             1,
		Students:           []student.Student{student_1},
	}

	class_2 := schoolClass.SchoolClass{
		ProfessorFirstName: "Marcos",
		ProfessorLastName:  "Caetano",
		Year:               2018,
		Season:             2,
		Students:           []student.Student{student_2},
	}

	contest_1 := contest.Contest{
		ID:    bson.NewObjectId(),
		Date:  "25/11/1997",
		Class: class_1,
		Tasks: []task.Task{task_1, task_2, task_3},
	}

	contest_2 := contest.Contest{
		ID:    bson.NewObjectId(),
		Date:  "25/11/2018",
		Class: class_2,
		Tasks: []task.Task{task_3, task_4},
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 								 INSERT CONTEST DB TEST 							  	 //
	///////////////////////////////////////////////////////////////////////////////////////////
	// Test if contest class array can be inserted in test database
	// Checks if err variable is not null

	if err := contest.CreateContests(db, []contest.Contest{contest_1, contest_2}, "apc_database_test", "contest_test"); err != nil {
		t.Errorf("Failed to insert contests in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							UPDATE LIST OF CONTEST FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	contest_1.Class.ProfessorFirstName = "Marcos"
	contest_1.Class.ProfessorLastName = "Caetano"

	if err := contest.UpdateContests(db, []contest.Contest{contest_1}, "apc_database_test", "contest_test"); err != nil {
		t.Errorf("Failed to update contests in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							DELETE LIST OF CONTEST FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	if err := contest.DeleteContests(db, []contest.Contest{contest_2}, "apc_database_test", "contest_test"); err != nil {
		t.Errorf("Failed to delete contests in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							   GET ALL CONTEST FROM DB TEST 				      		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	var contests []contest.Contest

	if contests, err = contest.GetContests(db, "apc_database_test", "contest_test"); err != nil {
		t.Errorf("Failed to get contests from Database : %s", err)
	}

	if len(contests) != 1 {
		t.Errorf("Invalid contests size, got: %d, want: %d.", len(contests), 1)
	}

	if contests[0].Class.ProfessorFirstName != "Marcos" {
		t.Errorf("Invalid contest[0] professor first name, got: %s, want: %s.", contests[0].Class.ProfessorFirstName, "Marcos")
	}

	if contests[0].Class.ProfessorLastName != "Caetano" {
		t.Errorf("Invalid contest[0] professor last name, got: %s, want: %s.", contests[0].Class.ProfessorLastName, "Caetano")
	}

	if len(contests[0].Tasks) != 3 {
		t.Errorf("Invalid tasks size from contest, got: %d, want: %d.", len(contests[0].Tasks), 3)
	}

}
