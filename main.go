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
	host := "localhost" // Replace with your desired host
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}
	// Additional logic to normalize host and port if needed
	return fmt.Sprintf("%s:%s", host, port)
}
