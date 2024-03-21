package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eldhopaulose/mongo-golang/router"
)

func main() {
	address := getAddress()

	fmt.Println("MongoDB Go API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(address, r))
}

func getAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}
	// Additional logic to normalize host and port if needed
	return ":" + port // Add colon before the port number
}
