package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Akishleroy/go-bookstore/pkg/jwt"
	"github.com/Akishleroy/go-bookstore/pkg/models"
	"github.com/Akishleroy/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"log"
	"time"

	//"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)
//var db *gorm.DB
var NewBooking models.Booking

func GetBooking(w http.ResponseWriter, r *http.Request) {
	newBooking := models.GetAllBookings()
	res, _ := json.Marshal(newBooking)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingId := vars["bookingId"]
	ID, err := strconv.ParseInt(bookingId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookingDetails, _ := models.GetBookingById(ID)
	res, _ := json.Marshal(bookingDetails)
	w.Header().Set("Content-type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookingByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookingDetails := models.GetAllBookingsByUserId(userID)
	res, _ := json.Marshal(bookingDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	CreateBooking := &models.Booking{}
	utils.ParseBody(r, CreateBooking)
	startDate, _ := time.Parse("2006-01-02", CreateBooking.StartDate)
	endDate, _ := time.Parse("2006-01-02", CreateBooking.EndDate)
	fmt.Println(startDate)
	fmt.Println(CreateBooking.StartDate)
	vars := mux.Vars(r)
	flatId := vars["flatId"]
	var header = r.Header.Get("Authorization")

	userId, _, _ := jwt.ExtractToken(header[7:])
	flatID, err := strconv.ParseInt(flatId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	var bookings []models.Booking

	models.Db.Where("flat_id=?", flatID).Find(&bookings)
	fmt.Println(bookings)
	if (len(bookings) != 0){
		for i := 0; i < len(bookings); i++ {
			start_date, _:= time.Parse("2006-01-02", bookings[i].StartDate)
			end_date, _:= time.Parse("2006-01-02", bookings[i].StartDate)
			if (start_date.After(startDate) && end_date.Before(startDate)) {
				log.Fatalf("Such flat is booked on that date %s", bookings[i].StartDate)
			}
			if (start_date.After(endDate) && end_date.Before(endDate)) {
				log.Fatalf("Such flat is booked on that date %s", bookings[i].EndDate)
			}
		}
	} else {
		CreateBooking.UserId = userId
		CreateBooking.FlatId = uint(flatID)
		b := CreateBooking.CreateBooking()
		res, _ := json.Marshal(b)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}


}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingId := vars["bookingId"]
	ID, err := strconv.ParseInt(bookingId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	booking := models.DeleteBooking(ID)
	res, _ := json.Marshal(booking)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	var updateBooking = &models.Booking{}
	utils.ParseBody(r, updateBooking)
	vars := mux.Vars(r)
	bookingId := vars["bookingId"]
	ID, err := strconv.ParseInt(bookingId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookingDetails, db := models.GetBookingById(ID)
	bookingDetails.UserId = updateBooking.UserId
	bookingDetails.FlatId = updateBooking.FlatId
	bookingDetails.StartDate = updateBooking.StartDate
	bookingDetails.EndDate = updateBooking.EndDate

	db.Save(&bookingDetails)
	res, _ := json.Marshal(bookingDetails)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
