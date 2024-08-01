package repository

import (
	"context"
	"log"
	"testing"

	"backend/mod/model"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error connecting to Mongo")
	}
	log.Println("Connected to Mongo")
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Error pinging Mongo")
	}
	log.Println("Pinged Mongo")
	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	user1 := uuid.New().String()
	// user2 := uuid.New().String()
	coll := mongoTestClient.Database("CarbonSmart").Collection("userAuth")
	userRepo := UserRepo{MongoCollection: coll}

	t.Run("Insert User 1 Data", func(t *testing.T) {
		user := model.User{
			Name:   "Shoyab Siddique",
			UserID: user1,
			Age:    22,
			Gender: "Male",
		}
		result, err := userRepo.InsertUser(&user)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Insert User 1 Data %v", result)
	})

	t.Run("Get Employee", func(t *testing.T) {
		result, err := userRepo.FindUserByID(user1)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Get Employee %v", result)
	})
}
