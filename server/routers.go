package main

import (
	"context"
	"net/http"
	"plataforma-apc/components/student"
	"plataforma-apc/components/schoolClass"

	"gopkg.in/mgo.v2/bson"
)

var studentClass1 = student.Student{
	ID:        bson.NewObjectId(),
	FirstName: "Thiago",
	LastName:  "Veras Machado",
	Matricula: "160156666",
	Handles:   []string{"Veras", "113065"},
	Password:  "HQFnf-1234",
	PhotoUrl:  "https://userpic.codeforces.com/546204/title/d2ac05baf39339f.jpg",
	Grade:     8.98,
}

var studentClass2 = student.Student{
	ID:        bson.NewObjectId(),
	FirstName: "Vitor",
	LastName:  "Fernandes Dullens",
	Matricula: "160571946",
	Handles:   []string{"vitordullens", "2353251"},
	Password:  "Hgqwge1234",
	PhotoUrl:  "https://userpic.codeforces.com/551311/title/95d04d8b95b95302.jpg",
	Grade:     9.08,
}

var studentClass3 = student.Student{
	ID:        bson.NewObjectId(),
	FirstName: "Giovanni",
	LastName:  "Guidini",
	Matricula: "136246666",
	Handles:   []string{"Gguidini", "11165"},
	Password:  "12rw-1234",
	PhotoUrl:  "https://userpic.codeforces.com/765049/title/2075d6432eadaae9.jpg",
	Grade:     9.98,
}

var class1 = schoolClass.SchoolClass{

	ProfessorFirstName : "Carla",
	ProfessorLastName  : "Castanho",
	Year               : 2019,
	Season             : 2,
	Students           : []student.Student{studentClass1, studentClass2, studentClass3},

}

///////////////////////////////////////////////////////////////////////////////////////////
// 									 STUDENTS			 								 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createStudents(w http.ResponseWriter, r *http.Request) {

	// Temporary
	collection := a.DB.Database("apc_database").Collection("student")
	collection.Drop(context.TODO())

	if err := student.CreateStudents(a.DB, []student.Student{studentClass1, studentClass2, studentClass3}, "apc_database", "student"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getStudents(w http.ResponseWriter, r *http.Request) {

	students, err := student.GetStudents(a.DB, "apc_database", "student")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, students)
}

func (a *App) updateStudents(w http.ResponseWriter, r *http.Request) {

	studentClass1.FirstName = "Guilherme"
	studentClass1.LastName = "Carvalho"

	studentClass3.FirstName = "Henrique"
	studentClass3.LastName = "Machado"

	if err := student.UpdateStudents(a.DB, []student.Student{studentClass1, studentClass2, studentClass3}, "apc_database", "student"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deleteStudents(w http.ResponseWriter, r *http.Request) {

	if err := student.DeleteStudents(a.DB, []student.Student{studentClass2}, "apc_database", "student"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								   CLASS OF STUDENTS         							 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createClasses(w http.ResponseWriter, r *http.Request) {
	// Temporary
	collection := a.DB.Database("apc_database").Collection("schoolClass")
	collection.Drop(context.TODO())

	if err :=  schoolClass.CreateClasses(a.DB, []schoolClass.SchoolClass{class1}, "apc_database", "schoolClass"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getClasses(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) updateClasses(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) deleteClasses(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								      SUBMISSION		 					     		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createSubmissions(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) getSubmissions(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) updateSubmissions(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) deleteSubmissions(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        TASK		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createTasks(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) getTasks(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) updateTasks(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) deleteTasks(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        CONTEST		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createContests(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) getContests(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) updateContests(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) deleteContests(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}
