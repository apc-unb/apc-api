package test

import (
	"testing"

	"github.com/plataforma-apc/components/student"
	"github.com/plataforma-apc/components/submission"
	"github.com/plataforma-apc/components/task"
)

func TestTask(t *testing.T) {

	student_1 := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoURL:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     8.98,
	}

	submission_1 := submission.Submission{
		Student:  student_1,
		Veredict: "WA",
		Time:     "Jun/04/2019 03:51",
	}

	submission_2 := submission.Submission{
		Student:  student_1,
		Veredict: "TLE",
		Time:     "Jun/04/2019 03:51",
	}

	submission_3 := submission.Submission{
		Student:  student_1,
		Veredict: "AC",
		Time:     "Jun/04/2019 03:51",
	}

	taskClass := task.Task{
		Statement:   "Deivis Express",
		Score:       4.5,
		Tags:        []string{"String", "Matrix", "Array"},
		Submissions: []submission.Submission{submission_1, submission_2, submission_3},
	}

	if taskClass.Statement != "Deivis Express" {
		t.Errorf("Invalid Task statement, got: %s, want: %s.", taskClass.Statement, "Deivis Express")
	}

	if taskClass.Score != 4.5 {
		t.Errorf("Invalid Task score, got: %f, want: %f.", taskClass.Score, 4.5)
	}

	if len(taskClass.Tags) != 3 {
		t.Errorf("Invalid Task tags size, got: %d, want: %d.", len(taskClass.Tags), 3)
	}

	if taskClass.Tags[0] != "String" {
		t.Errorf("Invalid Task tag[0], got: %s, want: %s.", taskClass.Tags[0], "String")
	}

	if taskClass.Tags[1] != "Matrix" {
		t.Errorf("Invalid Task tag[1], got: %s, want: %s.", taskClass.Tags[1], "Matrix")
	}

	if taskClass.Tags[2] != "Array" {
		t.Errorf("Invalid Task tag[2], got: %s, want: %s.", taskClass.Tags[1], "Array")
	}

	if len(taskClass.Submissions) != 3 {
		t.Errorf("Invalid Task submissions size, got: %d, want: %d.", len(taskClass.Submissions), 3)
	}

	if taskClass.Submissions[0].Student.FirstName != "Thiago" {
		t.Errorf("Invalid Task tag[0] Name, got: %s, want: %s.", taskClass.Submissions[0].Student.FirstName, "Thiago")
	}

	if taskClass.Submissions[0].Veredict != "WA" {
		t.Errorf("Invalid Task tag[0] veredict, got: %s, want: %s.", taskClass.Submissions[0].Veredict, "WA")
	}

	if taskClass.Submissions[1].Veredict != "TLE" {
		t.Errorf("Invalid Task tag[1] veredict, got: %s, want: %s.", taskClass.Submissions[1].Veredict, "TLE")
	}

	if taskClass.Submissions[2].Veredict != "AC" {
		t.Errorf("Invalid Task tag[2] veredict, got: %s, want: %s.", taskClass.Submissions[2].Veredict, "AC")
	}

}
