package web

import (
	"context"
	"net/http"
	"time"

	"github.com/apc-unb/apc-api/web/components/admin"

	"github.com/apc-unb/apc-api/web/components/exam"
	"github.com/apc-unb/apc-api/web/components/news"
	"github.com/apc-unb/apc-api/web/components/project"
	"github.com/apc-unb/apc-api/web/components/student"
	"github.com/apc-unb/apc-api/web/components/task"
	"github.com/apc-unb/apc-api/web/utils"
	"github.com/sirupsen/logrus"

	"github.com/apc-unb/apc-api/web/components/schoolClass"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/apc-unb/apc-api/web/config"
	"github.com/apc-unb/apc-api/web/middleware"
	"github.com/apc-unb/apc-api/web/prometheus"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Server is application struct data
type Server struct {
	*config.WebBuilder
}

// InitFromWebBuilder builds a Server instance
func (s *Server) InitFromWebBuilder(webBuilder *config.WebBuilder) *Server {
	s.WebBuilder = webBuilder
	logLevel, err := logrus.ParseLevel(s.LogLevel)
	if err != nil {
		logrus.Errorf("Not able to parse log level string. Setting default level: info.")
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)

	return s
}

func (s *Server) insertData(w http.ResponseWriter, r *http.Request) {

	var err error
	var classID primitive.ObjectID
	var examID primitive.ObjectID
	var studentID primitive.ObjectID
	var monitorID1 primitive.ObjectID
	var projectType1ID primitive.ObjectID
	var projectType2ID primitive.ObjectID

	classDAO := schoolClass.SchoolClassCreate{
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		ClassName:          "2019",
		Address:            "PJC 144",
		Year:               2019,
		Season:             2,
	}

	classID = s.insert("schoolClass", classDAO)

	studentDAO := student.StudentCreate{
		ClassID:   classID,
		FirstName: "Aluno",
		LastName:  "De Apc",
		Matricula: "123",
		Handles: student.StudentHandles{
			Codeforces: "Veras",
		},
		Email: "aluno@unb.com.br",
	}

	if studentDAO.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	studentID = s.insert("student", studentDAO)

	monitorDAO1 := admin.AdminCreate{
		ClassID:   classID,
		FirstName: "Jose",
		LastName:  "Leite",
		Matricula: "1612346666",
		Email:     "email.do.jose@gmail.com",
		Projects:  2,
	}

	if monitorDAO1.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	monitorID1 = s.insert("admin", monitorDAO1)

	monitorDAO2 := admin.AdminCreate{
		ClassID:   classID,
		FirstName: "Luis",
		LastName:  "Gebrim",
		Matricula: "160146666",
		Email:     "email.do.luis@gmail.com",
		Projects:  0,
	}

	if monitorDAO2.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	s.insert("admin", monitorDAO2)

	projectTypeDAO1 := project.ProjectType{
		Name:     "Trabalho 1",
		Order:    1,
		DeadLine: time.Now().Add(time.Minute * 30),
		Score:    10.0,
	}

	projectType1ID = s.insert("projectType", projectTypeDAO1)

	projectTypeDAO2 := project.ProjectType{
		Name:     "Trabalho 2",
		Order:    2,
		DeadLine: time.Now().Add(time.Minute * 60),
		Score:    4.0,
	}

	projectType2ID = s.insert("projectType", projectTypeDAO2)

	studentProject1 := project.Project{
		StudentID:     studentID,
		ProjectTypeID: projectType1ID,
		MonitorID:     monitorID1,
		SendTime:      time.Now(),
		FileName:      "Veras hehe",
		Status:        "Pending",
		Score:         0.0,
	}

	s.insert("projects", studentProject1)

	studentProject2 := project.Project{
		StudentID:     studentID,
		ProjectTypeID: projectType2ID,
		MonitorID:     monitorID1,
		SendTime:      time.Now(),
		FileName:      "Veras2 hehe",
		Status:        "Pending",
		Score:         0.0,
	}

	s.insert("projects", studentProject2)

	examDAO := exam.ExamCreate{
		ClassID: classID,
		Title:   "Prova 1 APC",
	}

	examID = s.insert("exam", examDAO)

	newsDAO := news.NewsCreate{
		ClassID:     classID,
		Title:       "Aula cancelada",
		Description: "Devido ao alinhamento da lua, hoje nao teremos aula",
		Tags:        []string{"Horóscopo", "É verdade esse bilhete"},
	}

	s.insert("news", newsDAO)

	taskDAO := task.TaskCreate{
		ExamID:    examID,
		Statement: "Some duas letras",
		Score:     7.5,
		Tags:      []string{"FFT", "Dinamic Programming", "Bitmask"},
	}

	s.insert("task", taskDAO)

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"result": "Data created!"})

}

func (s *Server) insert(collectionName string, data interface{}) primitive.ObjectID {

	var err error
	var dataID primitive.ObjectID
	var mongoReturn *mongo.InsertOneResult

	collection := s.DataBase.Database("apc_database").Collection(collectionName)

	if mongoReturn, err = collection.InsertOne(context.TODO(), data); err != nil {
		panic(err)
	} else {
		dataID = mongoReturn.InsertedID.(primitive.ObjectID)
	}

	return dataID

}

// Creates and run the server
func (s *Server) Run() error {

	prometheus.RecordUpTime()

	router := mux.NewRouter()
	router.Use(middleware.GetPrometheusMiddleware())
	router.Use(middleware.GetCorsMiddleware())

	router.HandleFunc("/student", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/student", s.getStudentLogin).Methods("POST")

	router.HandleFunc("/students", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/students", s.getStudents).Methods("GET")
	router.HandleFunc("/students/{classid}", s.getStudentsClass).Methods("GET")
	router.HandleFunc("/students", s.createStudents).Methods("POST")
	router.HandleFunc("/students", s.updateStudents).Methods("PUT")
	router.HandleFunc("/students", s.deleteStudents).Methods("DELETE")
	router.HandleFunc("/studentsFile", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/studentsFile", s.createStudentsFile).Methods("POST")

	router.HandleFunc("/admins", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/admins", s.getAdmins).Methods("GET")
	router.HandleFunc("/admins", s.createAdmins).Methods("POST")
	router.HandleFunc("/admins", s.updateAdmins).Methods("PUT")
	router.HandleFunc("/adminsFile", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/adminsFile", s.createAdminsFile).Methods("POST")
	router.HandleFunc("/admin/student", s.updateAdminStudent).Methods("PUT")

	router.HandleFunc("/classes", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/classes", s.getClasses).Methods("GET")
	router.HandleFunc("/classes", s.createClasses).Methods("POST")
	router.HandleFunc("/classes", s.updateClasses).Methods("PUT")
	router.HandleFunc("/classes", s.deleteClasses).Methods("DELETE")

	router.HandleFunc("/submissions", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/submissions", s.getSubmissions).Methods("GET")
	router.HandleFunc("/submissions", s.createSubmissions).Methods("POST")
	router.HandleFunc("/submissions", s.updateSubmissions).Methods("PUT")
	router.HandleFunc("/submissions", s.deleteSubmissions).Methods("DELETE")

	router.HandleFunc("/tasks", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/tasks", s.getTasks).Methods("GET")
	router.HandleFunc("/tasks/{examid}", s.getTasksExam).Methods("GET")
	router.HandleFunc("/tasks", s.createTasks).Methods("POST")
	router.HandleFunc("/tasks", s.updateTasks).Methods("PUT")
	router.HandleFunc("/tasks", s.deleteTasks).Methods("DELETE")

	router.HandleFunc("/exams", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/exams", s.getExams).Methods("GET")
	router.HandleFunc("/exams/{classid}", s.getExamsClass).Methods("GET")
	router.HandleFunc("/exams", s.createExams).Methods("POST")
	router.HandleFunc("/exams", s.updateExams).Methods("PUT")
	router.HandleFunc("/exams", s.deleteExams).Methods("DELETE")

	router.HandleFunc("/news", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/news", s.getNews).Methods("GET")
	router.HandleFunc("/news/{classid}", s.getNewsClass).Methods("GET")
	router.HandleFunc("/news", s.createNews).Methods("POST")
	router.HandleFunc("/news", s.updateNews).Methods("PUT")
	router.HandleFunc("/news", s.deleteNews).Methods("DELETE")

	router.HandleFunc("/projects/send", s.createProject).Methods("POST")
	router.HandleFunc("/projects/{studentid}", s.getProjectStudent).Methods("GET")

	router.HandleFunc("/data", s.insertData).Methods("GET")

	router.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:" + s.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Info("Initialized")
	err := srv.ListenAndServe()
	if err != nil {
		logrus.Fatal("server initialization error", err)
		return err
	}
	return nil

}
