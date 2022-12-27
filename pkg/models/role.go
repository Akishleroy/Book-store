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
	Db.NewRecord(r)
	Db.Create(&r)
}

func init() {
	var count int64
	config.Connect()
	Db = config.GetDB()
	Db.AutoMigrate(&Role{})
	Db.Model(&Role{}).Count(&count)

	if count != 3 {
		Db.Delete(&Role{})
		type1 := Role{Type: "admin"}
		type2 := Role{Type: "customer"}
		type3 := Role{Type: "seller"}
		Db.Create(&type1)
		Db.Create(&type2)
		Db.Create(&type3)
	}
}
