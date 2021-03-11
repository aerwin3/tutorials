package db

var (
	games   []*Game
	teams   map[string]*Team
	players []Player
)

func init() {
	players = []Player{
		{Id: "1", Name: "Bob", Team: "Zoomers", Jersey: 1},
		{Id: "2", Name: "Jill", Team: "Zoomers", Jersey: 2},
		{Id: "3", Name: "Chipper", Team: "Flyers", Jersey: 20},
		{Id: "4", Name: "Hoppey", Team: "Flyers", Jersey: 21},
	}
	teams = map[string]*Team{
		"Zoomers": {
			Name:    "Zoomers",
			Players: map[int]*Player{players[0].Jersey: &players[0], players[1].Jersey: &players[1]},
		},
		"Flyers": {
			Name:    "Flyers",
			Players: map[int]*Player{players[2].Jersey: &players[2], players[3].Jersey: &players[3]},
		},
	}
	games = []*Game{}
}
