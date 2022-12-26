package models

import (
	"fmt"
	"github.com/Akishleroy/go-bookstore/pkg/jwt"
)

type LoginInput struct {
	UserName string
	Password string
}

func (l *LoginInput) Login() (User, string, error) {
	var getUser User
	fmt.Println(l.UserName, l.Password)
	db.Where("user_name=? AND password=?", l.UserName, l.Password).Find(&getUser)
	//fmt.Println(getUser.UserName, getUser.Password)
	if getUser.UserName == l.UserName && l.Password == getUser.Password {
		token, err := jwt.CreateToken(getUser.ID)
		tokenModel := &Token{
			UserID: getUser.ID,
			Token:  token,
		}
		_ = tokenModel.InsertToken()
		return getUser, token, err
	}
	return getUser, "error1", nil

}
