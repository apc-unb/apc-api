package test

import (
	"testing"

	"github.com/plataforma-apc/components/contest"
	"github.com/plataforma-apc/components/schoolClass"
	"github.com/plataforma-apc/components/student"
	"github.com/plataforma-apc/components/submission"
	"github.com/plataforma-apc/components/task"
)

func TestContest(t *testing.T) {

	student_1 := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoURL:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     8.98,
	}

	class := schoolClass.SchoolClass{
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		Year:               2019,
		Season:             1,
		Students:           []student.Student{student_1},
	}

	submission_1 := submission.Submission{
		Student:  student_1,
		Veredict: "AC",
		Time:     "Jun/04/2019 03:51",
	}

	task_1 := task.Task{
		Statement:   "Some",
		Score:       1.5,
		Tags:        []string{"Ad Hoc"},
		Submissions: []submission.Submission{submission_1},
	}

	contest := contest.Contest{
		Date:  "25/11/1997",
		Class: class,
		Tasks: []task.Task{task_1},
	}

	if contest.Date != "25/11/1997" {
		t.Errorf("Invalid contest date, got: %s, want: %s.", contest.Date, "25/11/1997")
	}

	if contest.Class.ProfessorFirstName != "Carla" {
		t.Errorf("Invalid submission class professor first name, got: %s, want: %s.", contest.Class.ProfessorFirstName, "Carla")
	}

	if contest.Class.Students[0].FirstName != "Thiago" {
		t.Errorf("Invalid submission class student 0 first name, got: %s, want: %s.", contest.Class.Students[0].FirstName, "Thiago")
	}

	if contest.Tasks[0].Statement != "Some" {
		t.Errorf("Invalid contest task 0 statement, got: %s, want: %s.", contest.Tasks[0].Statement, "Some")
	}

	if contest.Tasks[0].Submissions[0].Veredict != "AC" {
		t.Errorf("Invalid contest task 0 submission 0 veredict, got: %s, want: %s.", contest.Tasks[0].Submissions[0].Veredict, "AC")
	}

}
