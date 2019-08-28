package test

import (
	"testing"

	"github.com/VerasThiago/plataforma-apc/components/student"
	"github.com/VerasThiago/plataforma-apc/components/submission"
)

func TestSubmission(t *testing.T) {

	// Instantiate grades for test
	grades := student.StudentGrades{
		Exams:    []float64{8.98, 2.3, 2.4},
		Projects: []float64{1.6, 3.1, 2.4},
		Lists:    []float64{1.2, 1.2, 1.2},
	}

	student1 := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoURL:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     grades,
	}

	class1 := submission.Submission{
		Student:  student1,
		Veredict: "AC",
		Time:     "Jun/04/2019 03:51",
	}

	if class1.Student.FirstName != "Thiago" {
		t.Errorf("Invalid Submission student first name, got: %s, want: %s.", class1.Student.FirstName, "Thiago")
	}

	if class1.Veredict != "AC" {
		t.Errorf("Invalid Submission veredicit, got: %s, want: %s.", class1.Veredict, "AC")
	}

	if class1.Time != "Jun/04/2019 03:51" {
		t.Errorf("Invalid Submission time, got: %s, want: %s.", class1.Time, "Jun/04/2019 03:51")
	}

}
