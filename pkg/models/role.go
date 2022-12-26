package models

import (
	"github.com/Akishleroy/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Type string `json:"type"`
}

func (r *Role) CreateRole() {
	db.NewRecord(r)
	db.Create(&r)
}

func init() {
	var count int64
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Role{})
	db.Model(&Role{}).Count(&count)

	if count != 3 {
		db.Delete(&Role{})
		type1 := Role{Type: "admin"}
		type2 := Role{Type: "customer"}
		type3 := Role{Type: "seller"}
		db.Create(&type1)
		db.Create(&type2)
		db.Create(&type3)
	}
}
