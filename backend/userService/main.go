package main

import (
	"fmt"
	"net/http"
	"userService/routes"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func main() {
	router := mux.NewRouter()

	routes.UserProfile(router)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", router)
}
