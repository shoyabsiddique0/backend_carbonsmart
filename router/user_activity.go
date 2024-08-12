package router

import (
	"backend/mod/usecase"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserActivityRouter(mongoClient *mongo.Client, mainRouter *mux.Router) mux.Router {
	router := mainRouter.PathPrefix("/userActivity").Subrouter()
	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection("user_activity")
	coll1 := mongoClient.Database(os.Getenv("DB_NAME")).Collection("userAuth")
	userAcSer := usecase.UserActivityService{MongoCollectionAct: coll, MongoCollectionUser: coll1}
	router.HandleFunc("/{userId}", userAcSer.GetUserActivityByUserIdHandler).Methods(http.MethodGet)
	router.HandleFunc("/time/{userId}", userAcSer.GetUserActivityByUserIdTimeStampHandler).Methods(http.MethodPost)
	router.HandleFunc("/create", userAcSer.AddUserActivityHandler).Methods(http.MethodPost)
	router.HandleFunc("/leaderboard", userAcSer.GetLeaderboardHandler).Methods(http.MethodGet)
	return *router
}
