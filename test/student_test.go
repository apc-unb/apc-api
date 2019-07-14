package test

import (
	"plataforma-apc/components/student"
	"testing"
)

func TestStudent(t *testing.T) {

	class_1 := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoURL:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     8.98,
	}

	if class_1.FirstName != "Thiago" {
		t.Errorf("Invalid first name, got: %s, want: %s.", class_1.FirstName, "Thiago")
	}

	if class_1.LastName != "Veras Machado" {
		t.Errorf("Invalid last name, got: %s, want: %s.", class_1.LastName, "Veras Machado")
	}

	if class_1.Matricula != "160156666" {
		t.Errorf("Invalid matricula, got: %s, want: %s.", class_1.Matricula, "160156666")
	}

	if len(class_1.Handles) != 2 {
		t.Errorf("Invalid handles size, got: %d, want: %d.", len(class_1.Handles), 2)
	}

	if class_1.Handles[0] != "Veras" {
		t.Errorf("Invalid handle[0], got: %s, want: %s.", class_1.Handles[0], "Veras")
	}

	if class_1.Handles[1] != "113065" {
		t.Errorf("Invalid handle[1], got: %s, want: %s.", class_1.Handles[1], "113065")
	}

	if class_1.Password != "HQFnf-1234" {
		t.Errorf("Invalid password, got: %s, want: %s.", class_1.Password, "HQFnf-1234")
	}

	if class_1.PhotoURL != "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg" {
		t.Errorf("Invalid photo url, got: %s, want: %s.", class_1.PhotoURL, "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg")
	}

	if class_1.Grade != 8.98 {
		t.Errorf("Invalid grade, got: %f, want: %f.", class_1.Grade, 8.98)
	}

}
