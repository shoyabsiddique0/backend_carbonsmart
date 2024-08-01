package repository

import (
	"backend/mod/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	MongoCollection *mongo.Collection
}

func (r *UserRepo) InsertUser(user *model.User) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *UserRepo) FindUserByID(userId string) (*model.User, error) {
	var user model.User
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "user_id", Value: userId}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) FindByEmailPass(email string) (*model.User, error) {
	var user model.User
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) FindAllUsers() ([]*model.User, error) {
	var users []*model.User
	cursor, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	// defer cursor.Close(context.Background())

	// for cursor.Next(context.Background()) {
	// 	var user model.User
	// 	err := cursor.Decode(&user)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	users = append(users, &user)
	// }
	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) UpdateUser(userId string, user *model.User) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(), bson.D{{Key: "user_id", Value: userId}}, bson.D{{Key: "$set", Value: user}})
	if err != nil {
		return 0, err
	}
	if result.MatchedCount == 0 {
		return 0, fmt.Errorf("user not found")
	}
	return result.MatchedCount, nil
}

func (r *UserRepo) DeleteUser(userId string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(), bson.D{{Key: "user_id", Value: userId}})
	if err != nil {
		return 0, err
	}
	if result.DeletedCount == 0 {
		return 0, fmt.Errorf("user not found")
	}
	return result.DeletedCount, nil
}

func (r *UserRepo) DeleteAll() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
