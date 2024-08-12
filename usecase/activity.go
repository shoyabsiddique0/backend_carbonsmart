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

type ActivityService struct {
	MongoCollection *mongo.Collection
}

func (s *ActivityService) AddActivityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	log.Println("Hit /activity/create")
	defer json.NewEncoder(w).Encode(&res)
	var activity = model.Activity{}
	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		res.Error = "Error decoding JSON"
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	activity.ActivityID = uuid.NewString()
	repo := repository.ActivityRepo{MongoCollection: s.MongoCollection}
	result, err := repo.InsertActivity(&activity)
	if err != nil {
		res.Error = "Error inserting activity"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = result
}

func (s *ActivityService) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	params := mux.Vars(r)
	activityID := params["id"]
	var updatedActivity = model.Activity{}
	err := json.NewDecoder(r.Body).Decode(&updatedActivity)
	if err != nil {
		res.Error = "Error decoding JSON"
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	repo := repository.ActivityRepo{MongoCollection: s.MongoCollection}
	result, err := repo.UpdateActivity(activityID, &updatedActivity)
	if err != nil {
		res.Error = "Error updating activity"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = result
}

func (s *ActivityService) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	params := mux.Vars(r)
	activityID := params["id"]
	repo := repository.ActivityRepo{MongoCollection: s.MongoCollection}
	result, err := repo.DeleteActivity(activityID)
	if err != nil {
		res.Error = "Error deleting activity"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println(result)
	res.Data = "Activity deleted successfully"
}

func (s *ActivityService) GetActivities(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit /activity")
	w.Header().Set("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	repo := repository.ActivityRepo{MongoCollection: s.MongoCollection}
	activities, err := repo.FindAllActivitys()
	if err != nil {
		res.Error = "Error retrieving activities"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = activities
}

func (s *ActivityService) GetActivityById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Hit /activity/{id}")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	params := mux.Vars(r)
	activityID := params["id"]
	repo := repository.ActivityRepo{MongoCollection: s.MongoCollection}
	activity, err := repo.FindActivityByID(activityID)
	if err != nil {
		res.Error = "Error retrieving activity"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if activity == nil {
		res.Error = "Activity not found"
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = activity
}

func (s *ActivityService) GetActivityByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Hit /activity/category")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	params := mux.Vars(r)
	category := params["id"]
	repo := repository.ActivityRepo{MongoCollection: s.MongoCollection}
	activities, err := repo.FindByCategoryID(category)
	if err != nil {
		res.Error = "Error retrieving activities by category"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if activities == nil {
		res.Error = "No activities found in this category"
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = activities
}
