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
	var classID2 primitive.ObjectID
	var examID primitive.ObjectID
	var studentID primitive.ObjectID
	var monitorID1 primitive.ObjectID
	var projectType1ID primitive.ObjectID
	var projectType2ID primitive.ObjectID

	classDAO := schoolClass.SchoolClassCreate{
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		ClassName:          "H",
		Address:            "PJC 144",
		Year:               2019,
		Season:             2,
	}

	classID = s.insert("schoolClass", classDAO)

	classDAO2 := schoolClass.SchoolClassCreate{
		ProfessorFirstName: "Caetano",
		ProfessorLastName:  "Veloso",
		ClassName:          "A",
		Address:            "PJC 101",
		Year:               2019,
		Season:             2,
	}

	classID2 = s.insert("schoolClass", classDAO2)

	classDAO3 := schoolClass.SchoolClassCreate{
		ProfessorFirstName: "Caetano",
		ProfessorLastName:  "Veloso",
		ClassName:          "A",
		Address:            "PJC 101",
		Year:               2020,
		Season:             1,
	}

	s.insert("schoolClass", classDAO3)

	studentDAO := student.StudentCreate{
		ClassID:   classID,
		FirstName: "Student",
		LastName:  "De Apc",
		Matricula: "321",
		Handles: student.StudentHandles{
			Codeforces: "Veras",
		},
		Email: "aluno@unb.com.br",
	}

	if studentDAO.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	studentID = s.insert("student", studentDAO)

	studentDAO2 := student.StudentCreate{
		ClassID:   classID,
		FirstName: "Aluno",
		LastName:  "De Apc",
		Matricula: "123",
		Handles: student.StudentHandles{
			Codeforces: "Veras",
		},
		Email: "aluno@unb.com.br",
	}

	if studentDAO2.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	s.insert("student", studentDAO2)

	studentDAO3 := student.StudentCreate{
		ClassID:   classID,
		FirstName: "Aluno",
		LastName:  "Altamente preparado de Apc",
		Matricula: "1234",
		Handles: student.StudentHandles{
			Codeforces: "Veras",
		},
		Email: "aluno@unb.com.br",
	}

	if studentDAO3.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	s.insert("student", studentDAO3)

	monitorDAO1 := admin.AdminCreate{
		ClassID:   classID,
		FirstName: "Jose",
		LastName:  "Leite",
		Matricula: "1612346666",
		Email:     "email.do.jose@gmail.com",
		Projects:  6,
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
		Projects:  4,
	}

	if monitorDAO2.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	s.insert("admin", monitorDAO2)

	monitorDAO3 := admin.AdminCreate{
		ClassID:   classID2,
		FirstName: "Vitor",
		LastName:  "Dullens",
		Matricula: "1612346666",
		Email:     "email.do.dullens@gmail.com",
		Projects:  2,
	}

	if monitorDAO3.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	monitorID1 = s.insert("admin", monitorDAO3)

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
		ClassID:       classID,
		CreatedAT:     time.Now(),
		FileName:      "Veras hehe",
		Status:        "Pending",
		Score:         0.0,
	}

	s.insert("projects", studentProject1)

	studentProject2 := project.Project{
		StudentID:     studentID,
		ProjectTypeID: projectType2ID,
		MonitorID:     monitorID1,
		ClassID:       classID,
		CreatedAT:     time.Now(),
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
		CreatedAT:   time.Now(),
		UpdatedAT:   time.Now(),
	}

	s.insert("news", newsDAO)

	newsDAO2 := news.NewsCreate{
		ClassID:     classID,
		Title:       "Cancelamento do cancelamento da aula",
		Description: "A lua voltou ao seu local normal, teremos aula",
		Tags:        []string{"Horóscopo", "É verdade esse bilhete"},
		CreatedAT:   time.Now(),
		UpdatedAT:   time.Now().Add(10 * time.Minute),
	}

	s.insert("news", newsDAO2)

	newsDAO3 := news.NewsCreate{
		ClassID:     classID,
		Title:       "Prova 1",
		Description: "Se ligue na prova 1 galera",
		Tags:        []string{"Prova 1", "Rumo ao MM"},
		CreatedAT:   time.Now(),
		UpdatedAT:   time.Now().Add(5 * time.Minute),
	}

	s.insert("news", newsDAO3)

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

	router.HandleFunc("/student/login", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/student/login", s.getStudentLogin).Methods("POST")

	router.HandleFunc("/student", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/student", s.getStudents).Methods("GET")
	router.HandleFunc("/student/{classid}", s.getStudentsClass).Methods("GET")
	router.HandleFunc("/student", s.createStudents).Methods("POST")
	router.HandleFunc("/student", s.updateStudents).Methods("PUT")
	router.HandleFunc("/student", s.deleteStudents).Methods("DELETE")
	router.HandleFunc("/student/file", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/student/file", s.createStudentsFile).Methods("POST")

	router.HandleFunc("/admin/login", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/admin/login", s.getAdminLogin).Methods("POST")

	router.HandleFunc("/admin", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/admin", s.getAdmins).Methods("GET")
	router.HandleFunc("/admin", s.createAdmins).Methods("POST")
	router.HandleFunc("/admin", s.updateAdmins).Methods("PUT")
	router.HandleFunc("/admin/file", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/admin/file", s.createAdminsFile).Methods("POST")
	router.HandleFunc("/admin/student", s.updateAdminStudent).Methods("PUT")

	router.HandleFunc("/class", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/class", s.getClasses).Methods("GET")
	router.HandleFunc("/class", s.createClasses).Methods("POST")
	router.HandleFunc("/class", s.updateClasses).Methods("PUT")
	router.HandleFunc("/class", s.deleteClasses).Methods("DELETE")

	router.HandleFunc("/submission", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/submission", s.getSubmissions).Methods("GET")
	router.HandleFunc("/submission", s.createSubmissions).Methods("POST")
	router.HandleFunc("/submission", s.updateSubmissions).Methods("PUT")
	router.HandleFunc("/submission", s.deleteSubmissions).Methods("DELETE")

	router.HandleFunc("/task", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/task", s.getTasks).Methods("GET")
	router.HandleFunc("/task/{examid}", s.getTasksExam).Methods("GET")
	router.HandleFunc("/task", s.createTasks).Methods("POST")
	router.HandleFunc("/task", s.updateTasks).Methods("PUT")
	router.HandleFunc("/task", s.deleteTasks).Methods("DELETE")

	router.HandleFunc("/exam", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/exam", s.getExams).Methods("GET")
	router.HandleFunc("/exam/{classid}", s.getExamsClass).Methods("GET")
	router.HandleFunc("/exam", s.createExams).Methods("POST")
	router.HandleFunc("/exam", s.updateExams).Methods("PUT")
	router.HandleFunc("/exam", s.deleteExams).Methods("DELETE")

	router.HandleFunc("/news", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/news", s.getNews).Methods("GET")
	router.HandleFunc("/news/{classid}", s.getNewsClass).Methods("GET")
	router.HandleFunc("/news", s.createNews).Methods("POST")
	router.HandleFunc("/news", s.updateNews).Methods("PUT")
	router.HandleFunc("/news", s.deleteNews).Methods("DELETE")

	router.HandleFunc("/project", s.createProject).Methods("POST")
	router.HandleFunc("/project/status", s.updateStatusProject).Methods("PUT")
	router.HandleFunc("/project/{studentid}", s.getProjectStudent).Methods("GET")

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
