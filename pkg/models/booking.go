package models

import (
	"fmt"
	"github.com/Akishleroy/go-bookstore/pkg/config"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type Booking struct {
	gorm.Model
	UserId uint       `json:"userid"`
	FlatId uint       `json:"flatid"`
	StartDate   string `json:"start_date"`
	EndDate string `json:"end_date"`
}

func (b Booking) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.StartDate, validation.Required),
		validation.Field(&b.EndDate, validation.Required),
	)
}

func init() {
	config.Connect()
	Db = config.GetDB()
	Db.AutoMigrate(&Booking{})
}

func (b *Booking) CreateBooking() error {
	err := b.Validate()
	if err != nil {
		return err
	} else {
		fmt.Printf("err = %T\n", err)
		Db.NewRecord(b)
		Db.Create(&b)
		return err
	}

}

func GetAllBookings() []Booking {
	var Bookings []Booking
	Db.Find(&Bookings)
	return Bookings
}

func GetAllBookingsByUserId(Id int64) []Booking {
	var getBooking []Booking
	Db.Where("user_id=?", Id).Find(&getBooking)
	return getBooking
}

func GetBookingById(Id int64) (*Booking, *gorm.DB) {
	var getBooking Booking
	db := Db.Where("ID=?", Id).Find(&getBooking)
	return &getBooking, db
}

func DeleteBooking(ID int64) Booking {
	var book Booking
	Db.Where("ID=?", ID).Delete(book)
	return book
}
