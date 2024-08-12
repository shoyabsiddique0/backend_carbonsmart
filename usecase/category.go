package usecase

import (
	"backend/mod/model"
	"backend/mod/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryService struct {
	MongoCollection *mongo.Collection
}

func (s *CategoryService) AddCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("Hit /category/create")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	var category = model.Category{}
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		res.Error = "Error decoding JSON"
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	category.CategoryID = uuid.NewString()
	repo := repository.CategoryRepo{MongoCollection: s.MongoCollection}
	result, err := repo.InsertCategory(&category)
	if err != nil {
		res.Error = "Error inserting category"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = result
}

func (s *CategoryService) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := &Response{}

	defer json.NewEncoder(w).Encode(&res)
	params := mux.Vars(r)
	categoryID := params["id"]
	var updatedCategory = model.Category{}
	err := json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		res.Error = "Error decoding JSON"
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	repo := repository.CategoryRepo{MongoCollection: s.MongoCollection}
	result, err := repo.UpdateCategory(categoryID, &updatedCategory)
	if err != nil {
		res.Error = "Error updating category"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = result
}

func (s *CategoryService) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	params := mux.Vars(r)
	categoryID := params["id"]
	repo := repository.CategoryRepo{MongoCollection: s.MongoCollection}
	result, err := repo.DeleteCategory(categoryID)
	if err != nil {
		res.Error = "Error deleting category"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println(result)
	res.Data = "Category deleted successfully"
}

func (s *CategoryService) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := &Response{}
	log.Println("Hit /category/")
	defer json.NewEncoder(w).Encode(&res)
	repo := repository.CategoryRepo{MongoCollection: s.MongoCollection}
	categories, err := repo.FindAllCategorys()
	if err != nil {
		res.Error = "Error retrieving categories"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = categories
}

func (s *CategoryService) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Hit /category/{id}")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	params := mux.Vars(r)
	categoryID := params["id"]
	repo := repository.CategoryRepo{MongoCollection: s.MongoCollection}
	category, err := repo.FindCategoryByID(categoryID)
	if err != nil {
		res.Error = "Error retrieving category"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if category == nil {
		res.Error = "Category not found"
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = category
}
