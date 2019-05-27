package test

import (
	"plataforma-apc/components/schoolClass"
	"plataforma-apc/components/student"
	"testing"
)

func TestSchoolClass(t *testing.T) {

	studentClass1 := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoUrl:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     8.98,
	}

	studentClass2 := student.Student{
		FirstName: "Vitor",
		LastName:  "Fernandes Dullens",
		Matricula: "160571946",
		Handles:   []string{"vitordullens", "2353251"},
		Password:  "Hgqwge1234",
		PhotoUrl:  "https://userpic.codeforces.com/551311/title/95d04d8b95b95302.jpg",
		Grade:     9.08,
	}

	studentClass3 := student.Student{
		FirstName: "Giovanni",
		LastName:  "Guidini",
		Matricula: "136246666",
		Handles:   []string{"Gguidini", "11165"},
		Password:  "12rw-1234",
		PhotoUrl:  "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Grade:     9.98,
	}

	apc_2018 := schoolClass.SchoolClass{
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		Year:               2019,
		Season:             1,
		Students:           []student.Student{studentClass1, studentClass2, studentClass3},
	}

	if apc_2018.ProfessorFirstName != "Carla" {
		t.Errorf("Invalid professor first name, got: %s, want: %s.", apc_2018.ProfessorFirstName, "Carla")
	}

	if apc_2018.ProfessorLastName != "Castanho" {
		t.Errorf("Invalid professor last name, got: %s, want: %s.", apc_2018.ProfessorLastName, "Castanho")
	}

	if apc_2018.Year != 2019 {
		t.Errorf("Invalid Year, got: %d, want: %d.", apc_2018.Year, 2019)
	}

	if apc_2018.Season != 1 {
		t.Errorf("Invalid Season, got: %d, want: %d.", apc_2018.Season, 1)
	}

	if len(apc_2018.Students) != 3 {
		t.Errorf("Invalid students size, got: %d, want: %d.", len(apc_2018.Students), 3)
	}

	if apc_2018.Students[0].FirstName != "Thiago" {
		t.Errorf("Invalid Student[0] name, got: %s, want: %s.", apc_2018.Students[0].FirstName, "Veras")
	}

	if apc_2018.Students[1].FirstName != "Vitor" {
		t.Errorf("Invalid Student[1] name, got: %s, want: %s.", apc_2018.Students[1].FirstName, "Vitor")
	}

	if apc_2018.Students[2].FirstName != "Giovanni" {
		t.Errorf("Invalid Student[2] name, got: %s, want: %s.", apc_2018.Students[2].FirstName, "Giovanni")
	}

}
