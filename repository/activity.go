package repository

import (
	"backend/mod/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ActivityRepo struct {
	MongoCollection *mongo.Collection
}

func (r *ActivityRepo) InsertActivity(activity *model.Activity) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), activity)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *ActivityRepo) FindActivityByID(activityId string) (*model.Activity, error) {
	var activity model.Activity
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "activity_id", Value: activityId}}).Decode(&activity)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *ActivityRepo) FindByEmailPass(email string) (*model.Activity, error) {
	var activity model.Activity
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&activity)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *ActivityRepo) FindAllActivitys() ([]*model.Activity, error) {
	var activitys []*model.Activity
	cursor, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	// defer cursor.Close(context.Background())

	// for cursor.Next(context.Background()) {
	// 	var activity model.Activity
	// 	err := cursor.Decode(&activity)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	activitys = append(activitys, &activity)
	// }
	err = cursor.All(context.Background(), &activitys)
	if err != nil {
		return nil, err
	}

	return activitys, nil
}

func (r *ActivityRepo) UpdateActivity(activityId string, activity *model.Activity) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(), bson.D{{Key: "activity_id", Value: activityId}}, bson.D{{Key: "$set", Value: activity}})
	if err != nil {
		return 0, err
	}
	if result.MatchedCount == 0 {
		return 0, fmt.Errorf("activity not found")
	}
	return result.MatchedCount, nil
}

func (r *ActivityRepo) DeleteActivity(activityId string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(), bson.D{{Key: "activity_id", Value: activityId}})
	if err != nil {
		return 0, err
	}
	if result.DeletedCount == 0 {
		return 0, fmt.Errorf("activity not found")
	}
	return result.DeletedCount, nil
}

func (r *ActivityRepo) DeleteAll() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (r *ActivityRepo) FindByCategoryID(categoryId string) ([]*model.Activity, error) {
	var activity []*model.Activity
	result, err := r.MongoCollection.Find(context.Background(), bson.D{{Key: "category_id", Value: categoryId}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = result.All(context.Background(), &activity)
	if err == mongo.ErrNoDocuments {
		log.Println(err)
		return nil, err
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return activity, nil
}
