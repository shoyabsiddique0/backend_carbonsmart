package usecase

import (
	"backend/mod/model"
	"backend/mod/repository"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"crypto/sha1"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (svc *UserService) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	var user = model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Invalid Body", err)
		res.Error = err.Error()
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte("ifeellikethekingoftheworld"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Error creating token", err)
		res.Error = err.Error()
		return
	}
	user.JWTToken = tokenString
	user.UserID = uuid.NewString()
	hasher := sha1.New()
	sha := base64.URLEncoding.EncodeToString(hasher.Sum([]byte(user.Password)))
	user.Password = sha
	user.Score = 0
	log.Println("User Password", user.Password)
	repo := repository.UserRepo{MongoCollection: svc.MongoCollection}
	existingUser, err := repo.FindByEmailPass(user.Email)
	if existingUser != nil {
		w.WriteHeader(http.StatusConflict)
		log.Print("User already exists", err)
		res.Error = "User already exists"
		return
	}

	insertID, err := repo.InsertUser(&user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Error inserting user", err)
		res.Error = err.Error()
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Inserted user", insertID)
	res.Data = user

}

func (svc *UserService) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := &Response{}
	loginBody := &LoginBody{}
	defer json.NewEncoder(w).Encode(res)

	err := json.NewDecoder(r.Body).Decode(loginBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Invalid Body", err)
		res.Error = err.Error()
		return
	}
	hasher := sha1.New()
	sha := base64.URLEncoding.EncodeToString(hasher.Sum([]byte(loginBody.Password)))
	repo := repository.UserRepo{MongoCollection: svc.MongoCollection}
	existingUser, err := repo.FindByEmailPass(loginBody.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Something Went Wrong", err)
		res.Error = "Something Went Wrong"
		return
	}
	log.Println("This is existing user", existingUser)
	if existingUser == nil {
		w.WriteHeader(http.StatusOK)
		log.Print("User Not Found")
		res.Error = "User Not Found"
		return
	}
	if existingUser.Password != sha {
		w.WriteHeader(http.StatusUnauthorized)
		log.Print("Incorrect Password")
		res.Error = "Incorrect Password"
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("User Logged In", existingUser.UserID)
	res.Data = existingUser
}

func (svc *UserService) UpdateUserScoreHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	userId := mux.Vars(r)["id"]
	score := mux.Vars(r)["score"]
	log.Println("User ID", userId)

	// err := json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	log.Print("Invalid Body", err)
	// 	res.Error = err.Error()
	// 	return
	// }

	repo := repository.UserRepo{MongoCollection: svc.MongoCollection}
	user, err := repo.FindUserByID(userId)
	var userModel = *user
	data, err := strconv.ParseFloat(score, 32)
	userModel.Score = float32(data)
	result, err := repo.UpdateUser(userId, &userModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Error updating user", err)
		res.Error = err.Error()
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = result
}

func (svc *UserService) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	userId := mux.Vars(r)
	log.Println("User ID", userId)

	repo := repository.UserRepo{MongoCollection: svc.MongoCollection}
	user, err := repo.FindUserByID(userId["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Print("User not found", err)
		res.Error = err.Error()
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = user
	log.Println("User fetched", user)

}
func (svc *UserService) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.UserRepo{MongoCollection: svc.MongoCollection}
	user, err := repo.FindAllUsers()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error", err)
		res.Error = err.Error()
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Data = user
	log.Println("Users fetched", user)
}
func (svc *UserService) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	var user = model.User{}
	userId := mux.Vars(r)["id"]
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Invalid Body", err)
		res.Error = err.Error()
		return
	}
	// user.UserID = uuid.NewString()
	repo := repository.UserRepo{MongoCollection: svc.MongoCollection}
	insertID, err := repo.UpdateUser(userId, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Error updating user", err)
		res.Error = err.Error()
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("updated user", insertID)
	res.Data = user
}
func (svc *UserService) DeleteUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	// var user = model.User{}
	// err := json.NewDecoder(r.Body).Decode(&user)
	userId := mux.Vars(r)["id"]

	repo := repository.UserRepo{MongoCollection: svc.MongoCollection}
	insertID, err := repo.DeleteUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Error Deleting user", err)
		res.Error = err.Error()
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Deleted user", insertID)
	res.Data = insertID
}
func (svc *UserService) DeleteAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	// var user = model.User{}
	// err := json.NewDecoder(r.Body).Decode(&user)
	// userId := mux.Vars(r)["id"]

	repo := repository.UserRepo{MongoCollection: svc.MongoCollection}
	insertID, err := repo.DeleteAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Error Deleting user", err)
		res.Error = err.Error()
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Deleted user", insertID)
	res.Data = insertID
}
