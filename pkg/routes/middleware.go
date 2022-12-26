package routes

import (
	"github.com/Akishleroy/go-bookstore/pkg/jwt"
	"net/http"
)

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")

		userId, _ := jwt.ExtractToken(header[7:])

		if userId == 0 {
			http.Redirect(w, r, "/user/login", 302)
		}

		HandlerFunc.ServeHTTP(w, r)
	}
}
