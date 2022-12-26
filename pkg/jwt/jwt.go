package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Server struct {
	Tokens       map[int64]*Token
	AccessSecret string
	AccessTTL    int
}

type Token struct {
	AccessToken  string
	RefreshToken string
}

type ServerTokens struct {
	tokens       map[int64]*Token
	accessSecret string
	accessTTL    int
}

func (server *Server) CreateToken(userId int64, userRole int64) (*Token, error) {
	accessTokenExp := time.Now().Add(time.Minute * time.Duration(server.AccessTTL)).Unix()

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["id"] = userId
	accessTokenClaims["iat"] = time.Now().Unix()
	accessTokenClaims["exp"] = accessTokenExp
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	signedAccessToken, err := accessToken.SignedString([]byte(server.AccessSecret))
	if err != nil {
		return nil, err
	}

	res := &Token{
		AccessToken: signedAccessToken,
	}

	// Remember this token
	server.Tokens[userId] = res

	return res, nil
}

func (server *Server) extractToken(tokenStr string) (int64, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, fmt.Errorf("Failed to extract token metadata, unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(server.AccessSecret), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userId, ok := claims["id"].(float64)
		if !ok {
			return 0, fmt.Errorf("No such key 'id'")
		}

		return int64(userId), nil
	}

	return 0, fmt.Errorf("Invalid token. Map claims not found")
}
