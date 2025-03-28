package models


type Player struct {
	//general
	Username string
	InGame bool
	GameId int
	Hand   []Card
	Deck   []Card
	Health int
}

func MakePlayer() Player{
	return Player{
		Username: "",
		InGame: false,
		GameId: 0,
		Hand: []Card{},
		Deck: []Card{},
		Health: 0,
	}
}