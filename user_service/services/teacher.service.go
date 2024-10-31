package services

import (
	"context"
	database "user_service/database"
	"user_service/models"
	pb "user_service/protos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TeacherService struct {
	pb.UnimplementedTeacherServiceServer
}

func NewTeacherService() *TeacherService {
	return &TeacherService{}
}

func (service *TeacherService) CreateNewTeacher(ctx context.Context, req *pb.CreateTeacherRequest) (*pb.CreateTeacherResponse, error) {
	// Kiểm tra xem tất cả thông tin đã được cung cấp
	if req.Username == "" {
		return &pb.CreateTeacherResponse{Message: "Username is required"}, nil
	}

	if req.TeacherName == "" {
		return &pb.CreateTeacherResponse{Message: "Name is required"}, nil
	}
	if req.Email == "" {
		return &pb.CreateTeacherResponse{Message: "Email is required"}, nil
	}
	if req.Password == "" {
		return &pb.CreateTeacherResponse{Message: "Password is required"}, nil
	}

	existedTeacher, err := database.FindTeacherByUsername(req.Username)
	if err == nil && existedTeacher.ID != primitive.NilObjectID {
		return &pb.CreateTeacherResponse{Message: "Tai khoan da ton tai"}, nil

	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	teacher := models.Teacher{
		TeacherName: req.TeacherName,
		Username:    req.Username,
		Email:       req.Email,
		Password:    string(hashedPassword),
	}

	// Chèn học sinh vào collection
	_, err = database.GetTeacherCollection().InsertOne(context.TODO(), teacher)

	if err != nil {
		return &pb.CreateTeacherResponse{Message: "Failed to create student"}, err
	}
	return &pb.CreateTeacherResponse{Message: "Them student thanh cong"}, nil
}

func (service *TeacherService) GetAllTeacher(ctx context.Context, emptypb *emptypb.Empty) (*pb.GetAllTeacherResponse, error) {
	cursors, err := database.GetTeacherCollection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var teacherList []*pb.Teacher
	for cursors.Next(context.TODO()) {
		var teacher models.Teacher
		cursors.Decode(&teacher)

		// Chuyển đổi ClassID và SubjectID từ primitive.ObjectID sang chuỗi
		classIDs := []string{}
		for _, classID := range teacher.ClassID {
			classIDs = append(classIDs, classID.Hex())
		}

		subjectIDs := []string{}
		for _, subjectID := range teacher.SubjectID {
			subjectIDs = append(subjectIDs, subjectID.Hex())
		}

		teacherList = append(teacherList, &pb.Teacher{
			ID:          teacher.ID.Hex(),
			TeacherName: teacher.TeacherName,
			ClassID:     classIDs,
			SubjectID:   subjectIDs,
		})
	}
	if len(teacherList) == 0 {
		return &pb.GetAllTeacherResponse{Message: "Khong tim thay data nao", Teachers: nil}, nil
	}

	return &pb.GetAllTeacherResponse{Message: "Tim thay tat ca data", Teachers: teacherList}, nil
}
