package db

import (
	"errors"
	"math/rand"
)

type Team struct {
	Players map[int]*Player `json:"players"`
	Name    string          `json:"name"`
}

func (t *Team) AddPlayer(name string) {
	assigned := false
	j := rand.Intn(99)

	for !assigned {
		if _, ok := t.Players[j]; !ok {
			t.Players[j] = &Player{Name: name, Jersey: j, Team: t.Name}
			assigned = true
		}
	}
}

func GetTeams() (ts []Team) {
	for _, v := range teams {
		ts = append(ts, *v)
	}
	return
}

func GetTeam(name string) (err error, ts Team) {
	if t, ok := teams[name]; !ok {
		err = errors.New("not found")
		return
	} else {
		ts = *t
	}
	return
}
