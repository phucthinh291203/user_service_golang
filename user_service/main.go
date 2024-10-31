package main

import (
	"context"
	"log"
	"net"
	"strings"
	"user_service/database"
	"user_service/middleware"
	pb "user_service/protos"
	"user_service/services"

	"google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

func main() {

	database.ConnectToDatabase()

	// Khởi tạo các roles yêu cầu
	requiredRoles := []string{"Admin", "Teacher"}

	// Thiết lập Interceptor với danh sách các roles cần thiết
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(roleBasedInterceptor(requiredRoles)),
	}

	grpcServer := grpc.NewServer(opts...)

	authService := services.NewAuthService()
	teacherService := services.NewTeacherService()

	pb.RegisterAuthServiceServer(grpcServer, authService)
	pb.RegisterTeacherServiceServer(grpcServer, teacherService)
	reflection.Register(grpcServer)

	// Lắng nghe trên cổng 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer.Serve(lis)

}

// Interceptor kiểm tra role chỉ cho TeacherService
func roleBasedInterceptor(requiredRoles []string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Chỉ xác thực role nếu service là TeacherService
		if strings.Contains(info.FullMethod, "TeacherService") {
			return middleware.JWTAuthInterceptor(requiredRoles)(ctx, req, info, handler)
		}
		// Nếu không phải TeacherService, bỏ qua Interceptor này
		return handler(ctx, req)
	}
}
