package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/plataforma-apc/components/contest"
	"github.com/plataforma-apc/components/news"
	"github.com/plataforma-apc/components/schoolClass"
	"github.com/plataforma-apc/components/student"
	"github.com/plataforma-apc/components/submission"
	"github.com/plataforma-apc/components/task"
)

///////////////////////////////////////////////////////////////////////////////////////////
// 									 STUDENTS			 								 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) getStudentLogin(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

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

func (a *App) createStudents(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

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

func (a *App) createStudentsFile(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	body, _ := ioutil.ReadAll(r.Body)
	bodyString := string(body)
	fmt.Println("Deu = ", bodyString)

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getStudents(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	students, err := student.GetStudents(a.DB, "apc_database", "student")

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, students)
}

func (a *App) updateStudents(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var studentUpdate student.StudentUpdate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&studentUpdate); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := student.UpdateStudents(a.DB, studentUpdate, "apc_database", "student"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deleteStudents(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

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

	enableCORS(&w)

	var classes []schoolClass.SchoolClassCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&classes); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := schoolClass.CreateClasses(a.DB, classes, "apc_database", "schoolClass"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getClasses(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	classes, err := schoolClass.GetClasses(a.DB, "apc_database", "schoolClass")

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, classes)

}

func (a *App) updateClasses(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var classes []schoolClass.SchoolClass

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&classes); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := schoolClass.UpdateClasses(a.DB, classes, "apc_database", "schoolClass"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deleteClasses(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var classes []schoolClass.SchoolClass

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&classes); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := schoolClass.DeleteClasses(a.DB, classes, "apc_database", "schoolClass"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

///////////////////////////////////////////////////////////////////////////////////////////
// 								      SUBMISSION		 					     		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createSubmissions(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var submissions []submission.SubmissionCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&submissions); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := submission.CreateSubmissions(a.DB, submissions, "apc_database", "submission"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getSubmissions(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	submissions, err := submission.GetSubmissions(a.DB, "apc_database", "submission")

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, submissions)
}

func (a *App) updateSubmissions(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var submissions []submission.Submission

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&submissions); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := submission.UpdateSubmissions(a.DB, submissions, "apc_database", "submission"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deleteSubmissions(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var submissions []submission.Submission

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&submissions); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := submission.DeleteSubmissions(a.DB, submissions, "apc_database", "submission"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        TASK		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createTasks(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var tasks []task.TaskCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tasks); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := task.CreateTasks(a.DB, tasks, "apc_database", "task"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getTasks(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	tasks, err := task.GetTasks(a.DB, "apc_database", "task")

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, tasks)
}

func (a *App) updateTasks(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var tasks []task.Task

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tasks); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := task.UpdateTasks(a.DB, tasks, "apc_database", "task"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (a *App) deleteTasks(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var tasks []task.Task

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tasks); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := task.DeleteTasks(a.DB, tasks, "apc_database", "task"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        CONTEST		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createContests(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var contests []contest.ContestCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&contests); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := contest.CreateContests(a.DB, contests, "apc_database", "contest"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getContests(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	contests, err := contest.GetContests(a.DB, "apc_database", "contest")

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, contests)
}

func (a *App) updateContests(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var contests []contest.Contest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&contests); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := contest.UpdateContests(a.DB, contests, "apc_database", "contest"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deleteContests(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var contests []contest.Contest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&contests); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := contest.DeleteContests(a.DB, contests, "apc_database", "contest"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        NEWS		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createNews(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

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

	enableCORS(&w)

	newsArray, err := news.GetNews(a.DB, "apc_database", "news")

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, newsArray)
}

func (a *App) updateNews(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var newsArray []news.News

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &newsArray); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := news.UpdateNews(a.DB, newsArray, "apc_database", "news"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deleteNews(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)

	var newsArray []news.News

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &newsArray); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := news.DeleteNews(a.DB, newsArray, "apc_database", "news"); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}
