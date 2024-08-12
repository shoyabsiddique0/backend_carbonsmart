package usecase

import (
	"backend/mod/model"
	"backend/mod/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserActivityService struct {
	MongoCollectionAct  *mongo.Collection
	MongoCollectionUser *mongo.Collection
}

func (s *UserActivityService) AddUserActivityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("Hit /userActivity/create")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	var userAc = model.UserActivity{}
	err := json.NewDecoder(r.Body).Decode(&userAc)
	if err != nil {
		res.Error = "Error decoding JSON"
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userAc.UserActivityId = uuid.NewString()
	userAc.Timestamp = time.Now().UTC().String()
	userRepo := repository.UserRepo{MongoCollection: s.MongoCollectionUser}
	repo := repository.UserActivityRepo{MongoCollection: s.MongoCollectionAct}
	user, err := userRepo.FindUserByID(userAc.UserId)

	if err != nil {
		res.Error = "User not found"
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Printf(
		"User: %+v\nActivity: %+v\nCarbon Footprint: %v\nScore: %v\n",
		user, userAc, userAc.CarbonFootprint, user.Score,
	)
	if userAc.CarbonFootprint > 2 {
		user.Score = (float32(user.Score) - userAc.CarbonFootprint/2)
	} else {
		user.Score = (user.Score + userAc.CarbonFootprint*10)
	}
	updated, err := userRepo.UpdateUser(user.UserID, user)
	if err != nil {
		res.Error = "Error updating User: " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(updated)
		return
	}
	result, err := repo.InsertUserActivity(userAc)
	if err != nil {
		res.Error = "Error inserting User Activity: " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = result
}

func (s *UserActivityService) GetUserActivityByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("Hit /userActivity/byUserId")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	params := mux.Vars(r)
	userId := params["userId"]
	repo := repository.UserActivityRepo{MongoCollection: s.MongoCollectionAct}
	activities, err := repo.FindUserActivitiesByUserID(userId)
	if err != nil && err != mongo.ErrNoDocuments {
		res.Error = "Error fetching User Activities: " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if activities == nil {
		res.Data = []model.UserActivity{}
		return
	}
	res.Data = activities
}

type timsestampdata struct {
	Timestamp string `json:"timestamp" bson:"timestamp"`
}

func (s *UserActivityService) GetUserActivityByUserIdTimeStampHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("Hit /userActivity/byUserIdAndTimestamp")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	params := mux.Vars(r)
	userId := params["userId"]
	timestamp := timsestampdata{}
	json.NewDecoder(r.Body).Decode(&timestamp)
	fmt.Println(timestamp.Timestamp, "Timestamp.Timestamp")
	repo := repository.UserActivityRepo{MongoCollection: s.MongoCollectionAct}
	layout := "2006-01-02"
	datetime, err := time.Parse(layout, strings.Split(timestamp.Timestamp, " ")[0])
	if err != nil {
		res.Error = "Error parsing timestamp: " + err.Error()
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(datetime, "DateTime.Timestamp")
	activity, err := repo.FindUserActivitiesByUserIDByTimeStamp(userId, datetime)
	if err != nil {
		res.Error = "Error fetching User Activity: " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = activity
}

func (s *UserActivityService) GetLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("Hit /leaderboard")
	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)
	// userId := mux.Vars(r)["userId"]
	userRepo := repository.UserRepo{MongoCollection: s.MongoCollectionUser}
	users, err := userRepo.FindAllUsers()

	if err != nil {
		res.Error = "Error fetching users: " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Sort users by score in descending order
	sort.Slice(users, func(i, j int) bool {
		return users[i].Score > users[j].Score
	})

	w.WriteHeader(http.StatusOK)
	res.Data = users
}
