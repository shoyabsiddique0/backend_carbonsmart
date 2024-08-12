package router

import (
	"backend/mod/usecase"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func CategoryRouter(mongoClient *mongo.Client, mainRouter *mux.Router) mux.Router {
	router := mainRouter.PathPrefix("/category").Subrouter()
	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection("category")
	categoryService := usecase.CategoryService{MongoCollection: coll}
	router.HandleFunc("/create", categoryService.AddCategoryHandler).Methods(http.MethodPost)
	router.HandleFunc("/", categoryService.GetCategories).Methods(http.MethodGet)
	router.HandleFunc("/{id}", categoryService.GetCategoryById).Methods(http.MethodGet)
	router.HandleFunc("/update/{id}", categoryService.UpdateCategory).Methods(http.MethodPut)
	router.HandleFunc("/delete/{id}", categoryService.DeleteCategory).Methods(http.MethodDelete)
	return *router
}
