package models

import (
	"github.com/Akishleroy/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model
	UserID uint   `json:"userid"`
	Token  string `json:"token"`
}

func init() {
	config.Connect()
	Db = config.GetDB()
	Db.AutoMigrate(&Token{})
}

func (t *Token) InsertToken() error {
	Db.NewRecord(t)
	Db.Create(&t)
	return nil
}

func GetUserByToken(tokenVal string) (*Token, *gorm.DB) {
	var token Token
	db := Db.Where("token=?", tokenVal).Find(&token)
	return &token, db
}

func GetTokenById(Id int64) (*Token, *gorm.DB) {
	var token Token
	db := Db.Where("ID=?", Id).Find(&token)
	return &token, db
}

func DeleteToken(ID int64) Token {
	var token Token
	Db.Where("ID=?", ID).Delete(token)
	return token
}
