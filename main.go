package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/minguu42/myapp/user"
	"github.com/minguu42/myapp/userCharacter"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/get", user.GetUsers).Methods("GET")
	r.HandleFunc("/user/create", user.CreateUser).Methods("POST")
	r.HandleFunc("/user/update", user.UpdateUser).Methods("PUT")

	r.HandleFunc("/character/list", userCharacter.CharacterList).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
