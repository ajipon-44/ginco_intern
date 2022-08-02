package userCharacter

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/minguu42/myapp/share"
)

func CharacterList(w http.ResponseWriter, _ *http.Request) {
	db := share.ConnectDb()
	defer db.Close()

	var userCharacters []UserCharacter

	db.Find(&userCharacters)

	response, err_res := json.Marshal(userCharacters)

	if err_res != nil {
		panic(err_res.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
