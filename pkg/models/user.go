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
	UserType  uint    `json:"usertype"`
	UserName  string `json:"username"`
}

func init() {
	config.Connect()
	Db = config.GetDB()
	Db.AutoMigrate(&User{})
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.FirstName, validation.Required),
		validation.Field(&u.LastName, validation.Required),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.UserType, validation.Required, validation.Min(1), validation.Max(3)),
	)
}

func (u *User) RegisterNewUser() error {
	err := u.Validate()
	if err != nil {
		return err
	} else {
		Db.NewRecord(u)
		Db.Create(&u)
		return err
	}
}

func GetAllUsers() []User {
	var Users []User
	Db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := Db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func (u *User) Login() string {
	var getUser User
	Db.Where("user_name=? AND password=?", u.UserName, u.Password).Find(&getUser)
	if getUser.UserName == u.UserName && u.Password == getUser.Password {
		return "successfully logged in"
	}
	return "Incorrect password or username"
}

func DeleteUser(ID int64) User {
	var user User
	Db.Where("ID=?", ID).Delete(user)
	return user
}
