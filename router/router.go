package router

import (
	"github.com/eldhopaulose/mongo-golang/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user", controllers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/api/user/all", controllers.GetUserHandler).Methods("GET")

	return router
}
