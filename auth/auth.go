package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
)

type AuthToken struct {
	Token string `json:"token"`
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request, user_id int) {
	claims := jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte("SECRET_KEY"))
	returnToken := AuthToken{}
	returnToken.Token = tokenString

	res, err_res := json.Marshal(returnToken)
	if err_res != nil {
		panic(err_res)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func VerifyToken(w http.ResponseWriter, r *http.Request) interface{} {
	tokenString := string(r.Header.Get("Authorization"))
	var Id interface{}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		Id = token.Claims.(jwt.MapClaims)["user_id"]
		return []byte("SECRET_KEY"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user_id"])
		fmt.Printf("exp: %v\n", int64(claims["exp"].(float64)))
	} else {
		fmt.Println(err)
	}
	return Id
}
