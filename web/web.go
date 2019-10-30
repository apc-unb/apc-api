package web

import (
	"context"
	"net/http"
	"time"

	"github.com/apc-unb/apc-api/web/components/admin"
	"github.com/apc-unb/apc-api/web/components/user"

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

	classDAO := schoolClass.SchoolClassCreate{
		ProfessorFirstName: "Carla",
		ProfessorLastName:  "Castanho",
		ClassName:          "H",
		Address:            "PJC 144",
		Year:               2019,
		Season:             2,
		ContestsIDs: []int{227662, 227824, 229195, 230372, 231393, 231394, 232351},
		GroupID: "qpBtprcUFF",
	}

	classID := s.insert("schoolClass", classDAO)

	classDAO2 := schoolClass.SchoolClassCreate{
		ProfessorFirstName: "Caetano",
		ProfessorLastName:  "Veloso",
		ClassName:          "A",
		Address:            "PJC 101",
		Year:               2019,
		Season:             2,
		ContestsIDs:[]int{227662, 227824, 229195, 230372, 231393, 231394, 232351},
		GroupID: "qpBtprcUFF",
	}

	classID2 := s.insert("schoolClass", classDAO2)

	classDAO3 := schoolClass.SchoolClassCreate{
		ProfessorFirstName: "Caetano",
		ProfessorLastName:  "Veloso",
		ClassName:          "A",
		Address:            "PJC 101",
		Year:               2020,
		Season:             1,
		ContestsIDs:[]int{227662, 227824, 229195, 230372, 231393, 231394, 232351},
		GroupID: "qpBtprcUFF",
	}

	s.insert("schoolClass", classDAO3)

	studentDAO := student.StudentCreate{
		ClassID:   classID,
		FirstName: "Student",
		LastName:  "De Apc",
		Matricula: "321",
		Handles: student.StudentHandles{
			Codeforces: "FeMaiaF",
		},
		Email: "aluno@unb.com.br",
	}

	studentID := s.insert("student", studentDAO)

	studentCredentialsDAO := user.UserCredentials{
		ID:        studentID,
		Matricula: studentDAO.Matricula,
	}

	if studentCredentialsDAO.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	s.insert("student_login", studentCredentialsDAO)

	studentDAO2 := student.StudentCreate{
		ClassID:   classID,
		FirstName: "Aluno",
		LastName:  "De Apc",
		Matricula: "123",
		Handles: student.StudentHandles{
			Codeforces: "bnlz",
		},
		Email: "aluno@unb.com.br",
	}

	studentID2 := s.insert("student", studentDAO2)

	studentCredentialsDAO2 := user.UserCredentials{
		ID:        studentID2,
		Matricula: studentDAO2.Matricula,
	}

	if studentCredentialsDAO2.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	s.insert("student_login", studentCredentialsDAO2)

	studentDAO3 := student.StudentCreate{
		ClassID:   classID,
		FirstName: "Aluno",
		LastName:  "Altamente preparado de Apc",
		Matricula: "1234",
		Handles: student.StudentHandles{
			Codeforces: "kyara",
		},
		Email: "aluno@unb.com.br",
	}

	studentID3 := s.insert("student", studentDAO3)

	studentCredentialsDAO3 := user.UserCredentials{
		ID:        studentID3,
		Matricula: studentDAO3.Matricula,
	}

	if studentCredentialsDAO3.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	s.insert("student_login", studentCredentialsDAO3)

	monitorDAO1 := admin.AdminCreate{
		ClassID:   classID,
		FirstName: "Jose",
		LastName:  "Leite",
		Matricula: "1612346666",
		Email:     "email.do.jose@gmail.com",
		Projects:  6,
		Teacher: true,
	}

	monitorID1 := s.insert("admin", monitorDAO1)

	adminCredentialsDAO1 := user.UserCredentials{
		ID:        monitorID1,
		Matricula: monitorDAO1.Matricula,
	}

	if adminCredentialsDAO1.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	s.insert("admin_login", adminCredentialsDAO1)

	monitorDAO2 := admin.AdminCreate{
		ClassID:   classID,
		FirstName: "Luis",
		LastName:  "Gebrim",
		Matricula: "160146666",
		Email:     "email.do.luis@gmail.com",
		Projects:  4,
		Teacher: false,
	}

	monitorID2 := s.insert("admin", monitorDAO2)

	adminCredentialsDAO2 := user.UserCredentials{
		ID:        monitorID2,
		Matricula: monitorDAO2.Matricula,
	}

	if adminCredentialsDAO2.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	s.insert("admin_login", adminCredentialsDAO2)

	monitorDAO3 := admin.AdminCreate{
		ClassID:   classID2,
		FirstName: "Vitor",
		LastName:  "Dullens",
		Matricula: "1612346666",
		Email:     "email.do.dullens@gmail.com",
		Projects:  2,
	}

	monitorID3 := s.insert("admin", monitorDAO3)

	adminCredentialsDAO3 := user.UserCredentials{
		ID:        monitorID3,
		Matricula: monitorDAO3.Matricula,
	}

	if adminCredentialsDAO3.Password, err = utils.HashAndSalt([]byte("123")); err != nil {
		panic(err)
	}

	s.insert("admin_login", adminCredentialsDAO3)

	projectTypeDAO1 := project.ProjectType{
		Name:     "Trabalho 1",
		Description:    "Codar em C, somar 2 numeros",
		ClassID: classID,
		Start: time.Now(),
		End: time.Now().Add(time.Minute * 30),
		Score:    10.0,
	}

	projectType1ID := s.insert("projectType", projectTypeDAO1)

	projectTypeDAO2 := project.ProjectType{
		Name:     "Trabalho 2",
		Description:    "Codar em C, somar 3 numeros",
		ClassID: classID,
		Start: time.Now().Add(time.Minute * 30),
		End: time.Now().Add(time.Minute * 90),
		Score:    10.0,
	}

	projectType2ID := s.insert("projectType", projectTypeDAO2)

	studentProject1 := project.Project{
		StudentID:     studentID,
		ProjectTypeID: projectType1ID,
		MonitorID:     monitorID1,
		ClassID:       classID,
		CreatedAT:     time.Now(),
		UpdatedAT:     time.Now(),
		FileName:      "Veras hehe",
		Status:        project.Created,
		Score:         0.0,
	}

	s.insert("projects", studentProject1)

	studentProject2 := project.Project{
		StudentID:     studentID,
		ProjectTypeID: projectType2ID,
		MonitorID:     monitorID1,
		ClassID:       classID,
		CreatedAT:     time.Now(),
		UpdatedAT:     time.Now(),
		FileName:      "Veras2 hehe",
		Status:        project.Created,
		Score:         0.0,
	}

	s.insert("projects", studentProject2)

	examDAO := exam.ExamCreate{
		ClassID: classID,
		Title:   "Prova 1 APC",
	}

	examID := s.insert("exam", examDAO)

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



	////////////////////
	// PUBLIC ROUTERS //
	////////////////////

	router.HandleFunc("/student/login", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/student/login", s.studentLogin).Methods("POST")
	router.HandleFunc("/admin/login", s.getOptions).Methods("OPTIONS")
	router.HandleFunc("/admin/login", s.adminLogin).Methods("POST")
	router.HandleFunc("/data", s.insertData).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())



	////////////////////
	// SECURE ROUTERS //
	////////////////////


	secureRouter := router.NewRoute().Subrouter()
	secureRouter.Use(middleware.SetMiddlewareAuthentication())
	secureRouter.Use(middleware.SetMiddlewareJSON())

	secureRouter.HandleFunc("/student", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/student", s.getStudents).Methods("GET")
	secureRouter.HandleFunc("/student/{classid}", s.getStudentsClass).Methods("GET")
	secureRouter.HandleFunc("/student", s.createStudents).Methods("POST")
	secureRouter.HandleFunc("/student", s.updateStudents).Methods("PUT")
	secureRouter.HandleFunc("/student", s.deleteStudents).Methods("DELETE")
	secureRouter.HandleFunc("/student/file", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/student/file", s.createStudentsFile).Methods("POST")
	secureRouter.HandleFunc("/student/contest/{studentid}", s.getStudentIndividualProgress).Methods("GET")

	secureRouter.HandleFunc("/admin", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/admin", s.getAdmins).Methods("GET")
	secureRouter.HandleFunc("/admin", s.createAdmins).Methods("POST")
	secureRouter.HandleFunc("/admin", s.updateAdmins).Methods("PUT")
	secureRouter.HandleFunc("/admin/file", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/admin/file", s.createAdminsFile).Methods("POST")
	secureRouter.HandleFunc("/admin/student", s.updateAdminStudent).Methods("PUT")

	secureRouter.HandleFunc("/class", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/class", s.getClasses).Methods("GET")
	secureRouter.HandleFunc("/class", s.createClasses).Methods("POST")
	secureRouter.HandleFunc("/class", s.updateClasses).Methods("PUT")
	secureRouter.HandleFunc("/class", s.deleteClasses).Methods("DELETE")

	secureRouter.HandleFunc("/submission", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/submission", s.getSubmissions).Methods("GET")
	secureRouter.HandleFunc("/submission", s.createSubmissions).Methods("POST")
	secureRouter.HandleFunc("/submission", s.updateSubmissions).Methods("PUT")
	secureRouter.HandleFunc("/submission", s.deleteSubmissions).Methods("DELETE")

	secureRouter.HandleFunc("/task", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/task", s.getTasks).Methods("GET")
	secureRouter.HandleFunc("/task/{examid}", s.getTasksExam).Methods("GET")
	secureRouter.HandleFunc("/task", s.createTasks).Methods("POST")
	secureRouter.HandleFunc("/task", s.updateTasks).Methods("PUT")
	secureRouter.HandleFunc("/task", s.deleteTasks).Methods("DELETE")

	secureRouter.HandleFunc("/exam", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/exam", s.getExams).Methods("GET")
	secureRouter.HandleFunc("/exam/{classid}", s.getExamsClass).Methods("GET")
	secureRouter.HandleFunc("/exam", s.createExams).Methods("POST")
	secureRouter.HandleFunc("/exam", s.updateExams).Methods("PUT")
	secureRouter.HandleFunc("/exam", s.deleteExams).Methods("DELETE")

	secureRouter.HandleFunc("/news", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/news", s.getNews).Methods("GET")
	secureRouter.HandleFunc("/news/{classid}", s.getNewsClass).Methods("GET")
	secureRouter.HandleFunc("/news", s.createNews).Methods("POST")
	secureRouter.HandleFunc("/news", s.updateNews).Methods("PUT")
	secureRouter.HandleFunc("/news", s.deleteNews).Methods("DELETE")

	secureRouter.HandleFunc("/project", s.createProject).Methods("POST")
	secureRouter.HandleFunc("/project", s.updateProject).Methods("PUT")
	secureRouter.HandleFunc("/project", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/project/type", s.getProjectType).Methods("GET")
	secureRouter.HandleFunc("/project/check", s.checkProject).Methods("POST")
	secureRouter.HandleFunc("/project/status", s.updateStatusProject).Methods("PUT")
	secureRouter.HandleFunc("/project/{studentid}", s.getProjectStudent).Methods("GET")


	professorRouter := secureRouter.NewRoute().Subrouter()
	professorRouter.Use()

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