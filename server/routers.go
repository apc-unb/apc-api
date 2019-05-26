package main

import (
	"fmt"
	"net/http"
)

func (a *App) createUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HELLOOOOOOOOOOOOOOOO")

	// collection := a.DB.Database("apc_database").Collection("apc_collection")

	// aluno := Teste{"Mikael Viado", 10, "Babaca"}

	// insertResult, err := collection.InsertOne(context.TODO(), aluno)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// var players []player.PlayerCreate
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&players); err != nil {
	// 	respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	// 	return
	// }
	// defer r.Body.Close()

	// if err := player.CreatePlayers(a.DB, players); err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {

	// collection := a.DB.Database("apc_database").Collection("apc_collection")

	// aluno := Teste{"Mikael Viado", 10, "Babaca"}

	// insertResult, err := collection.InsertOne(context.TODO(), aluno)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// var players []player.PlayerCreate
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&players); err != nil {
	// 	respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	// 	return
	// }
	// defer r.Body.Close()

	// if err := player.CreatePlayers(a.DB, players); err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}
