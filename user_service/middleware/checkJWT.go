package middleware

import (
	"context"
	"errors"
	"strings"

	auth "user_service/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func JWTAuthInterceptor(requiredTitles []string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Lấy metadata từ context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("missing metadata")
		}

		// Lấy token từ header Authorization
		token := md["authorization"]
		if len(token) == 0 {
			return nil, errors.New("authorization header missing")
		}

		// Xóa "Bearer " khỏi token
		tokenString := strings.TrimPrefix(token[0], "Bearer ")

		// Giải mã và kiểm tra token
		claims, err := auth.ParseJWT(tokenString)
		if err != nil {
			return nil, errors.New("invalid token")
		}

		// Kiểm tra vai trò của người dùng
		title := claims.Title
		roleAllowed := false
		for _, requiredTitle := range requiredTitles {
			if title == requiredTitle {
				roleAllowed = true
				break
			}
		}
		if !roleAllowed {
			return nil, errors.New("insufficient permissions")
		}

		// Thêm thông tin người dùng vào context
		newCtx := context.WithValue(ctx, "username", claims.Username)
		newCtx = context.WithValue(newCtx, "name", claims.Name)
		if claims.Title == "Teacher" {
			newCtx = context.WithValue(newCtx, "teacher_id", claims.UserID)
		}
		if claims.Title == "Admin" {
			newCtx = context.WithValue(newCtx, "admin_id", claims.UserID)
		}

		// Gọi handler với context mới
		return handler(newCtx, req)
	}
}
