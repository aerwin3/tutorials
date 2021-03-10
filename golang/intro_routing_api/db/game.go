package db

import "errors"

type Game struct {
	Teams  []*Team
	Id     string         `json:"id"`
	Plays  []Play         `json:"plays"`
	Scores map[string]int `json:"scores"`
	Winner string         `json:"winner"`
}

type Play struct {
	Player *Player `json:"player"`
	Time   string  `json:"created_at"`
}

func GetGames() (gs []Game) {
	for _, g := range games {
		gs = append(gs, *g)
	}
	return
}

func GetGame(id string) (err error, g Game) {
	for _, v := range games {
		if v.Id == id {
			g = *v
			return
		}
	}
	err = errors.New("not found")
	return
}
