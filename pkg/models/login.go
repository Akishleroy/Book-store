package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type LoginInput struct {
	UserName string
	Password string
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

func (l *LoginInput) Login() (User, string, error) {
	var getUser User
	fmt.Println(l.UserName, l.Password)
	db.Where("user_name=? AND password=?", l.UserName, l.Password).Find(&getUser)
	//fmt.Println(getUser.UserName, getUser.Password)
	if getUser.UserName == l.UserName && l.Password == getUser.Password {
		token, err := GenerateToken(getUser)
		return getUser, token, err
	}
	return getUser, "error1", nil

}

func GenerateToken(user User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte("123gsdfhhj12367124jdsf"))
}
