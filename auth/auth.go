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

func VerifyToken(w http.ResponseWriter, r *http.Request) int {
	// ヘッダーからトークンを取り出す
	tokenString := string(r.Header.Get("Authorization"))

	// 取り出したトークンからParseで検証を行う
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// ヘッダーに入っている署名方法を使った署名方法に型アサーションして真偽値を出す
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// falseならエラーを返す
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 秘密鍵のバイト配列を返す
		return []byte("SECRET_KEY"), nil
	})

	// 検証して出力されたデータを格納し、その返り値とtoken.Validに入っている真偽値(トークンの正当性がどうだったか)がどちらも真であった場合結果を出力
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user_id"])
		fmt.Printf("exp: %v\n", int64(claims["exp"].(float64)))
	} else {
		fmt.Println(err)
	}
	// ユーザーIDを返す
	return int(token.Claims.(jwt.MapClaims)["user_id"].(float64))
}
