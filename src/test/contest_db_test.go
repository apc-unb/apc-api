package test

import (
	"context"
	"log"
	"testing"

	"github.com/VerasThiago/plataforma-apc/components/contest"
	"github.com/VerasThiago/plataforma-apc/components/schoolClass"
	"github.com/VerasThiago/plataforma-apc/components/student"
	"github.com/VerasThiago/plataforma-apc/components/submission"
	"github.com/VerasThiago/plataforma-apc/components/task"
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
	task1 := task.Task{
		Statement:   "Some 2 números inteiros",
		Score:       2.5,
		Tags:        []string{"String", "Matrix", "Array"},
		Submissions: []submission.Submission{},
	}

	task2 := task.Task{
		Statement:   "Some 3 números inteiros",
		Score:       4.5,
		Tags:        []string{"Dp", "Array"},
		Submissions: []submission.Submission{},
	}

	task3 := task.Task{
		Statement:   "Some 5 números inteiros",
		Score:       5.5,
		Tags:        []string{"Sement Tree", "Trie"},
		Submissions: []submission.Submission{},
	}

	task4 := task.Task{
		Statement:   "Some",
		Score:       1.5,
		Tags:        []string{"Ad Hoc"},
		Submissions: []submission.Submission{},
	}

	class1 := schoolClass.SchoolClass{
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		Year:               2019,
		Season:             1,
		Students:           []student.Student{},
	}

	class2 := schoolClass.SchoolClass{
		ProfessorFirstName: "Marcos",
		ProfessorLastName:  "Caetano",
		Year:               2018,
		Season:             2,
		Students:           []student.Student{},
	}

	contest1 := contest.ContestCreate{
		Date:  "25/11/1997",
		Class: class1,
		Tasks: []task.Task{task1, task2},
	}

	contest2 := contest.ContestCreate{
		Date:  "25/11/2018",
		Class: class2,
		Tasks: []task.Task{task1, task2, task3, task4},
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 								 INSERT CONTEST DB TEST 							  	 //
	///////////////////////////////////////////////////////////////////////////////////////////
	// Test if contest class array can be inserted in test database
	// Checks if err variable is not null

	if err := contest.CreateContests(db, []contest.ContestCreate{contest1, contest2}, "apc_database_test", "contest_test"); err != nil {
		t.Errorf("Failed to insert contests in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							   GET ALL CONTEST FROM DB TEST 				      		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	var contests []contest.Contest

	if contests, err = contest.GetContests(db, "apc_database_test", "contest_test"); err != nil {
		t.Errorf("Failed to get contests from Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							UPDATE LIST OF CONTEST FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	contests[0].Class.ProfessorFirstName = "Marcos"
	contests[0].Class.ProfessorLastName = "Caetano"

	if err := contest.UpdateContests(db, []contest.Contest{contests[0]}, "apc_database_test", "contest_test"); err != nil {
		t.Errorf("Failed to update contests in Database : %s", err)
	}

	contests = nil

	if contests, err = contest.GetContests(db, "apc_database_test", "contest_test"); err != nil {
		t.Errorf("Failed to get contests from Database : %s", err)
	}

	if contests[0].Class.ProfessorFirstName != "Marcos" {
		t.Errorf("Invalid contest[0] professor first name, got: %s, want: %s.", contests[0].Class.ProfessorFirstName, "Marcos")
	}

	if contests[0].Class.ProfessorLastName != "Caetano" {
		t.Errorf("Invalid contest[0] professor last name, got: %s, want: %s.", contests[0].Class.ProfessorLastName, "Caetano")
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							DELETE LIST OF CONTEST FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	if err := contest.DeleteContests(db, []contest.Contest{contests[0]}, "apc_database_test", "contest_test"); err != nil {
		t.Errorf("Failed to delete contests in Database : %s", err)
	}

	contests = nil

	if contests, err = contest.GetContests(db, "apc_database_test", "contest_test"); err != nil {
		t.Errorf("Failed to get contests from Database : %s", err)
	}

	if len(contests) != 1 {
		t.Errorf("Invalid contests size, got: %d, want: %d.", len(contests), 1)
	}

	if len(contests[0].Tasks) != 4 {
		t.Errorf("Invalid tasks size from contest, got: %d, want: %d.", len(contests[0].Tasks), 3)
	}

}
