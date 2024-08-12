package model

type Category struct {
	CategoryID string `json:"category_id,omitempty" bson:"category_id"`
	Name       string `json:"name,omitempty" bson:"name"`
}
