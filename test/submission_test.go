package test

import (
	"plataforma-apc/components/student"
	"plataforma-apc/components/submission"
	"testing"
)

func TestSubmission(t *testing.T) {

	studentClass := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoUrl:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     8.98,
	}

	submissionClass := submission.Submission{
		Student:  studentClass,
		Veredict: "AC",
		Time:     "Jun/04/2019 03:51",
	}

	if submissionClass.Student.FirstName != "Thiago" {
		t.Errorf("Invalid Submission student first name, got: %s, want: %s.", submissionClass.Student.FirstName, "Thiago")
	}

	if submissionClass.Veredict != "AC" {
		t.Errorf("Invalid Submission veredicit, got: %s, want: %s.", submissionClass.Veredict, "AC")
	}

	if submissionClass.Time != "Jun/04/2019 03:51" {
		t.Errorf("Invalid Submission time, got: %s, want: %s.", submissionClass.Time, "Jun/04/2019 03:51")
	}

}
