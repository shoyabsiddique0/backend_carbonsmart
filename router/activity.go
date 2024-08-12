package router

import (
	"backend/mod/usecase"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func ActivityRouter(mongoClient *mongo.Client, mainRouter *mux.Router) mux.Router {
	router := mainRouter.PathPrefix("/activity").Subrouter()
	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection("activity")
	categoryService := usecase.ActivityService{MongoCollection: coll}
	router.HandleFunc("/create", categoryService.AddActivityHandler).Methods(http.MethodPost)
	router.HandleFunc("/", categoryService.GetActivities).Methods(http.MethodGet)
	router.HandleFunc("/category/{id}", categoryService.GetActivityByCategory).Methods(http.MethodGet)
	router.HandleFunc("/{id}", categoryService.GetActivityById).Methods(http.MethodGet)
	router.HandleFunc("/update/{id}", categoryService.UpdateActivity).Methods(http.MethodPut)
	router.HandleFunc("/delete/{id}", categoryService.DeleteActivity).Methods(http.MethodDelete)
	return *router
}
