package models

import (
	"fmt"
	"github.com/Akishleroy/go-bookstore/pkg/config"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	UserID int       `json:"userid"`
	FlatId int       `json:"flatid"`
	Time   time.Time `json:"timeid"`
	//User   *User     // поменять json
	//Flat   *Flat     // поменять json
}

func (b Booking) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.UserID, validation.Required),
		validation.Field(&b.FlatId, validation.Required),
		//validation.Field(&b.Time, validation.Required),
	)
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Booking{})
}

func (b *Booking) CreateBooking() error {
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

func GetAllBookings() []Booking {
	var Bookings []Booking
	db.Find(&Bookings)
	return Bookings
}

func GetAllBookingsByUserId(Id int64) []Booking {
	var getBooking []Booking
	db.Where("user_id=?", Id).Find(&getBooking)
	return getBooking
}

func GetBookingById(Id int64) (*Booking, *gorm.DB) {
	var getBooking Booking
	db := db.Where("ID=?", Id).Find(&getBooking)
	return &getBooking, db
}

func DeleteBooking(ID int64) Booking {
	var book Booking
	db.Where("ID=?", ID).Delete(book)
	return book
}
