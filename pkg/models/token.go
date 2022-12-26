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
	db = config.GetDB()
	db.AutoMigrate(&Token{})
}

func (t *Token) InsertToken() error {
	db.NewRecord(t)
	db.Create(&t)
	return nil
}

func GetTokenById(Id int64) (*Token, *gorm.DB) {
	var token Token
	db := db.Where("ID=?", Id).Find(&token)
	return &token, db
}

func DeleteToken(ID int64) Token {
	var token Token
	db.Where("ID=?", ID).Delete(token)
	return token
}
