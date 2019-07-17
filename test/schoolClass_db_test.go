package test

import (
	"context"
	"log"
	"testing"

	"github.com/plataforma-apc/components/schoolClass"
	"github.com/plataforma-apc/components/student"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func GetMongoDB(host, port string) (*mongo.Client, error) {

	db, err := mongo.Connect(context.TODO(), "mongodb://"+host+":"+port)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return db, nil
}

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

	grades := student.StudentGrades{
		Exams:    []float64{1.4, 2.3, 2.4},
		Projects: []float64{1.6, 3.1, 2.4},
		Lists:    []float64{1.2, 1.2, 1.2},
	}

	// Instantiate some school class objects
	student1 := student.Student{
		FirstName: "Thiago",
		LastName:  "Veras Machado",
		Matricula: "160156666",
		Handles:   []string{"Veras", "113065"},
		Password:  "HQFnf-1234",
		PhotoURL:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
		Grade:     grades,
	}

	student2 := student.Student{
		FirstName: "Vitor",
		LastName:  "Fernandes Dullens",
		Matricula: "160571946",
		Handles:   []string{"vitordullens", "2353251"},
		Password:  "Hgqwge1234",
		PhotoURL:  "https://userpic.codeforces.com/551311/title/95d04d8b95b95302.jpg",
		Grade:     grades,
	}

	student3 := student.Student{
		FirstName: "Giovanni",
		LastName:  "Guidini",
		Matricula: "136246666",
		Handles:   []string{"Gguidini", "11165"},
		Password:  "12rw-1234",
		PhotoURL:  "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Grade:     grades,
	}

	student4 := student.Student{
		FirstName: "Manoel",
		LastName:  "Josias",
		Matricula: "135426666",
		Handles:   []string{"Hehe", "11525"},
		Password:  "121521hh1234",
		PhotoURL:  "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
		Grade:     grades,
	}

	class1 := schoolClass.SchoolClassCreate{
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		Year:               2019,
		Season:             1,
		Students:           []student.Student{student1, student2},
	}
	class2 := schoolClass.SchoolClassCreate{
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		Year:               2018,
		Season:             2,
		Students:           []student.Student{student3, student4},
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 								 INSERT CLASS DB TEST 							     	 //
	///////////////////////////////////////////////////////////////////////////////////////////
	// Test if class class array can be inserted in test database
	// Checks if err variable is not null

	if err := schoolClass.CreateClasses(db, []schoolClass.SchoolClassCreate{class1, class2}, "apc_database_test", "schoolClass_test"); err != nil {
		t.Errorf("Failed to insert class in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							   GET ALL CLASSES FROM DB TEST 				      		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	var class []schoolClass.SchoolClass

	if class, err = schoolClass.GetClasses(db, "apc_database_test", "schoolClass_test"); err != nil {
		t.Errorf("Failed to get class from Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							UPDATE LIST OF CLASSES FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	class[0].ProfessorFirstName = "Marcos"
	class[0].ProfessorLastName = "Caetano"

	if err := schoolClass.UpdateClasses(db, []schoolClass.SchoolClass{class[0]}, "apc_database_test", "schoolClass_test"); err != nil {
		t.Errorf("Failed to update class in Database : %s", err)
	}

	class = nil

	if class, err = schoolClass.GetClasses(db, "apc_database_test", "schoolClass_test"); err != nil {
		t.Errorf("Failed to get class from Database : %s", err)
	}

	if class[0].ProfessorFirstName != "Marcos" {
		t.Errorf("Invalid class[0] professor first name, got: %s, want: %s.", class[0].ProfessorFirstName, "Marcos")
	}

	if class[0].ProfessorLastName != "Caetano" {
		t.Errorf("Invalid class[0] professor first name, got: %s, want: %s.", class[2].ProfessorLastName, "Caetano")
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							DELETE LIST OF CLASSES FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	if err := schoolClass.DeleteClasses(db, []schoolClass.SchoolClass{class[0]}, "apc_database_test", "schoolClass_test"); err != nil {
		t.Errorf("Failed to delete class in Database : %s", err)
	}

	class = nil

	if class, err = schoolClass.GetClasses(db, "apc_database_test", "schoolClass_test"); err != nil {
		t.Errorf("Failed to get class from Database : %s", err)
	}

	if len(class) != 1 {
		t.Errorf("Invalid students size, got: %d, want: %d.", len(class), 1)
	}

	if len(class[0].Students) != 2 {
		t.Errorf("Invalid students size from class, got: %d, want: %d.", len(class[0].Students), 2)
	}

}
