package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/VerasThiago/api/components/admin"
	"github.com/VerasThiago/api/components/exam"
	"github.com/VerasThiago/api/components/news"
	"github.com/VerasThiago/api/components/schoolClass"
	"github.com/VerasThiago/api/components/student"
	"github.com/VerasThiago/api/components/submission"
	"github.com/VerasThiago/api/components/task"
	"github.com/VerasThiago/api/utils"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

///////////////////////////////////////////////////////////////////////////////////////////
// 									 STUDENTS			 								 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) getStudentLogin(w http.ResponseWriter, r *http.Request) {

	var studentLogin student.StudentLogin
	var singleStudent student.StudentInfo
	var class schoolClass.SchoolClass
	var newsArray []news.News
	var err error

	decoder := json.NewDecoder(r.Body)

	if err = decoder.Decode(&studentLogin); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if singleStudent, err = student.AuthStudent(a.DB, studentLogin, "apc_database", "student"); err != nil {
		if err.Error() == "mongo: no documents in result" {
			ret := schoolClass.StudentPage{
				UserExist: false,
				Result:    "Invalid Login or Password",
			}
			utils.RespondWithJSON(w, http.StatusOK, ret)
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if class, err = schoolClass.GetClass(a.DB, singleStudent.ClassID, "apc_database", "schoolClass"); err != nil {
		if err.Error() == "mongo: no documents in result" {
			ret := schoolClass.StudentPage{
				UserExist: false,
				Result:    "Invalid student class",
			}
			utils.RespondWithJSON(w, http.StatusOK, ret)
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if newsArray, err = news.GetNewsClass(a.DB, singleStudent.ClassID, "apc_database", "news"); err != nil {
		if err.Error() == "mongo: no documents in result" {
			ret := schoolClass.StudentPage{
				UserExist: false,
				Result:    "Invalid student news",
			}
			utils.RespondWithJSON(w, http.StatusOK, ret)
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	ret := schoolClass.StudentPage{
		UserExist: true,
		Result:    "success",
		Student:   singleStudent,
		Class:     class,
		News:      newsArray,
	}

	utils.RespondWithJSON(w, http.StatusOK, ret)
}

func (a *App) createStudents(w http.ResponseWriter, r *http.Request) {

	var students []student.StudentCreate
	var studentsLits []student.StudentLogin
	var err error

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&students); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if studentsLits, err = student.CreateStudents(a.DB, a.API, students, "apc_database", "student"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonReturn := student.StudentCreatePage{
		Result:   "success",
		Students: studentsLits,
	}

	utils.RespondWithJSON(w, http.StatusCreated, jsonReturn)
}

func (a *App) createStudentsFile(w http.ResponseWriter, r *http.Request) {

	var studentsLits []student.StudentLogin
	var err error

	request, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if studentsLits, err = student.CreateStudentsFile(a.DB, string(request), "apc_database", "student"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonReturn := student.StudentCreatePage{
		Result:   "success",
		Students: studentsLits,
	}

	utils.RespondWithJSON(w, http.StatusCreated, jsonReturn)

}

func (a *App) getStudents(w http.ResponseWriter, r *http.Request) {

	students, err := student.GetStudents(a.DB, "apc_database", "student")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, students)
}

func (a *App) getStudentsClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	classID, err := primitive.ObjectIDFromHex(vars["classid"])

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Class ID")
		return
	}

	students, err := student.GetStudentsClass(a.DB, classID, "apc_database", "student")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, students)

}

func (a *App) updateStudents(w http.ResponseWriter, r *http.Request) {

	var studentUpdate student.StudentUpdate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&studentUpdate); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := student.UpdateStudents(a.DB, a.API, studentUpdate, "apc_database", "student"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if studentUpdate.Email != "" {
		utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success", "email": studentUpdate.Email})
	} else {
		utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
	}
}

func (a *App) deleteStudents(w http.ResponseWriter, r *http.Request) {

	var students []student.Student

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&students); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := student.DeleteStudents(a.DB, students, "apc_database", "student"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								   CLASS OF STUDENTS         							 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createClasses(w http.ResponseWriter, r *http.Request) {

	var classes []schoolClass.SchoolClassCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&classes); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := schoolClass.CreateClasses(a.DB, classes, "apc_database", "schoolClass"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getClasses(w http.ResponseWriter, r *http.Request) {

	classes, err := schoolClass.GetClasses(a.DB, "apc_database", "schoolClass")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, classes)

}

func (a *App) updateClasses(w http.ResponseWriter, r *http.Request) {

	var classes []schoolClass.SchoolClass

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&classes); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := schoolClass.UpdateClasses(a.DB, classes, "apc_database", "schoolClass"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deleteClasses(w http.ResponseWriter, r *http.Request) {

	var classes []schoolClass.SchoolClass

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&classes); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := schoolClass.DeleteClasses(a.DB, classes, "apc_database", "schoolClass"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

///////////////////////////////////////////////////////////////////////////////////////////
// 								      SUBMISSION		 					     		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createSubmissions(w http.ResponseWriter, r *http.Request) {

	var submissions []submission.SubmissionCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&submissions); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := submission.CreateSubmissions(a.DB, submissions, "apc_database", "submission"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getSubmissions(w http.ResponseWriter, r *http.Request) {

	submissions, err := submission.GetSubmissions(a.DB, "apc_database", "submission")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, submissions)
}

func (a *App) updateSubmissions(w http.ResponseWriter, r *http.Request) {

	var submissions []submission.Submission

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&submissions); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := submission.UpdateSubmissions(a.DB, submissions, "apc_database", "submission"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deleteSubmissions(w http.ResponseWriter, r *http.Request) {

	var submissions []submission.Submission

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&submissions); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := submission.DeleteSubmissions(a.DB, submissions, "apc_database", "submission"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        TASK		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createTasks(w http.ResponseWriter, r *http.Request) {

	var tasks []task.TaskCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tasks); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := task.CreateTasks(a.DB, tasks, "apc_database", "task"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := task.GetTasks(a.DB, "apc_database", "task")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, tasks)
}

func (a *App) getTasksExam(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	examID, err := primitive.ObjectIDFromHex(vars["examid"])

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Exam ID")
		return
	}

	tasks, err := task.GetTasksClass(a.DB, examID, "apc_database", "task")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tasks)

}

func (a *App) updateTasks(w http.ResponseWriter, r *http.Request) {

	var tasks []task.Task

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tasks); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := task.UpdateTasks(a.DB, tasks, "apc_database", "task"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

func (a *App) deleteTasks(w http.ResponseWriter, r *http.Request) {

	var tasks []task.Task

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tasks); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := task.DeleteTasks(a.DB, tasks, "apc_database", "task"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        EXAM		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createExams(w http.ResponseWriter, r *http.Request) {

	var exams []exam.ExamCreate
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&exams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := exam.CreateExams(a.DB, exams, "apc_database", "exam"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getExamsClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	classID, err := primitive.ObjectIDFromHex(vars["classid"])

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Class ID")
		return
	}

	exams, err := exam.GetExamsClass(a.DB, classID, "apc_database", "exam")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, exams)

}

func (a *App) getExams(w http.ResponseWriter, r *http.Request) {

	exams, err := exam.GetExams(a.DB, "apc_database", "exam")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, exams)
}

func (a *App) updateExams(w http.ResponseWriter, r *http.Request) {

	var exams []exam.Exam
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&exams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := exam.UpdateExams(a.DB, exams, "apc_database", "exam"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deleteExams(w http.ResponseWriter, r *http.Request) {

	var exams []exam.Exam
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&exams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := exam.DeleteExams(a.DB, exams, "apc_database", "exam"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        NEWS		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) createNews(w http.ResponseWriter, r *http.Request) {

	var newsArray []news.NewsCreate

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &newsArray); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := news.CreateNews(a.DB, newsArray, "apc_database", "news"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

func (a *App) getNews(w http.ResponseWriter, r *http.Request) {

	news, err := news.GetNews(a.DB, "apc_database", "news")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, news)
}

func (a *App) getNewsClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	classID, err := primitive.ObjectIDFromHex(vars["classid"])

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Class ID")
		return
	}

	newsArray, err := news.GetNewsClass(a.DB, classID, "apc_database", "news")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, newsArray)

}

func (a *App) updateNews(w http.ResponseWriter, r *http.Request) {

	var newsArray []news.News

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &newsArray); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := news.UpdateNews(a.DB, newsArray, "apc_database", "news"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deleteNews(w http.ResponseWriter, r *http.Request) {

	var newsArray []news.News

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &newsArray); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := news.DeleteNews(a.DB, newsArray, "apc_database", "news"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 									 ADMINS  			 								 //
///////////////////////////////////////////////////////////////////////////////////////////

func (a *App) getAdminLogin(w http.ResponseWriter, r *http.Request) {

	var adminLogin admin.AdminLogin
	var singleAdmin admin.AdminInfo
	var class schoolClass.SchoolClass
	var newsArray []news.News
	var err error

	decoder := json.NewDecoder(r.Body)

	if err = decoder.Decode(&adminLogin); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if singleAdmin, err = admin.AuthAdmin(a.DB, adminLogin, "apc_database", "admin"); err != nil {
		if err.Error() == "mongo: no documents in result" {
			ret := schoolClass.AdminPage{
				UserExist: false,
			}
			utils.RespondWithJSON(w, http.StatusOK, ret)
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if class, err = schoolClass.GetClass(a.DB, singleAdmin.ClassID, "apc_database", "schoolClass"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if newsArray, err = news.GetNewsClass(a.DB, singleAdmin.ClassID, "apc_database", "news"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ret := schoolClass.AdminPage{
		UserExist: true,
		Admin:     singleAdmin,
		Class:     class,
		News:      newsArray,
	}

	utils.RespondWithJSON(w, http.StatusOK, ret)
}

func (a *App) createAdmins(w http.ResponseWriter, r *http.Request) {

	var admins []admin.AdminCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&admins); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := admin.CreateAdmin(a.DB, a.API, admins, "apc_database", "admin"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) createAdminsFile(w http.ResponseWriter, r *http.Request) {

	request, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err := admin.CreateAdminFile(a.DB, string(request), "apc_database", "admin"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

func (a *App) getAdmins(w http.ResponseWriter, r *http.Request) {

	admins, err := admin.GetAdmins(a.DB, "apc_database", "admin")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, admins)
}

func (a *App) updateAdmins(w http.ResponseWriter, r *http.Request) {

	var adminUpdate admin.AdminUpdate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&adminUpdate); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := admin.UpdateAdmins(a.DB, a.API, adminUpdate, "apc_database", "admin"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) updateAdminStudent(w http.ResponseWriter, r *http.Request) {

	var adminUpdateStudent admin.AdminUpdateStudent

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&adminUpdateStudent); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := admin.UpdateAdminStudent(a.DB, a.API, adminUpdateStudent, "apc_database", "student"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

func (a *App) getOptions(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, nil)
}
