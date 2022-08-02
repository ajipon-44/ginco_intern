package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
		"net/http"
		"encoding/json"
		"github.com/gorilla/mux"
		"github.com/jinzhu/gorm"
		"log"
		"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the sql driver.\n")
}

type User struct {
	Id int
	Name string
}

func connectDB() *gorm.DB {
	dbconf := "root:Yanakei727@tcp(127.0.0.1:3306)/ginco_intern"
	db, err_db := gorm.Open("mysql", dbconf)
	if err_db != nil {
		panic(err_db.Error())
	}

	return db
}

func createUser(w http.ResponseWriter, r *http.Request){
	db := connectDB()
	defer db.Close()

	body, err_body := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
  if err_body != nil {
		panic(err_body)
	}

	var user User

	err_res := json.Unmarshal(body, &user)
	if err_res != nil {
		panic(err_res.Error())
	}

	db.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func getUsers(w http.ResponseWriter, r *http.Request){
	db := connectDB()
	defer db.Close()

	var users []User

	db.Find(&users)

	response, err_res := json.Marshal(users)

	if err_res != nil {
		panic(err_res.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func updateUser(w http.ResponseWriter, r *http.Request){
	db := connectDB()
	defer db.Close()

	body, err_body := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
  if err_body != nil {
		panic(err_body)
	}

	var user User

	err_res := json.Unmarshal(body, &user)
	if err_res != nil {
		panic(err_res.Error())
	}

	db.Save(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler)
	r.HandleFunc("/user/get", getUsers).Methods("GET")
	r.HandleFunc("/user/create", createUser).Methods("POST")
	r.HandleFunc("/user/update", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
