package utils

import (
	"context"
	"fmt"
	"strings"

	jwtutil "github.com/kittipat1413/go-common/util/jwt"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
	UserID   string `json:"uid"`
	UserRole string `json:"user_role"`
}

type TokenResponse struct {
	Status string
	Id     string
	Role   string
}

// ProcessToken is used to check the token for validity and return status, user id and user role
func ProcessToken(token string) TokenResponse {

	var splits = strings.Fields(token)

	if len(splits) > 0 {
		ctx := context.Background()
		signingKey := []byte("7b8ebceb27141aacfbc79027")
		manager, err := jwtutil.NewJWTManager(jwtutil.HS256, signingKey)
		if err != nil {
			return TokenResponse{
				Status: "Invalid token",
				Id:     "",
				Role:   "",
			}
		}

		parsedClaims := &MyCustomClaims{}
		err = manager.ParseAndValidateToken(ctx, splits[1], parsedClaims)
		tokenStatus := "Valid token"
		if err != nil {
			//log.Fatalf("Failed to validate token: %v", err)
			tokenStatus = "Invalid token"
		}
		fmt.Printf("%s will expire\n", parsedClaims.ExpiresAt)
		return TokenResponse{
			Status: tokenStatus,
			Id:     parsedClaims.UserID,
			Role:   parsedClaims.UserRole,
		}
	} else {
		return TokenResponse{
			Status: "Invalid token",
			Id:     "",
			Role:   "Invalid",
		}
	}

}
