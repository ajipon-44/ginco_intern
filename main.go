package main

import (
    _ "github.com/go-sql-driver/mysql"
		"net/http"
		"github.com/gorilla/mux"
		"log"
)

type User struct {
	Id int
	Name string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/get", GetUsers).Methods("GET")
	r.HandleFunc("/user/create", CreateUser).Methods("POST")
	r.HandleFunc("/user/update", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
