package models

import "errors"

type Game struct {
	Turn    int
	Mana    int
	Players []*Player
	Active  bool
	Lanes   []Lane
}

func (g *Game) SetLanes() error {
	if len(g.Players) == 2 {
		for i := 0; i < 4; i++ {
			g.Lanes = append(g.Lanes, Lane{
				Players: [2]*Player{g.Players[0], g.Players[1]}, // Assigning players explicitly
				Cards:   [2]*Card{nil, nil},                     // Empty card slots
			})	
		}
		return nil
	} else if len(g.Players) == 3 {
		for i := 0; i < 2; i++ {
			g.Lanes = append(g.Lanes, 
				Lane{
				Players: [2]*Player{g.Players[0], g.Players[1]},
				Cards:   [2]*Card{nil, nil},
				},
				Lane{
					Players: [2]*Player{g.Players[0], g.Players[2]},
					Cards:   [2]*Card{nil, nil},
				},
				Lane{
					Players: [2]*Player{g.Players[1], g.Players[2]},
					Cards:   [2]*Card{nil, nil},
				},
			)
		}
		return nil
	}
	return errors.New("incorrect amount of players")
}

type Lane struct {
	Players [2]*Player
	Cards   [2]*Card
}