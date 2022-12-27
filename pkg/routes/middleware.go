package routes

import (
	"fmt"
	"github.com/Akishleroy/go-bookstore/pkg/jwt"
	"net/http"
)

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")

		userId, userType, _ := jwt.ExtractToken(header[7:])
		fmt.Println("userType")
		fmt.Println(userType)

		if userId == 0 {
			http.Redirect(w, r, "/user/login", 302)
		}

		HandlerFunc.ServeHTTP(w, r)
	}
}

func canEdit(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")

		_ , userType, _ := jwt.ExtractToken(header[7:])
		fmt.Println("userType")
		fmt.Println(userType)

		if userType == 2 {
			http.Redirect(w, r, "/flat", 302)
		} else {
			HandlerFunc.ServeHTTP(w, r)
		}
	}
}
