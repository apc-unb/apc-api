package test

import (
	"testing"

	"github.com/VerasThiago/plataforma-apc/components/student"
)

func TestStudent(t *testing.T) {

	// Instantiate grades for test
	grades := student.StudentGrades{
		Exams:    []float64{8.98, 2.3, 2.4},
		Projects: []float64{1.6, 3.1, 2.4},
		Lists:    []float64{1.2, 1.2, 1.2},
	}

	class1 := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoURL:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     grades,
	}

	if class1.FirstName != "Thiago" {
		t.Errorf("Invalid first name, got: %s, want: %s.", class1.FirstName, "Thiago")
	}

	if class1.LastName != "Veras Machado" {
		t.Errorf("Invalid last name, got: %s, want: %s.", class1.LastName, "Veras Machado")
	}

	if class1.Matricula != "160156666" {
		t.Errorf("Invalid matricula, got: %s, want: %s.", class1.Matricula, "160156666")
	}

	if len(class1.Handles) != 2 {
		t.Errorf("Invalid handles size, got: %d, want: %d.", len(class1.Handles), 2)
	}

	if class1.Handles[0] != "Veras" {
		t.Errorf("Invalid handle[0], got: %s, want: %s.", class1.Handles[0], "Veras")
	}

	if class1.Handles[1] != "113065" {
		t.Errorf("Invalid handle[1], got: %s, want: %s.", class1.Handles[1], "113065")
	}

	if class1.Password != "HQFnf-1234" {
		t.Errorf("Invalid password, got: %s, want: %s.", class1.Password, "HQFnf-1234")
	}

	if class1.PhotoURL != "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg" {
		t.Errorf("Invalid photo url, got: %s, want: %s.", class1.PhotoURL, "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg")
	}

	if class1.Grade.Exams[0] != 8.98 {
		t.Errorf("Invalid grade, got: %f, want: %f.", class1.Grade.Exams[0], 8.98)
	}

}
