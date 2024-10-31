package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Admin struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email       string             `json:"email" bson:"email"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password" bson:"password"`
	Name        string             `json:"name" bson:"name"`
	// Thêm các trường khác nếu cần
}