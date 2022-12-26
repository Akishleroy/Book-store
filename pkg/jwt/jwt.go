package jwt

import (
	"time"

	_ "github.com/Akishleroy/go-bookstore/config"
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

func (server *Server) CreateToken(userId int64, userRole int64) (*Token, error) {
	accessTokenExp := time.Now().Add(time.Minute * time.Duration(server.accessTTL)).Unix()

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["id"] = userId
	accessTokenClaims["role"] = userRole
	accessTokenClaims["iat"] = time.Now().Unix()
	accessTokenClaims["exp"] = accessTokenExp
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	signedAccessToken, err := accessToken.SignedString([]byte(server.accessSecret))
	if err != nil {
		return nil, err
	}

	res := &Token{
		AccessToken: signedAccessToken,
	}

	// Remember this token
	server.tokens[userId] = res

	return res, nil
}
