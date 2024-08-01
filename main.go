package main

import (
	"backend/mod/usecase"
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
	router := mux.NewRouter()
	log.Println(os.Getenv("DB_NAME"))
	log.Println(os.Getenv("COLLECTION_NAME"))
	log.Println("number of databases", mongoClient)
	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	userService := usecase.UserService{MongoCollection: coll}
	router.HandleFunc("/user", userService.GetUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", userService.GetUserByIDHandler).Methods(http.MethodGet)
	router.HandleFunc("/user/login", userService.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/user/create", userService.AddUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/user/update/{id}", userService.UpdateUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/user/update/{id}/{score}", userService.UpdateUserScoreHandler).Methods(http.MethodPut)
	router.HandleFunc("/user/delete/{id}", userService.DeleteUserByIDHandler).Methods(http.MethodDelete)
	router.HandleFunc("/user/delete", userService.DeleteAllUsersHandler).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	// serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	defer mongoClient.Disconnect(context.TODO())
	handleUserRequest()
}
