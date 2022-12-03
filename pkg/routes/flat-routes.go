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
	router.HandleFunc("/flat/{flatId}", controllers.UpdateFlat).Methods("PUT")
	router.HandleFunc("/flat/{flatId}", controllers.DeleteFlat).Methods("DELETE")
	//user
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
	//login

}
