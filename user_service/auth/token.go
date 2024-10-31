package auth

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var SecretKey = []byte(os.Getenv("secretKey"))

type BaseClaims struct {
	Username string             `json:"username"`
	Title    string             `json:"title"`
	Name     string             `json:"name"`
	UserID   primitive.ObjectID `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(claims BaseClaims) (string, error) {
	expirationTime := time.Now().Add(5 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix() // Thiết lập thời gian hết hạn

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)

}

func ParseJWT(tokenString string) (*BaseClaims, error) {
	claims := &BaseClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}


// TokenMiddleware là một interceptor để xác thực token

func TokenMiddleware(

    ctx context.Context,

    req interface{},

    info *grpc.UnaryServerInfo,

    handler grpc.UnaryHandler,

) (resp interface{}, err error) {

    // Lấy metadata từ context

    md, ok := metadata.FromIncomingContext(ctx)

    if !ok {

        return nil, fmt.Errorf("metadata is not provided")

    }


    // Lấy token từ metadata

    tokenStrs := md["authorization"]

    if len(tokenStrs) == 0 {

        return nil, fmt.Errorf("authorization token is not provided")

    }


    // Xác thực token

    tokenStr := tokenStrs[0]

    tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

        // Kiểm tra thuật toán

        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

        }

        return SecretKey, nil

    })


    if err != nil || !token.Valid {

        return nil, fmt.Errorf("invalid token: %v", err)

    }


    // Gọi handler tiếp theo

    return handler(ctx, req)

}