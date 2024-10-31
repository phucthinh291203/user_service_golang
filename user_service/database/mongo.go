package database

import (
	"context"
	"log"
	"os"
	"user_service/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient   *mongo.Client
	allCollection models.AllData
)

func ConnectToDatabase() {
	godotenv.Load()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Fatal("Kết nối đến database thất bại:", err)
		return
	}

	// Kiểm tra kết nối
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Kết nối đến database thất bại:", err)
		return
	}

	mongoClient = client
	allCollection = models.AllData{
		TeacherCollection: mongoClient.Database(GetDBName()).Collection("teacherCollection"),
		AdminCollection:   mongoClient.Database(GetDBName()).Collection("adminCollection"),
	}

	if mongoClient == nil {
		log.Fatal("mongoClient chưa được khởi tạo")
	} else {
		log.Println("mongoClient đã khởi tạo thành công")
	}
	log.Print("Kết nối đến database thành công!")
}

func GetDBName() string {
	return os.Getenv("dbName")
}

func GetTeacherCollection() *mongo.Collection {
	return allCollection.TeacherCollection
}

func GetAdminCollection() *mongo.Collection {
	return allCollection.AdminCollection
}

func FindUserByUsername(username string) (models.Admin, error) {
	var user models.Admin
	err := GetAdminCollection().FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.Admin{}, err
	}
	log.Print(user)
	return user, nil
}

func FindTeacherByUsername(username string) (models.Teacher, error) {
	var user models.Teacher
	err := GetAdminCollection().FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.Teacher{}, err
	}
	log.Print(user)
	return user, nil
}

//$2a$10$Qv.OChOXn50KfQUxkjFKwuHnaP0ttiB7kp0WfmdZWSVh9eElCLVja
