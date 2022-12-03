package models

import (
	"fmt"
	"github.com/Akishleroy/go-bookstore/pkg/config"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Price    float64 `json:"price"`
	Size     int     `json:"size"`
	Address  string  `json:"address"`
	City     string  `json:"city"`
	IsActive bool    `json:"isActive"`
}

func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Price, validation.Required),
		validation.Field(&b.Size, validation.Required, validation.Min(1)),
		validation.Field(&b.Address, validation.Required),
		validation.Field(&b.City, validation.Required),
	)
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() error {
	err := b.Validate()
	if err != nil {
		return err
	} else {
		fmt.Printf("err = %T\n", err)
		db.NewRecord(b)
		db.Create(&b)
		return err
	}

}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
