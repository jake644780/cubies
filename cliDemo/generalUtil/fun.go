package generalutil

import (
	"cliDemo/models"
	"errors"
)



func DrawCard(cardList *[]models.Card) (models.Card, error){
	if len(*cardList) == 0{
		return models.Card{}, errors.New("deck is empty")
	}
	returnCard := (*cardList)[len(*cardList)-1]
	*cardList = (*cardList)[:len(*cardList)-1]
	return returnCard, nil
}

func DrawCards(cardList *[]models.Card, n int) ([]models.Card, error){
	var returnList []models.Card
	for i := 0; i < n; i++ {
		if len(*cardList) == 0{
			return returnList, errors.New("deck is empty")
		}
		returnList = append(returnList, (*cardList)[len(*cardList)-1])
		*cardList = (*cardList)[:len(*cardList)-1]
	}
	return returnList, nil
}

