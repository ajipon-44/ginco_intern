package character

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/minguu42/myapp/share"
	"local.packages/auth"
)

type Results struct {
	List []Result `json:"results"`
}

type Result struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}

type operation struct {
	Times   int `json:"times"`
	GachaID int `json:"gachaID"`
}

type UserCharacter struct {
	UserID      string
	CharacterID string
}

func GachaDraw(w http.ResponseWriter, r *http.Request) {
	id := auth.VerifyToken(w, r)

	db := share.ConnectDb()
	defer db.Close()

	var characters []Character
	var characterRate = []float64{}

	body, err_body := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err_body != nil {
		code, msg := share.StatusCode(err_body)
		http.Error(w, msg, code)
		return
	}

	var operation operation

	err_json := json.Unmarshal(body, &operation)
	if err_json != nil {
		code, msg := share.StatusCode(err_json)
		http.Error(w, msg, code)
		return
	}

	db.Where("gacha_id = ?", operation.GachaID).Find(&characters)

	for _, c := range characters {
		characterRate = append(characterRate, c.Rate)
	}
	sort.Float64s(characterRate)

	boundariesFloat := make([]float64, len(characterRate))
	boundariesInt := make([]int, len(characterRate))

	for i := 0; i < len(characterRate); i++ {
		if i == 0 {
			boundariesFloat[i] = (boundariesFloat[i] + characterRate[i]) * 100
		} else {
			boundariesFloat[i] = boundariesFloat[i-1] + characterRate[i]*100
		}
		boundariesInt[i] = int(boundariesFloat[i])
	}

	rand.Seed(time.Now().UnixNano())

	result := []Result{}
	results := Results{}

	for i := 0; i < operation.Times; i++ {
		draw := rand.Intn(100) + 1
		for i, boundary := range boundariesInt {
			if draw <= boundary {
				gachaResult := Result{strconv.Itoa(characters[i].CharacterID), characters[i].Name}
				result = append(result, gachaResult)
				user_character := UserCharacter{UserID: strconv.Itoa(id), CharacterID: strconv.Itoa(characters[i].CharacterID)}
				db.Create(&user_character)
				break
			}
		}
	}
	results.List = result

	res, err_res := json.Marshal(results)
	if err_res != nil {
		code, msg := share.StatusCode(err_res)
		http.Error(w, msg, code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
