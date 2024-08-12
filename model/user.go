package model

type User struct {
	UserID   string  `json:"user_id,omitempty" bson:"user_id"`
	Name     string  `json:"name,omitempty" bson:"name"`
	Email    string  `json:"email,omitempty" bson:"email"`
	Password string  `json:"password,omitempty" bson:"password"`
	Age      uint    `json:"age,omitempty" bson:"age"`
	Gender   string  `json:"gender,omitempty" bson:"gender"`
	JWTToken string  `json:"jwt_token,omitempty" b:"jwt_token"`
	Score    float32 `json:"score,omitempty" bson:"score"`
}
