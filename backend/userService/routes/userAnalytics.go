package routes

import (
	"userService/controllers"

	"github.com/gorilla/mux"
)

func UserAnalytics(r *mux.Router) {
	r.HandleFunc("/api/users/{id:[0-9]+}/analytics", controllers.GetUserAnalytics).Methods("GET")
}
