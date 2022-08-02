package main

import (
    _ "github.com/go-sql-driver/mysql"
		"net/http"
		"encoding/json"
		"io/ioutil"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	db := ConnectDb()
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
