package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user", HomeHandler).Methods("GET")

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", router)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hehe"))
}
