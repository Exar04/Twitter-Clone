package routes

import (
	"net/http"
	"userService/controllers"

	"github.com/gorilla/mux"
)

func UserProfile(router *mux.Router) {
	router.HandleFunc("user/getProfile/{id}", controllers.GetUserProfile).Methods(http.MethodGet)
	router.HandleFunc("user/updateProfile", controllers.UpdateUserProfile).Methods(http.MethodPut)
}
