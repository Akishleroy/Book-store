package routes

import (
	"github.com/Akishleroy/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	//flat
	router.HandleFunc("/flat", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/flat", controllers.GetBook).Methods("GET")
	router.HandleFunc("/flat/{flatId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/flat/{flatId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/flat/{flatId}", controllers.DeleteBook).Methods("DELETE")
	//user
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
}
