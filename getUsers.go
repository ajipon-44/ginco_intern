package main

import (
    _ "github.com/go-sql-driver/mysql"
		"net/http"
		"encoding/json"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	db := ConnectDb()
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
