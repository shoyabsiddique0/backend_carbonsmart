package model

type Activity struct {
	ActivityID      string `json:"activity_id,omitempty" bson:"activity_id"`
	CategoryID      string `json:"category_id,omitempty" bson:"category_id"`
	Name            string `json:"name,omitempty" bson:"name"`
	CarbonFootprint int    `json:"carbon_footprint,omitempty" bson:"carbon_footprint"`
	Unit            string `json:"unit,omitempty" bson:"unit"`
}
