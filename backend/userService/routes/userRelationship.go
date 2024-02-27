package routes

import (
	"userService/controllers"

	"github.com/gorilla/mux"
)

func UserRelationship(r *mux.Router) {
	r.HandleFunc("/user/{id:[0-9]+}/follow", controllers.FollowUser).Methods("POST")
	r.HandleFunc("/user/{id:[0-9]+}/unfollow", controllers.UnfollowUser).Methods("POST")
	r.HandleFunc("/user/{id:[0-9]+}/followers", controllers.GetUserFollowers).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}/following", controllers.GetUserFollowing).Methods("GET")
}
