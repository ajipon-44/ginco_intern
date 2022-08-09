package userCharacter

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/minguu42/myapp/share"
	"local.packages/auth"
)

type Results struct {
	List []Result `json:"characters"`
}

type Result struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
	Name            string `json:"name"`
}

type Character struct {
	id   int
	Name string
	rate float64
}

func CharacterList(w http.ResponseWriter, r *http.Request) {
	id := auth.VerifyToken(w, r)

	db := share.ConnectDb()
	defer db.Close()

	var userCharacters []UserCharacter
	var character []Character

	db.Where("user_id = ?", strconv.Itoa(id)).Find(&userCharacters)
	db.Find(&character)

	result := []Result{}
	results := Results{}

	for _, userCharacter := range userCharacters {
		char_id := userCharacter.CharacterID

		buf := Result{strconv.Itoa(userCharacter.UserCharacterID), strconv.Itoa(char_id), character[char_id-1].Name}
		result = append(result, buf)
	}
	results.List = result

	response, err_res := json.Marshal(results)
	if err_res != nil {
		panic(err_res.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
