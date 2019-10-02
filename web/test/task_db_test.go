package test

import (
	"context"
	"log"
	"testing"

	"github.com/components/submission"
	"github.com/components/task"
)

func TestTaskDB(t *testing.T) {

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
	collection := db.Database("apc_database_test").Collection("task_test")

	// Drop all content to start testing
	collection.Drop(context.TODO())

	// Instantiate some tasks objects
	task1 := task.TaskCreate{
		Statement:   "Some 2 números inteiros",
		Score:       2.5,
		Tags:        []string{"String", "Matrix", "Array"},
		Submissions: []submission.Submission{},
	}

	task2 := task.TaskCreate{
		Statement:   "Some 3 números inteiros",
		Score:       4.5,
		Tags:        []string{"Dp", "Array"},
		Submissions: []submission.Submission{},
	}

	task3 := task.TaskCreate{
		Statement:   "Some 4 números inteiros",
		Score:       5.5,
		Tags:        []string{"Sement Tree", "Trie"},
		Submissions: []submission.Submission{},
	}

	task4 := task.TaskCreate{
		Statement:   "Some",
		Score:       1.5,
		Tags:        []string{"Ad Hoc"},
		Submissions: []submission.Submission{},
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 								 INSERT TASKS DB TEST 								     //
	///////////////////////////////////////////////////////////////////////////////////////////
	//
	// Test if task class array can be inserted in test database
	// Checks if err variable is not null

	if err := task.CreateTasks(db, []task.TaskCreate{task1, task2, task3, task4}, "apc_database_test", "task_test"); err != nil {
		t.Errorf("Failed to insert tasks in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							   GET ALL TASKS FROM DB TEST 				      		 //
	///////////////////////////////////////////////////////////////////////////////////////////
	// 1º
	//
	// Test if can get all tasks from database
	// Checks if err variable is not null
	//
	// 2º
	//
	// Test if tasks array len equals to 3 because INSERT TEST inserted 4 tasks but DELETE test deleted 1 task, remaining only 3 tasks
	//
	// 3º Check some tasks tags
	// It is expected that the output is in the same order as the input

	var tasks []task.Task

	if tasks, err = task.GetTasks(db, "apc_database_test", "task_test"); err != nil {
		t.Errorf("Failed to get tasks from Database : %s", err)
	}

	if tasks[0].Score != 2.5 {
		t.Errorf("Invalid tasks[0].Score, got: %f, want: %f.", tasks[1].Score, 2.5)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							UPDATE LIST OF TASKS FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////
	//
	// Test if task class array can be updated in test database
	// Change some task name, then update that class on DB
	// Checks if err variable is not null

	tasks[1].Tags[0] = "Recursion"
	tasks[1].Tags[1] = "Matrix"

	if err := task.UpdateTasks(db, []task.Task{tasks[1]}, "apc_database_test", "task_test"); err != nil {
		t.Errorf("Failed to update tasks in Database : %s", err)
	}

	tasks = nil

	if tasks, err = task.GetTasks(db, "apc_database_test", "task_test"); err != nil {
		t.Errorf("Failed to get tasks from Database : %s", err)
	}

	if tasks[1].Tags[0] != "Recursion" {
		t.Errorf("Invalid tasks[1].Tags[0], got: %s, want: %s.", tasks[1].Tags[0], "Recursion")
	}

	if tasks[1].Tags[1] != "Matrix" {
		t.Errorf("Invalid tasks[1].Tags[1], got: %s, want: %s.", tasks[1].Tags[1], "Matrix")
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							DELETE LIST OF TASKS FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////
	//
	// Test if task class array can be deleted in test database
	// Checks if err variable is not null

	if err := task.DeleteTasks(db, []task.Task{tasks[2]}, "apc_database_test", "task_test"); err != nil {
		t.Errorf("Failed to delete task in Database : %s", err)
	}

	tasks = nil

	if tasks, err = task.GetTasks(db, "apc_database_test", "task_test"); err != nil {
		t.Errorf("Failed to get tasks from Database : %s", err)
	}

	if len(tasks) != 3 {
		t.Errorf("Invalid tasks size, got: %d, want: %d.", len(tasks), 3)
	}

}
