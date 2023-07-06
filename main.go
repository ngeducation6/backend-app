package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("testdb").Collection("data")

	// Define CORS options
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.ExposedHeaders([]string{}),
		handlers.AllowCredentials(),
		handlers.MaxAge(86400),
	)

	router := http.NewServeMux()
	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request for /data")

		var result []Data

		cur, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}

		defer cur.Close(context.Background())

		for cur.Next(context.Background()) {
			var data Data
			err := cur.Decode(&data)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, data)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		log.Printf("Retrieved %d documents\n", len(result))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Fatal(err)
		}
	})

	router.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	// Apply CORS middleware
	handler := cors(router)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("Server running on port %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
