package main

import (
	"backend/mod/router"
	"context"
	"log"
	"net/http"
	"os"

	// "os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	// loading .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}
	log.Println("Env Loaded")
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Error connecting", err)
	}
	log.Println("Connected", mongoClient == nil)
}

// var uri string = "mongodb+srv://shoyabsiddique:Shoyab%40786@cluster0.t5bidza.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

func handleUserRequest() {
	mainRouter := mux.NewRouter()
	log.Println(os.Getenv("DB_NAME"))
	log.Println(os.Getenv("COLLECTION_NAME"))
	log.Println("number of databases", mongoClient)
	router.UserRouter(mongoClient, mainRouter)
	router.CategoryRouter(mongoClient, mainRouter)
	router.ActivityRouter(mongoClient, mainRouter)
	router.BarcodeRouter(mongoClient, mainRouter)
	router.UserActivityRouter(mongoClient, mainRouter)
	mainRouter.StrictSlash(true)
	log.Fatal(http.ListenAndServe(":3000", mainRouter))
}

func main() {
	// serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	defer mongoClient.Disconnect(context.TODO())
	handleUserRequest()
}
