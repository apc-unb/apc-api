package test

import (
	"plataforma-apc/components/student"
	"testing"
)

func TestStudent(t *testing.T) {

	studentClass := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoUrl:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     8.98,
	}

	if studentClass.FirstName != "Thiago" {
		t.Errorf("Invalid first name, got: %s, want: %s.", studentClass.FirstName, "Thiago")
	}

	if studentClass.LastName != "Veras Machado" {
		t.Errorf("Invalid last name, got: %s, want: %s.", studentClass.LastName, "Veras Machado")
	}

	if studentClass.Matricula != "160156666" {
		t.Errorf("Invalid matricula, got: %s, want: %s.", studentClass.Matricula, "160156666")
	}

	if len(studentClass.Handles) != 2 {
		t.Errorf("Invalid handles size, got: %d, want: %d.", len(studentClass.Handles), 2)
	}

	if studentClass.Handles[0] != "Veras" {
		t.Errorf("Invalid handle[0], got: %s, want: %s.", studentClass.Handles[0], "Veras")
	}

	if studentClass.Handles[1] != "113065" {
		t.Errorf("Invalid handle[1], got: %s, want: %s.", studentClass.Handles[1], "113065")
	}

	if studentClass.Password != "HQFnf-1234" {
		t.Errorf("Invalid password, got: %s, want: %s.", studentClass.Password, "HQFnf-1234")
	}

	if studentClass.PhotoUrl != "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg" {
		t.Errorf("Invalid photo url, got: %s, want: %s.", studentClass.PhotoUrl, "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg")
	}

	if studentClass.Grade != 8.98 {
		t.Errorf("Invalid grade, got: %f, want: %f.", studentClass.Grade, 8.98)
	}

}
