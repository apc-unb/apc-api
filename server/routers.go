package main

import (
	"net/http"
)

func (a *App) createUsers(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}
