package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/VerasThiago/apc-api/components/exam"
	"github.com/VerasThiago/apc-api/components/news"
	"github.com/VerasThiago/apc-api/components/student"
	"github.com/VerasThiago/apc-api/components/task"
	"github.com/VerasThiago/apc-api/utils"

	"github.com/VerasThiago/apc-api/components/schoolClass"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/togatoga/goforces"

	"github.com/VerasThiago/apc-api/config"
	"github.com/VerasThiago/apc-api/middleware"
	"github.com/VerasThiago/apc-api/prometheus"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// App is application struct data
type App struct {
	Router *mux.Router
	DB     *mongo.Client
	API    *goforces.Client
}

// Initialize is a function that initialize  all tools for application
// Connect to mongo DB
// Connect to Codeforces API
// Calls function that create all routes
func (a *App) Initialize(host, port, codeforcesKey, codeforcesSecret string) {

	var err error

	if a.DB, err = config.GetMongoDB(host, port); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}

	if a.API, err = goforces.NewClient(log.New(os.Stderr, "*** ", log.LstdFlags)); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to Codeforces API!")
	}

	a.API.SetAPIKey(codeforcesKey)
	a.API.SetAPISecret(codeforcesSecret)

	a.Router = mux.NewRouter()
	a.Router.Use(middleware.GetPrometheusMiddleware())
	a.Router.Use(middleware.GetCorsMiddleware())

	a.initializeRoutes()
	a.insertData()
}

func (a *App) insertData() {

	var err error
	var classID primitive.ObjectID
	var examID primitive.ObjectID
	var studentID primitive.ObjectID
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

	classID = a.insert("schoolClass", classDAO)

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

	studentID = a.insert("student", studentDAO)

	projectTypeDAO1 := student.ProjectType{
		Name:     "Trabalho 1",
		Order:    1,
		DeadLine: time.Now().Add(time.Minute * 30),
		Score:    10.0,
	}

	projectType1ID = a.insert("projectType", projectTypeDAO1)

	projectTypeDAO2 := student.ProjectType{
		Name:     "Trabalho 2",
		Order:    2,
		DeadLine: time.Now().Add(time.Minute * 60),
		Score:    4.0,
	}

	projectType2ID = a.insert("projectType", projectTypeDAO2)

	studentProject1 := student.StudentProject{
		StudentID:     studentID,
		ProjectTypeID: projectType1ID,
		SendTime:      time.Now(),
		FileName:      "Veras hehe",
		Status:        "Pending",
		Score:         0.0,
	}

	a.insert("projects", studentProject1)

	studentProject2 := student.StudentProject{
		StudentID:     studentID,
		ProjectTypeID: projectType2ID,
		SendTime:      time.Now(),
		FileName:      "Veras2 hehe",
		Status:        "Pending",
		Score:         0.0,
	}

	a.insert("projects", studentProject2)

	examDAO := exam.ExamCreate{
		ClassID: classID,
		Title:   "Prova 1 APC",
	}

	examID = a.insert("exam", examDAO)

	newsDAO := news.NewsCreate{
		ClassID:     classID,
		Title:       "Aula cancelada",
		Description: "Devido ao alinhamento da lua, hoje nao teremos aula",
		Tags:        []string{"Horóscopo", "É verdade esse bilhete"},
	}

	a.insert("news", newsDAO)

	taskDAO := task.TaskCreate{
		ExamID:    examID,
		Statement: "Some duas letras",
		Score:     7.5,
		Tags:      []string{"FFT", "Dinamic Programming", "Bitmask"},
	}

	a.insert("task", taskDAO)

}

func (a *App) insert(collectionName string, data interface{}) primitive.ObjectID {

	var err error
	var dataID primitive.ObjectID
	var mongoReturn *mongo.InsertOneResult

	collection := a.DB.Database("apc_database").Collection(collectionName)

	if mongoReturn, err = collection.InsertOne(context.TODO(), data); err != nil {
		panic(err)
	} else {
		dataID = mongoReturn.InsertedID.(primitive.ObjectID)
	}

	return dataID

}

// initializeRoutes create all server routes
func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/student", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/student", a.getStudentLogin).Methods("POST")

	a.Router.HandleFunc("/students", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/students", a.getStudents).Methods("GET")
	a.Router.HandleFunc("/students/{classid}", a.getStudentsClass).Methods("GET")
	a.Router.HandleFunc("/students", a.createStudents).Methods("POST")
	a.Router.HandleFunc("/students", a.updateStudents).Methods("PUT")
	a.Router.HandleFunc("/students", a.deleteStudents).Methods("DELETE")
	a.Router.HandleFunc("/studentsFile", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/studentsFile", a.createStudentsFile).Methods("POST")

	a.Router.HandleFunc("/admins", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/admins", a.getAdmins).Methods("GET")
	a.Router.HandleFunc("/admins", a.createAdmins).Methods("POST")
	a.Router.HandleFunc("/admins", a.updateAdmins).Methods("PUT")
	a.Router.HandleFunc("/adminsFile", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/adminsFile", a.createAdminsFile).Methods("POST")
	a.Router.HandleFunc("/admin/student", a.updateAdminStudent).Methods("PUT")

	a.Router.HandleFunc("/classes", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/classes", a.getClasses).Methods("GET")
	a.Router.HandleFunc("/classes", a.createClasses).Methods("POST")
	a.Router.HandleFunc("/classes", a.updateClasses).Methods("PUT")
	a.Router.HandleFunc("/classes", a.deleteClasses).Methods("DELETE")

	a.Router.HandleFunc("/submissions", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/submissions", a.getSubmissions).Methods("GET")
	a.Router.HandleFunc("/submissions", a.createSubmissions).Methods("POST")
	a.Router.HandleFunc("/submissions", a.updateSubmissions).Methods("PUT")
	a.Router.HandleFunc("/submissions", a.deleteSubmissions).Methods("DELETE")

	a.Router.HandleFunc("/tasks", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/tasks", a.getTasks).Methods("GET")
	a.Router.HandleFunc("/tasks/{examid}", a.getTasksExam).Methods("GET")
	a.Router.HandleFunc("/tasks", a.createTasks).Methods("POST")
	a.Router.HandleFunc("/tasks", a.updateTasks).Methods("PUT")
	a.Router.HandleFunc("/tasks", a.deleteTasks).Methods("DELETE")

	a.Router.HandleFunc("/exams", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/exams", a.getExams).Methods("GET")
	a.Router.HandleFunc("/exams/{classid}", a.getExamsClass).Methods("GET")
	a.Router.HandleFunc("/exams", a.createExams).Methods("POST")
	a.Router.HandleFunc("/exams", a.updateExams).Methods("PUT")
	a.Router.HandleFunc("/exams", a.deleteExams).Methods("DELETE")

	a.Router.HandleFunc("/news", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/news", a.getNews).Methods("GET")
	a.Router.HandleFunc("/news/{classid}", a.getNewsClass).Methods("GET")
	a.Router.HandleFunc("/news", a.createNews).Methods("POST")
	a.Router.HandleFunc("/news", a.updateNews).Methods("PUT")
	a.Router.HandleFunc("/news", a.deleteNews).Methods("DELETE")

	a.Router.HandleFunc("/projects/send", a.createProject).Methods("POST")
	a.Router.HandleFunc("/projects/{studentid}", a.getProjectStudent).Methods("GET")

	a.Router.Handle("/metrics", promhttp.Handler())

}

// Run server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func Start() {

	prometheus.RecordUpTime()

	a := App{}
	//mongoHost := os.Getenv("CONN")
	//mongoHost := "apc-mongo"
	mongoHost := "localhost"
	a.Initialize(mongoHost, "27017", "f3d968eea83ad8d5f21cad0365edcc200439c6f0", "b30c206b689d5ba004534c6780aa7be8e234a7f3")

	defer a.DB.Disconnect(nil)

	a.Run(":8080")
}
