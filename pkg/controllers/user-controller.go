package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Akishleroy/go-bookstore/pkg/models"
	"github.com/Akishleroy/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

var NewUser models.User

//var tmpl = template.Must(template.ParseGlob("templates/"))

type ErrorResponse struct {
	Err string
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.GetAllUsers()
	res, _ := json.Marshal(newUser)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("eror while parsing")
	}
	userDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	pass, err := bcrypt.GenerateFromPassword([]byte(CreateUser.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	CreateUser.Password = string(pass)
	utils.ParseBody(r, CreateUser)
	u := CreateUser.CreateUser()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, db := models.GetUserById(ID)
	if updateUser.FirstName != "" {
		userDetails.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		userDetails.LastName = updateUser.LastName
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if updateUser.Password != "" {
		userDetails.Password = updateUser.Password
	}
	if updateUser.UserType != "" {
		userDetails.UserType = updateUser.UserType
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//func Login(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	username := r.Form.Get("username")
//	p := r.Form.Get("password")
//	password := GetMD5Hash(p)
//
//	db := config.GetDB()
//	var id int
//	var getUser models.User
//	row := db.Where("username=? AND password=?", username, password).Find(&getUser)
//	switch err := row.Scan(&id); err {
//	case sql.ErrNoRows:
//		tmpl.ExecuteTemplate(w, "Index", nil)
//	default:
//		tmpl.ExecuteTemplate(w, "Success", nil)
//	}
//	defer db.Close()
//}
//func GetMD5Hash(text string) string {
//	hasher := md5.New()
//	hasher.Write([]byte(text))
//	return hex.EncodeToString(hasher.Sum(nil))
//}

//
//func Login(w http.ResponseWriter, r *http.Request) {
//	user := &models.User{}
//	err := json.NewDecoder(r.Body).Decode(user)
//	if err != nil {
//		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
//		json.NewEncoder(w).Encode(resp)
//		return
//	}
//	resp := FindOne(user.Email, user.Password)
//	json.NewEncoder(w).Encode(resp)
//}
//
//func FindOne(email, password string) map[string]interface{} {
//	user := &models.User{}
//
//	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
//		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
//		return resp
//	}
//	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
//
//	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
//	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
//		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
//		return resp
//	}
//
//	tk := &models.Token{
//		ID:        user.ID,
//		firstName: user.FirstName,
//		lastName:  user.LastName,
//		email:     user.Email,
//		StandardClaims: &jwt.StandardClaims{
//			ExpiresAt: expiresAt,
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
//
//	tokenString, error := token.SignedString([]byte("secret"))
//	if error != nil {
//		fmt.Println(error)
//	}
//
//	var resp = map[string]interface{}{"status": false, "message": "logged in"}
//	resp["token"] = tokenString //Store the token in the response
//	resp["user"] = user
//	return resp
//}

//func loginHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "POST" {
//		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
//		return
//	}
//	// ParseForm parses the raw query from the URL and updates r.Form
//	err := r.ParseForm()
//	if err != nil {
//		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
//		return
//	}
//	users := &models.User{}
//	// Get username and password from the parsed form
//	email := r.Form.Get("email")
//	password := r.Form.Get("password")
//
//	// Check if user exists
//	storedPassword, exists := users[email]
//	if exists {
//		// It returns a new session if the sessions doesn't exist
//		session, _ := store.Get(r, "session.id")
//		if storedPassword == password {
//			session.Values["authenticated"] = true
//			// Saves all sessions used during the current request
//			session.Save(r, w)
//		} else {
//			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
//		}
//		w.Write([]byte("Login successfully!"))
//	}
//}
