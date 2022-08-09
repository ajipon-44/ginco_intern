package character

type Character struct {
	Id          int
	CharacterID int `json:"characterID"`
	Name        string
	Rate        float64
}
