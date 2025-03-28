package models

import "github.com/gorilla/websocket"

type Player struct {
	Hand   []Card
	Deck   []Card
	Health int
	Conn *websocket.Conn
	InGame bool
	GameId int
}
