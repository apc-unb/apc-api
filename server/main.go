package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/plataforma-apc/config"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Client
}

func (a *App) Initialize(host, port string) {

	var err error

	a.DB, err = config.GetMongoDB(host, port)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/student", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/student", a.getStudentLogin).Methods("POST")

	a.Router.HandleFunc("/students", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/students", a.getStudents).Methods("GET")
	a.Router.HandleFunc("/students", a.createStudents).Methods("POST")
	a.Router.HandleFunc("/students", a.updateStudents).Methods("PUT")
	a.Router.HandleFunc("/students", a.deleteStudents).Methods("DELETE")

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
	a.Router.HandleFunc("/tasks", a.createTasks).Methods("POST")
	a.Router.HandleFunc("/tasks", a.updateTasks).Methods("PUT")
	a.Router.HandleFunc("/tasks", a.deleteTasks).Methods("DELETE")

	a.Router.HandleFunc("/contests", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/contests", a.getContests).Methods("GET")
	a.Router.HandleFunc("/contests", a.createContests).Methods("POST")
	a.Router.HandleFunc("/contests", a.updateContests).Methods("PUT")
	a.Router.HandleFunc("/contests", a.deleteContests).Methods("DELETE")

	a.Router.HandleFunc("/news", a.getOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/news", a.getNews).Methods("GET")
	a.Router.HandleFunc("/news", a.createNews).Methods("POST")
	a.Router.HandleFunc("/news", a.updateNews).Methods("PUT")
	a.Router.HandleFunc("/news", a.deleteNews).Methods("DELETE")

}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	a := App{}

	a.Initialize("localhost", "27017")

	defer a.DB.Disconnect(nil)

	a.Run(":8080")

}
