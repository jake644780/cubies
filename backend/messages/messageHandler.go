package messages

import (
	"backend/models"
	"encoding/json"
	"fmt"
)

func UnmarshalEntity(data []byte) (models.Entity, error) {
	var temp struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return nil, err
	}

	var entity models.Entity
	switch temp.Type {
	case "SetUsername":
		entity = &models.SetUsernameMessage{}
	case "ChallengePlayer":
		entity = &models.ChallengePlayer{}
	case "AcceptChallenge":
		entity = &models.AcceptChallenge{}
	case "PlayCard":
		entity = &models.PlayCard{}
	case "SetDeck":
		entity = &models.SetDeck{}
	case "Emote":
		entity = &models.Emote{}
	default:
		return nil, fmt.Errorf("unknown message type: %s", temp.Type)
	}

	if err := json.Unmarshal(data, entity); err != nil {
		return nil, err
	}

	return entity, nil
}
