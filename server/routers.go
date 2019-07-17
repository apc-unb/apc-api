package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/plataforma-apc/components/news"
	"github.com/plataforma-apc/components/schoolClass"
	"github.com/plataforma-apc/components/student"
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
	var singleStudent student.StudentInfo
	var newsArray []news.News
	var err error

	decoder := json.NewDecoder(r.Body)

	if err = decoder.Decode(&studentLogin); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if singleStudent, err = student.AuthStudent(a.DB, studentLogin, "apc_database", "student"); err != nil {
		if err.Error() == "mongo: no documents in result" {
			ret := student.StudentPage{
				UserExist: false,
			}
			respondWithJSON(w, http.StatusOK, ret)
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if newsArray, err = news.GetNews(a.DB, "apc_database", "news"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ret := student.StudentPage{
		UserExist: true,
		User:      singleStudent,
		News:      newsArray,
	}

	respondWithJSON(w, http.StatusOK, ret)
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

	var students []student.Student

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

	var students []student.Student

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&students); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := student.DeleteStudents(a.DB, students, "apc_database", "student"); err != nil {
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

	if err := schoolClass.CreateClasses(a.DB, []schoolClass.SchoolClassCreate{}, "apc_database", "schoolClass"); err != nil {
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

///////////////////////////////////////////////////////////////////////////////////////////
// 								        NEWS		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createNews(w http.ResponseWriter, r *http.Request) {

	var newsArray []news.NewsCreate

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &newsArray); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := news.CreateNews(a.DB, newsArray, "apc_database", "news"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

func (a *App) getNews(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) updateNews(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}

func (a *App) deleteNews(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "function not implemented"})
}
