package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Akishleroy/go-bookstore/pkg/models"
	"github.com/Akishleroy/go-bookstore/pkg/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.LoginInput{}
	utils.ParseBody(r, user)

	l, _, err := user.Login()
	if err != nil {
		fmt.Println("error2")
	}
	res, _ := json.Marshal(l)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
