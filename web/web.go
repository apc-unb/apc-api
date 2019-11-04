package web

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/apc-unb/apc-api/web/config"
	"github.com/apc-unb/apc-api/web/middleware"
	"github.com/apc-unb/apc-api/web/prometheus"

	"github.com/gorilla/mux"
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
	secureRouter.Use(middleware.SetMiddlewareAuthentication(s.JwtSecret))
	secureRouter.Use(middleware.SetMiddlewareJSON())

	secureRouter.HandleFunc("/student", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/student", s.getStudents).Methods("GET")
	secureRouter.HandleFunc("/student/{classid}", s.getStudentsClass).Methods("GET")
	secureRouter.HandleFunc("/student", s.createStudents).Methods("POST")
	secureRouter.HandleFunc("/student", s.updateStudents).Methods("PUT")
	secureRouter.HandleFunc("/student/file", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/student/file", s.createStudentsFile).Methods("POST")
	secureRouter.HandleFunc("/student/contest/{studentid}", s.getStudentIndividualProgress).Methods("GET")

	secureRouter.HandleFunc("/admin", s.getAdmins).Methods("GET")
	secureRouter.HandleFunc("/admin", s.updateAdmins).Methods("PUT")
	secureRouter.HandleFunc("/admin/file", s.getOptions).Methods("OPTIONS")
	secureRouter.HandleFunc("/admin/file", s.createAdminsFile).Methods("POST")
	secureRouter.HandleFunc("/admin/student", s.updateAdminStudent).Methods("PUT")

	secureRouter.HandleFunc("/class", s.getClasses).Methods("GET")
	secureRouter.HandleFunc("/class/{professorid}", s.getClassProfessor).Methods("GET")


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
	secureRouter.HandleFunc("/project/status", s.updateStatusProject).Methods("PUT")
	secureRouter.HandleFunc("/project/{studentid}", s.getProjectStudent).Methods("GET")

	////////////////////
	// SECURE ROUTERS //
	////////////////////

	professorRouter := router.NewRoute().Subrouter()
	professorRouter.Use(middleware.SetMiddlewareAuthenticationProfessor(s.JwtSecret))
	professorRouter.Use(middleware.SetMiddlewareJSON())

	professorRouter.HandleFunc("/admin", s.getOptions).Methods("OPTIONS")
	professorRouter.HandleFunc("/admin", s.createAdmins).Methods("POST")
	professorRouter.HandleFunc("/admin", s.deleteAdmin).Methods("DELETE")

	professorRouter.HandleFunc("/student", s.deleteStudents).Methods("DELETE")

	professorRouter.HandleFunc("/class", s.getOptions).Methods("OPTIONS")
	professorRouter.HandleFunc("/class", s.createClasses).Methods("POST")
	professorRouter.HandleFunc("/class", s.updateClass).Methods("PUT")
	professorRouter.HandleFunc("/class", s.deleteClasses).Methods("DELETE")

	professorRouter.HandleFunc("/project/type", s.getOptions).Methods("OPTIONS")
	professorRouter.HandleFunc("/project/type", s.createProjectType).Methods("POST")
	professorRouter.HandleFunc("/project/type", s.updateProjectType).Methods("PUT")
	professorRouter.HandleFunc("/project/type", s.deleteProjectType).Methods("DELETE")

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