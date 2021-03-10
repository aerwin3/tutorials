package db

import "errors"

type Player struct {
	Id     string `json:"id"`
	Jersey int    `json:"jersey"`
	Name   string `json:"name"`
	Team   string `json:"team"`
}

func GetPlayers() (ts []Player) {
	for _, v := range players {
		ts = append(ts, v)
	}
	return
}

func GetPlayerById(id string) (err error, p Player) {
	for _, v := range players {
		if v.Id == id {
			p = v
			return
		}
	}
	err = errors.New("not found")
	return
}
