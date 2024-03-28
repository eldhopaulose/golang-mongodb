package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	// Connect to MongoDB

	err := godotenv.Load() // Update the file path to load the .env file
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dbLink := os.Getenv("MONG_URI")
	if dbLink == "" {
		log.Fatal("MONG_URI is not set")
	}

	clientOptions := options.Client().ApplyURI(dbLink)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Set the collection variable to a collection from the database
	collection = client.Database("Go").Collection("users")
	if collection == nil {
		log.Fatal("Failed to get collection reference")
	}

	fmt.Println("Collection instance created!")
}

// GetCollection returns the collection variable
func GetCollection() *mongo.Collection {
	return collection
}
