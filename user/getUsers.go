package user

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/minguu42/myapp/share"
	"local.packages/auth"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	id := auth.VerifyToken(w, r)

	db := share.ConnectDb()
	defer db.Close()

	var user User

	db.Select("name").Where("id = ?", id).Find(&user)

	response, err_res := json.Marshal(user)
	if err_res != nil {
		panic(err_res.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
