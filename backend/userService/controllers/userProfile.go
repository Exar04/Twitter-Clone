package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Println(userID)

}

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {

}

func UpdateUserProfilePicture(w http.ResponseWriter, r *http.Request) {

}

func UpdateUserProfileBackgroundPicture(w http.ResponseWriter, r *http.Request) {

}
