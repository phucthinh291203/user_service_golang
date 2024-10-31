package models

import "go.mongodb.org/mongo-driver/mongo"


type AllData struct {
	TeacherCollection *mongo.Collection
	AdminCollection *mongo.Collection
}