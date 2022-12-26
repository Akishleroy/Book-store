package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	tokens       map[int64]*Token
	accessSecret string
	accessTTL    int
}

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Htbo8731*@tcp(localhost:3306)/world?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
