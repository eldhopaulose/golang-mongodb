package router

import (
	"github.com/eldhopaulose/mongo-golang/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	s := router.PathPrefix("/api").Subrouter() //Base Path

	s.HandleFunc("/user", controllers.CreateUserHandler).Methods("POST")
	s.HandleFunc("/user/all", controllers.GetUserHandler).Methods("GET")
	s.HandleFunc("/user/single/{id}", controllers.GetUserByIDHandler).Methods("GET")

	return router
}
