package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/minguu42/myapp/share"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := share.ConnectDb()
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
