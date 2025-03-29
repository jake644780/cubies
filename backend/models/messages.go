package models


// Define a common interface
type Entity interface {
	GetMessageID() string
	GetType() string
}

type Message struct {
	MessageID string `json:"messageID"`
	Type      string `json:"type"` // Identify the message type
}

func (m Message) GetMessageID() string {
	return m.MessageID
}

func (m Message) GetType() string {
	return m.Type
}

type SetUsernameMessage struct {
	Message
	SetUsername string `json:"setUsername"`
}

type ChallengePlayer struct {
	Message
	Challenged string `json:"challengedName"`
}

type AcceptChallenge struct {
	Message
	Response bool `json:"response"`
}

type PlayCard struct {
	Message
	CardId  int `json:"cardID"`
	LaneNum int `json:"laneNum"`
}

type SetDeck struct {
	Message
	Cards []int `json:"cards"`
}

type Emote struct {
	Message
	EmoteId int `json:"emoteId"`
}
