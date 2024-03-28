package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eldhopaulose/mongo-golang/db"
	"github.com/eldhopaulose/mongo-golang/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func CreateUser(user models.User) error {

	collection := db.GetCollection() // Get the collection
	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("error inserting user: %v", err)
	}
	log.Printf("Inserted a single document with ID: %v\n", insertResult.InsertedID)
	return nil
}

// CreateUserHandler creates a new user

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	if err := CreateUser(user); err != nil {
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

// GetUserHandler gets all users

func GetUser() ([]primitive.M, error) {
	collection := db.GetCollection()

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error getting users: %v", err)
	}

	var users []primitive.M
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, fmt.Errorf("error iterating over users: %v", err)
	}

	return users, nil
}

// GetUserHandler gets all users
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := GetUser() // Assign the result of GetUser() to a variable
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting users: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert users to []models.User
	var convertedUsers []models.User
	for _, user := range users {
		convertedUser := models.User{
			ID:       user["_id"].(primitive.ObjectID),
			Username: user["username"].(string),
			Age:      int(user["age"].(int32)), // Convert int32 to int
		}
		convertedUsers = append(convertedUsers, convertedUser)
	}

	type Response struct {
		Data []models.User `json:"data"`
	}

	response := Response{
		Data: convertedUsers,
	}

	json.NewEncoder(w).Encode(response)
}

// GetUserByID gets a user by ID
func GetUserByID(id string) (models.User, error) {
	collection := db.GetCollection()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, fmt.Errorf("error converting id to objectID: %v", err)
	}

	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("error getting user: %v", err)
	}

	return user, nil
}

// GetUserByIDHandler gets a user by ID
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id := params["id"]

	if id == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	user, err := GetUserByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting user: %v", err), http.StatusInternalServerError)
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
