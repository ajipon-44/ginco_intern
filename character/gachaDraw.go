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
	Times int `json:"times"`
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

	db.Find(&characters)

	body, err_body := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err_body != nil {
		panic(err_body)
	}

	var operation operation

	err_json := json.Unmarshal(body, &operation)
	if err_json != nil {
		panic(err_json.Error())
	}

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
				gachaResult := Result{strconv.Itoa(characters[i].Id), characters[i].Name}
				result = append(result, gachaResult)
				user_character := UserCharacter{UserID: strconv.Itoa(int(id.(float64))), CharacterID: strconv.Itoa(characters[i].Id)}
				db.Create(&user_character)
				break
			}
		}
	}
	results.List = result

	res, err_res := json.Marshal(results)
	if err_res != nil {
		panic(err_res)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
