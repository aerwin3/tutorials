package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"

	"github.com/aerwin3/tutorials/golang/intro_routing_api/db"
)

func (as *apiServer) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to GOPHER BALL!!!")
}

func (as *apiServer) getGames(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(db.GetGames())
	if err != nil {
		fmt.Printf("Error encoding get games request. %v", err.Error())
	}
}

func (as *apiServer) getGame(w http.ResponseWriter, r *http.Request) {
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
		return
	}
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func (as *apiServer) startGame(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var req struct {
		Teams []string `json:"teams"`
	}

	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if len(req.Teams) != 2 {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	game := db.CreateGame(req.Teams)
	as.gameChan <- game
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"id": game.Id})
}

func (as *apiServer) getTeams(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(db.GetTeams())
	if err != nil {
		fmt.Printf("Error encoding get Teams request. %v", err.Error())
	}
}

func (as *apiServer) getTeam(w http.ResponseWriter, r *http.Request) {
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

func (as *apiServer) getPlayers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(db.GetPlayers())
	if err != nil {
		fmt.Printf("Error encoding get Teams request. %v", err.Error())
	}
}

func (as *apiServer) getPlayer(w http.ResponseWriter, r *http.Request) {
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
