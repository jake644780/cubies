package main

import (
	"cliDemo/models"
	"fmt"
	"math/rand"
	"time"
)

var cards []models.Card

func main(){

	p1 := models.Player{Health: 30}
	p2 := models.Player{Health: 30}

	rand.Seed(time.Now().UnixNano())
	fmt.Println("game created!");
	for i := 1; i <= 10; i++ {
		cards = append(cards, models.Card{Cost: i,Health: i,Damage: i,})
	}

	//for now, there will only be 10, x/x/x type cards, and both decks will have 30 random of these


	for i := 0; i < 30; i++ {
		randomNumber := rand.Intn(10)
		p1.Deck = append(p1.Deck, cards[randomNumber]);
		p2.Deck = append(p2.Deck, cards[randomNumber]);
	}

	fmt.Printf("player1: %v", p1.Health)
	fmt.Printf("\nplayer1 deck: %v", p1.Deck)
	fmt.Printf("\nplayer2: %v", p2.Health)
	fmt.Printf("\nplayer2 deck: %v", p2.Deck)

	for{
		
	}

}