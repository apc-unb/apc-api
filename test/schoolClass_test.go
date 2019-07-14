package test

import (
	"plataforma-apc/components/schoolClass"
	"plataforma-apc/components/student"
	"testing"
)

func TestSchoolClass(t *testing.T) {

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

	student_3 := student.Student{
		FirstName: "Giovanni",
		LastName:  "Guidini",
		Matricula: "136246666",
		Handles:   []string{"Gguidini", "11165"},
		Password:  "12rw-1234",
		PhotoURL:  "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Grade:     9.98,
	}

	class_1 := schoolClass.SchoolClass{
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		Year:               2019,
		Season:             1,
		Students:           []student.Student{student_1, student_2, student_3},
	}

	if class_1.ProfessorFirstName != "Carla" {
		t.Errorf("Invalid professor first name, got: %s, want: %s.", class_1.ProfessorFirstName, "Carla")
	}

	if class_1.ProfessorLastName != "Castanho" {
		t.Errorf("Invalid professor last name, got: %s, want: %s.", class_1.ProfessorLastName, "Castanho")
	}

	if class_1.Year != 2019 {
		t.Errorf("Invalid Year, got: %d, want: %d.", class_1.Year, 2019)
	}

	if class_1.Season != 1 {
		t.Errorf("Invalid Season, got: %d, want: %d.", class_1.Season, 1)
	}

	if len(class_1.Students) != 3 {
		t.Errorf("Invalid students size, got: %d, want: %d.", len(class_1.Students), 3)
	}

	if class_1.Students[0].FirstName != "Thiago" {
		t.Errorf("Invalid Student[0] name, got: %s, want: %s.", class_1.Students[0].FirstName, "Veras")
	}

	if class_1.Students[1].FirstName != "Vitor" {
		t.Errorf("Invalid Student[1] name, got: %s, want: %s.", class_1.Students[1].FirstName, "Vitor")
	}

	if class_1.Students[2].FirstName != "Giovanni" {
		t.Errorf("Invalid Student[2] name, got: %s, want: %s.", class_1.Students[2].FirstName, "Giovanni")
	}

}
