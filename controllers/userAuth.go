package controllers

//register accout using mux

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eldhopaulose/mongo-golang/db"
	"github.com/eldhopaulose/mongo-golang/models"
)

func CreateAutUser(user models.User) error {

	collection := db.GetCollection() // Get the collection
	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("error inserting user: %v", err)
	}
	log.Printf("Inserted a single document with ID: %v\n", insertResult.InsertedID)
	return nil
}

// CreateUserHandler creates a new user

func CreateAuthUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	if err := CreateAutUser(user); err != nil {
		http.Error(w, fmt.Sprintf("error creating user: %v", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		models.User
		ID interface{} `json:"id"`
	}{
		User: user,
		ID:   user.ID,
	}

	json.NewEncoder(w).Encode(response)
}
