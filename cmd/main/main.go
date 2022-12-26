package main

import (
	"github.com/gorilla/securecookie"
	"log"
	"net/http"

	"github.com/Akishleroy/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"

	"github.com/Akishleroy/go-bookstore/jwt"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func main() {
	r := mux.NewRouter()
	routes.RegisterFlatRoutes(r)
	http.Handle("/", r)
	server, err := &config.serverTokens{
		accessSecret: "secretkey123"
		accessTTL: 86400
	}
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
