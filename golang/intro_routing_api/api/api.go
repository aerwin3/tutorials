package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// creates a new instance of a mux router
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/games", getGames).Methods("GET")
	r.HandleFunc("/games/{id}", getGame).Methods("GET")
	r.HandleFunc("/teams", getTeams).Methods("GET")
	r.HandleFunc("/teams/{name}", getTeam).Methods("GET")
	r.HandleFunc("/players", getPlayers).Methods("GET")
	r.HandleFunc("/players/{id}", getPlayer).Methods("GET")

	// Start serving traffic
	log.Fatal(http.ListenAndServe(":10000", r))
}
