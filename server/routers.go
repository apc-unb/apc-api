package main

import (
	"context"
	"encoding/json"
	"net/http"
	"plataforma-apc/components/schoolClass"
	"plataforma-apc/components/student"
)

///////////////////////////////////////////////////////////////////////////////////////////
// 									 STUDENTS			 								 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createStudents(w http.ResponseWriter, r *http.Request) {

	var students []student.StudentCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&students); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := student.CreateStudents(a.DB, students, "apc_database", "student"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getStudent(w http.ResponseWriter, r *http.Request) {

	var studentLogin student.StudentLogin

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&studentLogin); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if student, err := student.AuthStudent(a.DB, studentLogin, "apc_database", "student"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		respondWithJSON(w, http.StatusOK, student)
	}

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

	var students []student.StudentUpdate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&students); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := student.UpdateStudents(a.DB, students, "apc_database", "student"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deleteStudents(w http.ResponseWriter, r *http.Request) {

	if err := student.DeleteStudents(a.DB, []student.Student{}, "apc_database", "student"); err != nil {
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

	if err := schoolClass.CreateClasses(a.DB, []schoolClass.SchoolClass{}, "apc_database", "schoolClass"); err != nil {
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
