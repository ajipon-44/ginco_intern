package main

import (
    //"database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
		"net/http"
		"encoding/json"
		"github.com/gorilla/mux"
		"github.com/jinzhu/gorm"
		_ "github.com/jinzhu/gorm/dialects/mysql"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the sql driver.\n")
}

type User struct {
	id int
	name string
}

func getUsers(w http.ResponseWriter, r *http.Request){
	dbconf := "root:Yanakei727@tcp(127.0.0.1:3306)/ginco_intern"
	db, _ := gorm.Open("mysql", dbconf)
	defer db.Close()

	var users []User

	result := db.Find(&users)

	response, _ := json.Marshal(users)

	fmt.Println(users)
	fmt.Println(result.RowsAffected)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler)
	r.HandleFunc("/users", getUsers).Methods("GET")

	http.ListenAndServe(":8000", r)
}
