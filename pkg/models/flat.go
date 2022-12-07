package models

import (
	"fmt"
	"github.com/Akishleroy/go-bookstore/pkg/config"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Flat struct {
	gorm.Model
	Price    float64 `json:"price"`
	Size     int     `json:"size"`
	Address  string  `json:"address"`
	City     string  `json:"city"`
	IsActive bool    `json:"isActive"`
	Booking  []Booking
}

func (b Flat) Validate() error {
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
	db.AutoMigrate(&Flat{})
}

func (b *Flat) CreateFlat() error {
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

func GetAllFlats() []Flat {
	var Flats []Flat
	db.Find(&Flats)
	return Flats

}

func GetFlatById(Id int64) (*Flat, *gorm.DB) {
	var getFlat Flat
	db := db.Where("ID=?", Id).Find(&getFlat)
	return &getFlat, db
}

func DeleteFlat(ID int64) Flat {
	var book Flat
	db.Where("ID=?", ID).Delete(book)
	return book
}

func FilterByCity(city string) []Flat {
	var Flats []Flat
	db.Where("city=?", city).Find(&Flats)
	return Flats
}

func GetFlatsByFilter(params [5]string) []Flat {
	var getFlats []Flat
	if params[0] != "" {
		city := params[0]
		db.Where("city=?", city).Find(&getFlats)
		return getFlats
	}
	if params[1] != "" && params[2] != "" {
		fromSize := params[1]
		toSize := params[2]
		db.Where("size>=? AND size<=?", fromSize, toSize).Find(&getFlats)
		return getFlats
	}
	if params[1] != "" && params[2] == "" {
		fromSize := params[1]
		db.Where("size>=?", fromSize).Find(&getFlats)
		return getFlats
	}
	if params[1] == "" && params[2] != "" {
		toSize := params[2]
		db.Where("size<=?", toSize).Find(&getFlats)
		return getFlats
	}
	if params[3] != "" && params[4] != "" {
		fromPrice := params[3]
		toPrice := params[4]
		db.Where("price>=? AND price<=?", fromPrice, toPrice).Find(&getFlats)
		return getFlats
	}
	if params[3] != "" && params[4] == "" {
		fromPrice := params[3]
		db.Where("price>=?", fromPrice).Find(&getFlats)
		return getFlats
	}
	if params[3] == "" && params[4] != "" {
		toPrice := params[4]
		db.Where("price<=?", toPrice).Find(&getFlats)
		return getFlats
	}
	return getFlats
}
