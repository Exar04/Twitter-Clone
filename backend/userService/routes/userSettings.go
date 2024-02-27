package routes

import (
	"userService/controllers"

	"github.com/gorilla/mux"
)

func UserSettings(r *mux.Router) {
	r.HandleFunc("/user/settings", controllers.GetUserSettings).Methods("GET")
	r.HandleFunc("/user/settings", controllers.UpdateUserSettings).Methods("PUT")
}
