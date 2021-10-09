package helper_user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	//config := GetConfiguration()
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://vignesh:vignesh@cluster0.gq2m8.mongodb.net/insta_go?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("insta_go").Collection("users")

	return collection
}

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

type Configuration struct {
	Port             string
	ConnectionString string
}

