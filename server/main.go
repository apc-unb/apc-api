package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Client
}

func (a *App) Initialize(host, port string) {

	var err error

	a.DB, err = GetMongoDB(host, port)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/users", a.createUsers).Methods("POST")
	// a.Router.HandleFunc("/users", a.updateUsers).Methods("PUT")
	// a.Router.HandleFunc("/users", a.deleteUsers).Methods("DELETE")
	// a.Router.HandleFunc("/users/{id}", a.getUser).Methods("GET")
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
