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

	router.HandleFunc("/student/login", s.studentLogin).Methods("POST", "OPTIONS")
	router.HandleFunc("/admin/login", s.adminLogin).Methods("POST", "OPTIONS")
	router.HandleFunc("/data", s.insertData).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())


	////////////////////
	// SECURE ROUTERS //
	////////////////////

	secureRouter := router.NewRoute().Subrouter()
	secureRouter.Use(middleware.SetMiddlewareAuthentication(s.JwtSecret))
	secureRouter.Use(middleware.SetMiddlewareJSON())

	secureRouter.HandleFunc("/student", s.getStudents).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/student/{classid}", s.getStudentsClass).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/student", s.createStudents).Methods("POST", "OPTIONS")
	secureRouter.HandleFunc("/student", s.updateStudents).Methods("PUT", "OPTIONS")
	secureRouter.HandleFunc("/student/file", s.createStudentsFile).Methods("POST", "OPTIONS")
	secureRouter.HandleFunc("/student/contest/{studentid}", s.getStudentIndividualProgress).Methods("GET", "OPTIONS")

	secureRouter.HandleFunc("/admin", s.getAdmins).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/admin", s.updateAdmins).Methods("PUT", "OPTIONS")
	secureRouter.HandleFunc("/admin/file", s.createAdminsFile).Methods("POST", "OPTIONS")
	secureRouter.HandleFunc("/admin/student", s.updateAdminStudent).Methods("PUT", "OPTIONS")

	secureRouter.HandleFunc("/class", s.getClasses).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/class/{professorid}", s.getClassProfessor).Methods("GET", "OPTIONS")


	secureRouter.HandleFunc("/submission", s.getSubmissions).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/submission", s.createSubmissions).Methods("POST", "OPTIONS")
	secureRouter.HandleFunc("/submission", s.updateSubmissions).Methods("PUT", "OPTIONS")
	secureRouter.HandleFunc("/submission", s.deleteSubmissions).Methods("DELETE", "OPTIONS")

	secureRouter.HandleFunc("/task", s.getTasks).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/task/{examid}", s.getTasksExam).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/task", s.createTasks).Methods("POST", "OPTIONS")
	secureRouter.HandleFunc("/task", s.updateTasks).Methods("PUT", "OPTIONS")
	secureRouter.HandleFunc("/task", s.deleteTasks).Methods("DELETE", "OPTIONS")

	secureRouter.HandleFunc("/exam", s.getExams).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/exam/{classid}", s.getExamsClass).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/exam", s.createExams).Methods("POST", "OPTIONS")
	secureRouter.HandleFunc("/exam", s.updateExams).Methods("PUT", "OPTIONS")
	secureRouter.HandleFunc("/exam", s.deleteExams).Methods("DELETE", "OPTIONS")

	secureRouter.HandleFunc("/news", s.getNews).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/news/{classid}", s.getNewsClass).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/news", s.createNews).Methods("POST", "OPTIONS")
	secureRouter.HandleFunc("/news", s.updateNews).Methods("PUT", "OPTIONS")
	secureRouter.HandleFunc("/news", s.deleteNews).Methods("DELETE", "OPTIONS")

	secureRouter.HandleFunc("/project", s.createProject).Methods("POST", "OPTIONS")
	secureRouter.HandleFunc("/project", s.updateProject).Methods("PUT", "OPTIONS")
	secureRouter.HandleFunc("/project/type", s.getProjectType).Methods("GET", "OPTIONS")
	secureRouter.HandleFunc("/project/status", s.updateStatusProject).Methods("PUT", "OPTIONS")
	secureRouter.HandleFunc("/project/{studentid}", s.getProjectStudent).Methods("GET", "OPTIONS")

	////////////////////
	// SECURE ROUTERS //
	////////////////////

	professorRouter := router.NewRoute().Subrouter()
	professorRouter.Use(middleware.SetMiddlewareAuthenticationProfessor(s.JwtSecret))
	professorRouter.Use(middleware.SetMiddlewareJSON())

	professorRouter.HandleFunc("/admin", s.createAdmins).Methods("POST", "OPTIONS")
	professorRouter.HandleFunc("/admin", s.deleteAdmin).Methods("DELETE", "OPTIONS")

	professorRouter.HandleFunc("/student", s.deleteStudents).Methods("DELETE", "OPTIONS")

	professorRouter.HandleFunc("/class", s.createClasses).Methods("POST", "OPTIONS")
	professorRouter.HandleFunc("/class", s.updateClass).Methods("PUT", "OPTIONS")
	professorRouter.HandleFunc("/class", s.deleteClasses).Methods("DELETE", "OPTIONS")

	professorRouter.HandleFunc("/project/type", s.createProjectType).Methods("POST", "OPTIONS")
	professorRouter.HandleFunc("/project/type", s.updateProjectType).Methods("PUT", "OPTIONS")
	professorRouter.HandleFunc("/project/type", s.deleteProjectType).Methods("DELETE", "OPTIONS")

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