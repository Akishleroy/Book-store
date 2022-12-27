package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

type ServerTokens struct {
	tokens       map[int64]*Token
	accessSecret string
	accessTTL    int
}

func CreateToken(userId uint, userRole int64) (string, error) {
	accessTokenExp := time.Now().Add(time.Minute * time.Duration(86400)).Unix()

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["id"] = userId
	accessTokenClaims["iat"] = time.Now().Unix()
	accessTokenClaims["userType"] = userRole
	accessTokenClaims["exp"] = accessTokenExp
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	signedAccessToken, _ := accessToken.SignedString([]byte("secretKey123"))

	return signedAccessToken, nil
}

func ExtractToken(tokenStr string) (uint, int64, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, fmt.Errorf("Failed to extract token metadata, unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secretKey123"), nil
	})

	if err != nil {
		return 0, 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userId, ok := claims["id"].(float64)
		userType, ok := claims["userType"].(float64)
		if !ok {
			return 0, 0, fmt.Errorf("No such key 'id'")
		}

		return uint(userId), int64(userType), nil
	}

	return 0, 0, fmt.Errorf("Invalid token. Map claims not found")
}
