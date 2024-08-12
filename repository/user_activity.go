package repository

import (
	"backend/mod/model"
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserActivityRepo struct {
	MongoCollection *mongo.Collection
}

func (r *UserActivityRepo) InsertUserActivity(userActivity model.UserActivity) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), userActivity)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (r *UserActivityRepo) FindUserActivitiesByUserID(userID string) ([]model.UserActivity, error) {
	var userActivities []model.UserActivity
	filter := bson.D{{Key: "userId", Value: userID}}

	cursor, err := r.MongoCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var userActivity model.UserActivity
		err := cursor.Decode(&userActivity)
		if err != nil {
			return nil, err
		}
		userActivities = append(userActivities, userActivity)
	}

	return userActivities, nil
}

func (r *UserActivityRepo) FindUserActivitiesByUserIDByTimeStamp(userID string, timeStamp time.Time) ([]model.UserActivity, error) {
	activities, err := r.FindUserActivitiesByUserID(userID)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}
	filteredActivities := []model.UserActivity{}
	fmt.Println(activities)
	layout := "2006-01-02"
	reqDate, err := time.Parse(layout, strings.Split(timeStamp.String(), " ")[0])

	if err != nil {
		fmt.Println("Error parsing time", err)
		return nil, err
	}
	fmt.Println(reqDate)
	fmt.Println(timeStamp, "timeStamp")
	for index, activity := range activities {
		datetime, err := time.Parse(layout, strings.Split(activity.Timestamp, " ")[0])
		if err != nil {
			fmt.Println("Error parsing time", err)
			continue
		}
		fmt.Println(datetime.Equal(reqDate), "condition")
		if datetime.Equal(reqDate) {
			filteredActivities = append(filteredActivities, activities[index])
		}
	}
	return filteredActivities, nil
}
