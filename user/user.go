package user

type User struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
}
