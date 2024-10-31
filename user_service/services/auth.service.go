package services

import (
	"context"
	"errors"
	"user_service/auth"
	database "user_service/database"
	pb "user_service/protos"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (service *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var userType = req.Type
	var token string
	var name string
	var err error

	if userType == "Admin" {
		token, name, err = loginAdmin(req.Username, req.Password)
	} else if userType == "Teacher" {
		token, name, err = loginTeacher(req.Username, req.Password)
	} else {
		return &pb.LoginResponse{Message: "Chưa nhập type"}, nil
	}

	if err != nil {
		return &pb.LoginResponse{Message: err.Error()}, nil
	}
	return &pb.LoginResponse{Message: "Hello " + name, Token: token}, nil
}

func loginAdmin(username, password string) (token string, name string, err error) {
	user, err := database.FindUserByUsername(username)
	if err != nil {
		return "", "", errors.New("Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("Invalid email or password")
	}

	claims := auth.BaseClaims{
		UserID:   user.ID,
		Username: user.Username,
		Name:     user.Name,
		Title:    "Admin",
	}

	token, err = auth.GenerateJWT(claims)
	if err != nil {
		return "", "", errors.New("Failed to generate token")
	}

	return token, user.Name, nil
}

func loginTeacher(username, password string) (token string, name string, err error) {
	user, err := database.FindTeacherByUsername(username)
	if err != nil {
		return "", "", errors.New("Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("Invalid email or password")
	}

	claims := auth.BaseClaims{
		UserID:   user.ID,
		Username: user.Username,
		Name:     user.TeacherName,
		Title:    "Teacher",
	}

	token, err = auth.GenerateJWT(claims)
	if err != nil {
		return "", "", errors.New("Failed to generate token")
	}

	return token, user.TeacherName, nil
}
