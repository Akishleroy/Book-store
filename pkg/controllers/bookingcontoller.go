package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Akishleroy/go-bookstore/pkg/models"
	"github.com/Akishleroy/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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
	//vars := mux.Vars(r)
	//userId := vars["userId"]
	//flatId := vars["flatId"]
	//flatID, err1 := strconv.ParseInt(flatId, 0, 0)
	//userID, err2 := strconv.ParseInt(userId, 0, 0)
	//if err1 != nil {
	//	fmt.Println("error while parsing")
	//}
	//if err2 != nil {
	//	fmt.Println("error while parsing")
	//}
	//flatDetails, _ := models.GetFlatById(flatID)
	//userDetails, _ := models.GetUserById(userID)
	//CreateBooking.User = userDetails
	//CreateBooking.Flat = flatDetails
	utils.ParseBody(r, CreateBooking)
	b := CreateBooking.CreateBooking()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

//	func CreateBooking(w http.ResponseWriter, r *http.Request) {
//		CreateBooking := &models.Booking{}
//		pass, err := bcrypt.GenerateFromPassword([]byte(CreateBooking.Password), bcrypt.DefaultCost)
//		if err != nil {
//			fmt.Println(err)
//			err := ErrorResponse{
//				Err: "Password Encryption  failed",
//			}
//			json.NewEncoder(w).Encode(err)
//		}
//
//		CreateBooking.Password = string(pass)
//		utils.ParseBody(r, CreateBooking)
//		u := CreateBooking.CreateBooking()
//		res, _ := json.Marshal(u)
//		w.WriteHeader(http.StatusOK)
//		w.Write(res)
//	}
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
	bookingDetails.UserID = updateBooking.UserID
	bookingDetails.FlatId = updateBooking.FlatId
	bookingDetails.Time = updateBooking.Time

	db.Save(&bookingDetails)
	res, _ := json.Marshal(bookingDetails)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
