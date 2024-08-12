package router

import (
	"backend/mod/usecase"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func BarcodeRouter(mongoClient *mongo.Client, mainRouter *mux.Router) mux.Router {
	router := mainRouter.PathPrefix("/barcode").Subrouter()
	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection("barcode_history")
	barcodeService := usecase.BarcodeService{MongoCollection: coll}
	router.HandleFunc("/search/{id}/{user_id}", barcodeService.GetDetailsHandler).Methods(http.MethodGet)
	router.HandleFunc("/history/{user_id}", barcodeService.GetProductHistoryHandler).Methods(http.MethodGet)
	return *router
}
