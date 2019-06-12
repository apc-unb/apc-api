package test

import (
	"context"
	"log"
	"plataforma-apc/components/schoolClass"
	"plataforma-apc/components/student"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestClassDB(t *testing.T) {

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

	collection := db.Database("apc_database_test").Collection("schoolClass_test")

	// Drop all content to start testing

	collection.Drop(context.TODO())

	// Instantiate some students objects

	student_1 := student.Student{
		ID:        bson.NewObjectId(),
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoUrl:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     8.98,
	}

	student_2 := student.Student{
		ID:        bson.NewObjectId(),
		FirstName: "Vitor",
		LastName:  "Fernandes Dullens",
		Matricula: "160571946",
		Handles:   []string{"vitordullens", "2353251"},
		Password:  "Hgqwge1234",
		PhotoUrl:  "https://userpic.codeforces.com/551311/title/95d04d8b95b95302.jpg",
		Grade:     9.08,
	}

	student_3 := student.Student{
		ID:        bson.NewObjectId(),
		FirstName: "Giovanni",
		LastName:  "Guidini",
		Matricula: "136246666",
		Handles:   []string{"Gguidini", "11165"},
		Password:  "12rw-1234",
		PhotoUrl:  "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Grade:     9.98,
	}

	student_4 := student.Student{
		ID:        bson.NewObjectId(),
		FirstName: "Manoel",
		LastName:  "Josias",
		Matricula: "135426666",
		Handles:   []string{"Hehe", "11525"},
		Password:  "121521hh1234",
		PhotoUrl:  "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Grade:     6.27,
	}

	apc2019_1 := schoolClass.SchoolClass{
		ID:                 bson.NewObjectId(),
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		Year:               2019,
		Season:             1,
		Students:           []student.Student{student_1, student_2},
	}
	apc2018_2 := schoolClass.SchoolClass{
		ID:                 bson.NewObjectId(),
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		Year:               2018,
		Season:             2,
		Students:           []student.Student{student_3, student_4},
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 								 INSERT CLASS DB TEST 							     	 //
	///////////////////////////////////////////////////////////////////////////////////////////
	// Test if class class array can be inserted in test database
	// Checks if err variable is not null

	if err := schoolClass.CreateClasses(db, []schoolClass.SchoolClass{apc2019_1, apc2018_2}, "apc_database_test", "schoolClass_test"); err != nil {
		t.Errorf("Failed to insert class in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							UPDATE LIST OF CLASSES FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							DELETE LIST OF CLASSES FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							   GET ALL CLASSES FROM DB TEST 				      		 //
	///////////////////////////////////////////////////////////////////////////////////////////

}
