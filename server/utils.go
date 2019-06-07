package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetMongoDB(host, port string) (*mongo.Client, error) {

	db, err := mongo.Connect(context.TODO(), "mongodb://"+host+":"+port)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return db, nil
}
