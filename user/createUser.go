package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/minguu42/myapp/share"
	"local.packages/auth"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := share.ConnectDb()
	defer db.Close()

	body, err_body := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	if err_body != nil {
		code, msg := share.StatusCode(err_body)
		http.Error(w, msg, code)
		return
	}

	var user User

	err_res := json.Unmarshal(body, &user)
	if err_res != nil {
		code, msg := share.StatusCode(err_res)
		http.Error(w, msg, code)
		return
	}

	if user.Name == "" {
		code, msg := share.StatusCode(err_res)
		http.Error(w, msg, code)
	} else {
		//db.Create(&user)
		auth.GetTokenHandler(w, r, user.Id)
	}
}
