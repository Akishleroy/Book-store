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
	//fmt.Println(user.UserName, user.Password)
	l, token, err := user.Login()
	if err != nil {
		fmt.Println("error2")
	}
	res, _ := json.Marshal(l)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Println(token)
}

//
//import (
//	"errors"
//	"github.com/Akishleroy/go-bookstore/pkg/config"
//	"github.com/Akishleroy/go-bookstore/pkg/models"
//	"golang.org/x/crypto/bcrypt"
//	"html/template"
//	"net/http"
//)
//
//type UserInput struct {
//	Username string `validate:"required"`
//	Password string `validate:"required"`
//}
//
//var userModel = models.NewUserModel()
//var validation = libraries.NewValidation()
//
//func Index(w http.ResponseWriter, r *http.Request) {
//
//	session, _ := config.Store.Get(r, config.SESSION_ID)
//
//	if len(session.Values) == 0 {
//		http.Redirect(w, r, "/login", http.StatusSeeOther)
//	} else {
//
//		if session.Values["loggedIn"] != true {
//			http.Redirect(w, r, "/login", http.StatusSeeOther)
//		} else {
//
//			data := map[string]interface{}{
//				"username": session.Values["username"],
//			}
//
//			temp, _ := template.ParseFiles("views/index.html")
//			temp.Execute(w, data)
//		}
//
//	}
//}
//
//func Login(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method == http.MethodGet {
//		temp, _ := template.ParseFiles("views/login.html")
//		temp.Execute(w, nil)
//	} else if r.Method == http.MethodPost {
//		// proses login
//		r.ParseForm()
//		UserInput := &UserInput{
//			Username: r.Form.Get("username"),
//			Password: r.Form.Get("password"),
//		}
//
//		errorMessages := validation.Struct(UserInput)
//
//		if errorMessages != nil {
//
//			data := map[string]interface{}{
//				"validation": errorMessages,
//			}
//
//			temp, _ := template.ParseFiles("views/login.html")
//			temp.Execute(w, data)
//
//		} else {
//
//			var user models.User
//			db.Where(&user, "username", UserInput.Username)
//
//			var message error
//			if user.Email == "" {
//				message = errors.New("Email is reuired field!")
//			} else {
//				errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
//				if errPassword != nil {
//					message = errors.New("Password is required field!")
//				}
//			}
//
//			if message != nil {
//
//				data := map[string]interface{}{
//					"error": message,
//				}
//
//				temp, _ := template.ParseFiles("views/login.html")
//				temp.Execute(w, data)
//			} else {
//				// set session
//				session, _ := config.Store.Get(r, config.SESSION_ID)
//
//				session.Values["loggedIn"] = true
//				session.Values["email"] = user.Email
//				session.Values["username"] = user.UserName
//				session.Values["firstname"] = user.FirstName
//
//				session.Save(r, w)
//
//				http.Redirect(w, r, "/", http.StatusSeeOther)
//			}
//		}
//
//	}
//
//}
//
//func Logout(w http.ResponseWriter, r *http.Request) {
//	session, _ := config.Store.Get(r, config.SESSION_ID)
//	// delete session
//	session.Options.MaxAge = -1
//	session.Save(r, w)
//
//	http.Redirect(w, r, "/login", http.StatusSeeOther)
//}
//
//func Register(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method == http.MethodGet {
//
//		temp, _ := template.ParseFiles("views/register.html")
//		temp.Execute(w, nil)
//
//	} else if r.Method == http.MethodPost {
//		// melakukan proses registrasi
//
//		// mengambil inputan form
//		r.ParseForm()
//
//		user := models.User{
//			FirstName: r.Form.Get("nama_lengkap"),
//			Email:     r.Form.Get("email"),
//			UserName:  r.Form.Get("username"),
//			Password:  r.Form.Get("password"),
//		}
//
//		errorMessages := validation.Struct(user)
//
//		if errorMessages != nil {
//
//			data := map[string]interface{}{
//				"validation": errorMessages,
//				"user":       user,
//			}
//
//			temp, _ := template.ParseFiles("views/register.html")
//			temp.Execute(w, data)
//		} else {
//
//			// hashPassword
//			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
//			user.Password = string(hashPassword)
//
//			// insert ke database
//			user.CreateUser()
//
//			data := map[string]interface{}{
//				"pesan": "Registrasi berhasil",
//			}
//			temp, _ := template.ParseFiles("views/register.html")
//			temp.Execute(w, data)
//		}
//	}
//
//}
