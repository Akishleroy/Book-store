package models

import (
	"github.com/Akishleroy/go-bookstore/pkg/config"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserType  string `json:"usertype"`
	UserName  string `json:"username"`
	Booking   []Booking
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.FirstName, validation.Required),
		validation.Field(&u.LastName, validation.Required),
		validation.Field(&u.Password, validation.Required),
	)
}

func (u *User) CreateUser() error {
	err := u.Validate()
	if err != nil {
		return err
	} else {
		db.NewRecord(u)
		db.Create(&u)
		return err
	}
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func Login(Email string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("Email=?", Email).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}
