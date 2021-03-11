package db

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const (
	winCondition = 100
)

type Game struct {
	players []*Player
	Id      string         `json:"id"`
	Plays   []Play         `json:"-"`
	Scores  map[string]int `json:"scores"`
	Winner  string         `json:"winner"`
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

func CreateGame(teamNames []string) *Game {
	g := &Game{Id: uuid.NewString(), players: []*Player{}, Scores: map[string]int{}}
	for _, name := range teamNames {
		for _, p := range teams[name].Players {
			g.players = append(g.players, p)
		}

		g.Scores[name] = 0
	}
	games = append(games, g)
	return g
}

func (g *Game) GetRandomPlayer() Player {
	i := rand.Intn(len(g.players))
	return players[i]
}

func (g *Game) RunPlay() {
	p := g.GetRandomPlayer()
	g.Scores[p.Team] += 1
	g.Plays = append(g.Plays, Play{Player: &p, Time: time.Now().Format("2006-01-02T15:04:05")})
	if g.Scores[p.Team] == winCondition {
		g.Winner = p.Team
	}
}
