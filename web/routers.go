package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/apc-unb/apc-api/web/components/admin"
	"github.com/apc-unb/apc-api/web/components/exam"
	"github.com/apc-unb/apc-api/web/components/news"
	"github.com/apc-unb/apc-api/web/components/project"
	"github.com/apc-unb/apc-api/web/components/schoolClass"
	"github.com/apc-unb/apc-api/web/components/student"
	"github.com/apc-unb/apc-api/web/components/submission"
	"github.com/apc-unb/apc-api/web/components/task"
	"github.com/apc-unb/apc-api/web/components/user"
	"github.com/apc-unb/apc-api/web/utils"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

///////////////////////////////////////////////////////////////////////////////////////////
// 									 STUDENTS			 								 //
///////////////////////////////////////////////////////////////////////////////////////////

func (s *Server) studentLogin(w http.ResponseWriter, r *http.Request) {

	var UserCredentials user.UserCredentials
	var singleStudent student.StudentInfo
	var class schoolClass.SchoolClass
	var newsArray []news.News
	var err error

	decoder := json.NewDecoder(r.Body)

	if err = decoder.Decode(&UserCredentials); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if singleStudent, err = student.AuthStudent(s.DataBase, UserCredentials, "apc_database", "student"); err != nil {
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

	if class, err = schoolClass.GetClass(s.DataBase, singleStudent.ClassID, "apc_database", "schoolClass"); err != nil {
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

	if newsArray, err = news.GetNewsClass(s.DataBase, singleStudent.ClassID, "apc_database", "news"); err != nil {
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

func (s *Server) createStudents(w http.ResponseWriter, r *http.Request) {

	var students []student.StudentCreate
	var studentsList []user.UserCredentials
	var err error

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&students); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if studentsList, err = student.CreateStudents(s.DataBase, s.GoForces, students, "apc_database", "student"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonReturn := student.StudentCreatePage{
		Result:   "success",
		Students: studentsList,
	}

	utils.RespondWithJSON(w, http.StatusCreated, jsonReturn)
}

func (s *Server) createStudentsFile(w http.ResponseWriter, r *http.Request) {

	var studentsList []user.UserCredentials
	var err error

	request, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if studentsList, err = student.CreateStudentsFile(s.DataBase, string(request), "apc_database", "student"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonReturn := student.StudentCreatePage{
		Result:   "success",
		Students: studentsList,
	}

	utils.RespondWithJSON(w, http.StatusCreated, jsonReturn)

}

func (s *Server) getStudents(w http.ResponseWriter, r *http.Request) {

	students, err := student.GetStudents(s.DataBase, "apc_database", "student")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, students)
}

func (s *Server) getStudentsClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	classID, err := primitive.ObjectIDFromHex(vars["classid"])

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Class ID")
		return
	}

	students, err := student.GetStudentsClass(s.DataBase, classID, "apc_database", "student")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, students)

}

func (s *Server) updateStudents(w http.ResponseWriter, r *http.Request) {

	var studentUpdate student.StudentUpdate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&studentUpdate); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := student.UpdateStudents(s.DataBase, s.GoForces, studentUpdate, "apc_database", "student"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if studentUpdate.Email != "" {
		utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success", "email": studentUpdate.Email})
	} else {
		utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
	}
}

func (s *Server) deleteStudents(w http.ResponseWriter, r *http.Request) {

	var students []student.Student

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&students); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := student.DeleteStudents(s.DataBase, students, "apc_database", "student"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								   CLASS OF STUDENTS         							 //
///////////////////////////////////////////////////////////////////////////////////////////

func (s *Server) createClasses(w http.ResponseWriter, r *http.Request) {

	var classes []schoolClass.SchoolClassCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&classes); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := schoolClass.CreateClasses(s.DataBase, classes, "apc_database", "schoolClass"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (s *Server) getClasses(w http.ResponseWriter, r *http.Request) {

	classes, err := schoolClass.GetClasses(s.DataBase, "apc_database", "schoolClass")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, classes)

}

func (s *Server) updateClasses(w http.ResponseWriter, r *http.Request) {

	var classes []schoolClass.SchoolClass

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&classes); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := schoolClass.UpdateClasses(s.DataBase, classes, "apc_database", "schoolClass"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (s *Server) deleteClasses(w http.ResponseWriter, r *http.Request) {

	var classes []schoolClass.SchoolClass

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&classes); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := schoolClass.DeleteClasses(s.DataBase, classes, "apc_database", "schoolClass"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

///////////////////////////////////////////////////////////////////////////////////////////
// 								      SUBMISSION		 					     		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (s *Server) createSubmissions(w http.ResponseWriter, r *http.Request) {

	var submissions []submission.SubmissionCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&submissions); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := submission.CreateSubmissions(s.DataBase, submissions, "apc_database", "submission"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (s *Server) getSubmissions(w http.ResponseWriter, r *http.Request) {

	submissions, err := submission.GetSubmissions(s.DataBase, "apc_database", "submission")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, submissions)
}

func (s *Server) updateSubmissions(w http.ResponseWriter, r *http.Request) {

	var submissions []submission.Submission

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&submissions); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := submission.UpdateSubmissions(s.DataBase, submissions, "apc_database", "submission"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (s *Server) deleteSubmissions(w http.ResponseWriter, r *http.Request) {

	var submissions []submission.Submission

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&submissions); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := submission.DeleteSubmissions(s.DataBase, submissions, "apc_database", "submission"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        TASK		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (s *Server) createTasks(w http.ResponseWriter, r *http.Request) {

	var tasks []task.TaskCreate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tasks); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := task.CreateTasks(s.DataBase, tasks, "apc_database", "task"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (s *Server) getTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := task.GetTasks(s.DataBase, "apc_database", "task")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, tasks)
}

func (s *Server) getTasksExam(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	examID, err := primitive.ObjectIDFromHex(vars["examid"])

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Exam ID")
		return
	}

	tasks, err := task.GetTasksClass(s.DataBase, examID, "apc_database", "task")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tasks)

}

func (s *Server) updateTasks(w http.ResponseWriter, r *http.Request) {

	var tasks []task.Task

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tasks); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := task.UpdateTasks(s.DataBase, tasks, "apc_database", "task"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

func (s *Server) deleteTasks(w http.ResponseWriter, r *http.Request) {

	var tasks []task.Task

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tasks); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := task.DeleteTasks(s.DataBase, tasks, "apc_database", "task"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        EXAM		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (s *Server) createExams(w http.ResponseWriter, r *http.Request) {

	var exams []exam.ExamCreate
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&exams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := exam.CreateExams(s.DataBase, exams, "apc_database", "exam"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (s *Server) getExamsClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	classID, err := primitive.ObjectIDFromHex(vars["classid"])

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Class ID")
		return
	}

	exams, err := exam.GetExamsClass(s.DataBase, classID, "apc_database", "exam")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, exams)

}

func (s *Server) getExams(w http.ResponseWriter, r *http.Request) {

	exams, err := exam.GetExams(s.DataBase, "apc_database", "exam")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, exams)
}

func (s *Server) updateExams(w http.ResponseWriter, r *http.Request) {

	var exams []exam.Exam
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&exams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := exam.UpdateExams(s.DataBase, exams, "apc_database", "exam"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (s *Server) deleteExams(w http.ResponseWriter, r *http.Request) {

	var exams []exam.Exam
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&exams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := exam.DeleteExams(s.DataBase, exams, "apc_database", "exam"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        NEWS		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (s *Server) createNews(w http.ResponseWriter, r *http.Request) {

	var newsArray []news.NewsCreate

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &newsArray); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := news.CreateNews(s.DataBase, newsArray, "apc_database", "news"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

func (s *Server) getNews(w http.ResponseWriter, r *http.Request) {

	news, err := news.GetNews(s.DataBase, "apc_database", "news")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, news)
}

func (s *Server) getNewsClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	classID, err := primitive.ObjectIDFromHex(vars["classid"])

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Class ID")
		return
	}

	newsArray, err := news.GetNewsClass(s.DataBase, classID, "apc_database", "news")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, newsArray)

}

func (s *Server) updateNews(w http.ResponseWriter, r *http.Request) {

	var newsArray []news.News

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &newsArray); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := news.UpdateNews(s.DataBase, newsArray, "apc_database", "news"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (s *Server) deleteNews(w http.ResponseWriter, r *http.Request) {

	var newsArray []news.News

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &newsArray); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := news.DeleteNews(s.DataBase, newsArray, "apc_database", "news"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

///////////////////////////////////////////////////////////////////////////////////////////
// 									 ADMINS  			 								 //
///////////////////////////////////////////////////////////////////////////////////////////

func (s *Server) adminLogin(w http.ResponseWriter, r *http.Request) {

	var UserCredentials user.UserCredentials
	var singleAdmin admin.AdminInfo
	var class schoolClass.SchoolClass
	var newsArray []news.News
	var err error

	decoder := json.NewDecoder(r.Body)

	if err = decoder.Decode(&UserCredentials); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if singleAdmin, err = admin.AuthAdmin(s.DataBase, UserCredentials, "apc_database", "admin"); err != nil {
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

	if class, err = schoolClass.GetClass(s.DataBase, singleAdmin.ClassID, "apc_database", "schoolClass"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if newsArray, err = news.GetNewsClass(s.DataBase, singleAdmin.ClassID, "apc_database", "news"); err != nil {
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

func (s *Server) createAdmins(w http.ResponseWriter, r *http.Request) {

	var admins []admin.AdminCreate
	var adminsList []user.UserCredentials
	var err error

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&admins); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if adminsList, err = admin.CreateAdmin(s.DataBase, s.GoForces, admins, "apc_database", "admin"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonReturn := admin.AdminCreatePage{
		Result: "success",
		Admins: adminsList,
	}

	utils.RespondWithJSON(w, http.StatusCreated, jsonReturn)
}

func (s *Server) createAdminsFile(w http.ResponseWriter, r *http.Request) {

	var adminList []user.UserCredentials
	var err error

	request, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if adminList, err = admin.CreateAdminFile(s.DataBase, string(request), "apc_database", "student"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonReturn := student.StudentCreatePage{
		Result:   "success",
		Students: adminList,
	}

	utils.RespondWithJSON(w, http.StatusCreated, jsonReturn)

}

func (s *Server) getAdmins(w http.ResponseWriter, r *http.Request) {

	admins, err := admin.GetAdmins(s.DataBase, "apc_database", "admin")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, admins)
}

func (s *Server) updateAdmins(w http.ResponseWriter, r *http.Request) {

	var adminUpdate admin.AdminUpdate

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&adminUpdate); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := admin.UpdateAdmins(s.DataBase, s.GoForces, adminUpdate, "apc_database", "admin"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (s *Server) updateAdminStudent(w http.ResponseWriter, r *http.Request) {

	var adminUpdateStudent admin.AdminUpdateStudent

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&adminUpdateStudent); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := admin.UpdateAdminStudent(s.DataBase, s.GoForces, adminUpdateStudent, "apc_database", "student", "admin_login"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

func (s *Server) getOptions(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, nil)
}

///////////////////////////////////////////////////////////////////////////////////////////
// 								        PROJECTS		 				            		 //
///////////////////////////////////////////////////////////////////////////////////////////

func (s *Server) getProjectStudent(w http.ResponseWriter, r *http.Request) {

	var studentProjects []project.Project
	var err error
	vars := mux.Vars(r)
	studentID, err := primitive.ObjectIDFromHex(vars["studentid"])

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Student ID")
		return
	}

	if studentProjects, err = project.GetProjects(s.DataBase, studentID, "apc_database", "projects"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, studentProjects)

}

func (s *Server) createProject(w http.ResponseWriter, r *http.Request) {

	var projectInfo project.Project
	var monitorInfo admin.Admin
	var err error

	decoder := json.NewDecoder(r.Body)

	if err = decoder.Decode(&projectInfo); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if monitorInfo, err = project.CreateProject(s.DataBase, projectInfo, "apc_database"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	monitorReturn := map[string]interface{}{
		"status":       "success",
		"content" : map[string]interface{}{
			"monitorID":monitorInfo.ID,
			"monitorName":  monitorInfo.FirstName + " " + monitorInfo.LastName,
			"monitorEmail": monitorInfo.Email,
		},
	}


	utils.RespondWithJSON(w, http.StatusCreated, monitorReturn)

}

func (s *Server) updateStatusProject(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Recebi um request no update")

	var projectInfo project.Project

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&projectInfo); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := project.UpdateStatusProject(s.DataBase, projectInfo, "apc_database", "projects"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}


func (s *Server) getProjectType(w http.ResponseWriter, r *http.Request) {

	var types []project.ProjectType
	var err error

	if types, err = project.GetProjectsType(s.DataBase, "apc_database", "projectType"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types)

}