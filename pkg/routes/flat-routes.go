package routes

import (
	"github.com/Akishleroy/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterFlatRoutes = func(router *mux.Router) {
	//flat
	router.HandleFunc("/flat", controllers.CreateFlat).Methods("POST")
	router.HandleFunc("/flat", controllers.GetFlat).Methods("GET")
	router.HandleFunc("/flat/{flatId}", controllers.GetFlatById).Methods("GET")
	router.HandleFunc("/flats/filter", controllers.GetFlatsByFilter).Methods("GET")
	router.HandleFunc("/flat/{flatId}", controllers.UpdateFlat).Methods("PUT")
	router.HandleFunc("/flat/{flatId}", controllers.DeleteFlat).Methods("DELETE")

	//user
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
	//booking
	router.HandleFunc("/booking/{flatId}", Auth(controllers.CreateBooking)).Methods("POST")
	router.HandleFunc("/booking", Auth(controllers.GetBooking)).Methods("GET")
	router.HandleFunc("/booking/{bookingId}", Auth(controllers.GetBookingById)).Methods("GET")
	router.HandleFunc("/bookings/{userId}", Auth(controllers.GetBookingByUserId)).Methods("GET")
	router.HandleFunc("/booking/{bookingId}", Auth(controllers.UpdateBooking)).Methods("PUT")
	router.HandleFunc("/booking/{bookingId}", Auth(controllers.DeleteBooking)).Methods("DELETE")
	//login
}
