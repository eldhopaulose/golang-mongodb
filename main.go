package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eldhopaulose/mongo-golang/router"
)

func main() {
	fmt.Println("MongoDB Go API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
