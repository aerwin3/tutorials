package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/aerwin3/tutorials/golang/intro_routing_api/db"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getGames(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(db.GetGames())
	if err != nil {
		fmt.Printf("Error encoding get games request. %v", err.Error())
	}
}

func getGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, ok := vars["id"]; ok {
		err, g := db.GetGame(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		err = json.NewEncoder(w).Encode(g)
		if err != nil {
			fmt.Printf("Error encoding get game request. %v", err.Error())
		}
	}
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(db.GetTeams())
	if err != nil {
		fmt.Printf("Error encoding get Teams request. %v", err.Error())
	}
}

func getTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if name, ok := vars["name"]; !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else {
		err, g := db.GetTeam(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		err = json.NewEncoder(w).Encode(g)
		if err != nil {
			fmt.Printf("Error encoding get team request. %v", err.Error())
		}
	}
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(db.GetPlayers())
	if err != nil {
		fmt.Printf("Error encoding get Teams request. %v", err.Error())
	}
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, ok := vars["id"]; !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else {
		err, p := db.GetPlayerById(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		err = json.NewEncoder(w).Encode(p)
		if err != nil {
			fmt.Printf("Error encoding get player request. %v", err.Error())
		}
	}
}
