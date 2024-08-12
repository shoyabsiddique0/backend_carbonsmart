package router

import (
	"backend/mod/usecase"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRouter(mongoClient *mongo.Client, mainRouter *mux.Router) mux.Router {
	router := mainRouter.PathPrefix("/user").Subrouter()
	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	userService := usecase.UserService{MongoCollection: coll}
	router.HandleFunc("/", userService.GetUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/{id}", userService.GetUserByIDHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", userService.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/create", userService.AddUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/update/{id}", userService.UpdateUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/update/{id}/{score}", userService.UpdateUserScoreHandler).Methods(http.MethodPut)
	router.HandleFunc("/delete/{id}", userService.DeleteUserByIDHandler).Methods(http.MethodDelete)
	router.HandleFunc("/delete", userService.DeleteAllUsersHandler).Methods(http.MethodDelete)
	return *router
}
