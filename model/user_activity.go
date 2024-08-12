package model

type UserActivity struct {
	UserActivityId  string  `json:"userActivityId" bson:"userActivityId"`
	UserId          string  `json:"userId" bson:"userId"`
	ActivityId      string  `json:"activityId" bson:"activityId"`
	CarbonFootprint float32 `json:"carbonFootprint" bson:"carbonFootprint"`
	Timestamp       string  `json:"timestamp" bson:"timestamp"`
}
