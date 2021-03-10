package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aerwin3/tutorials/golang/intro_routing_api/db"
	"github.com/gorilla/mux"
)

type Server interface {
	Start()
	Stop()
}

type apiServer struct {
	server   *http.Server
	gameChan chan *db.Game
}

func NewApiServer(gameChan chan *db.Game) Server {
	return &apiServer{gameChan: gameChan}
}

func (as *apiServer) Start() {
	// creates a new instance of a mux router
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", as.indexHandler).Methods("GET")
	r.HandleFunc("/games", as.startGame).Methods("POST")
	r.HandleFunc("/games", as.getGames).Methods("GET")
	r.HandleFunc("/games/{id}", as.getGame).Methods("GET")
	r.HandleFunc("/teams", as.getTeams).Methods("GET")
	r.HandleFunc("/teams/{name}", as.getTeam).Methods("GET")
	r.HandleFunc("/players", as.getPlayers).Methods("GET")
	r.HandleFunc("/players/{id}", as.getPlayer).Methods("GET")
	as.server = &http.Server{
		Addr:    ":10000",
		Handler: r,
	}
	go func() {
		if err := as.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func (as *apiServer) Stop() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := as.server.Shutdown(ctx)
	if err != nil {
		fmt.Printf("Failed to Shutdown API")
	}
	log.Print("Api server shutdown.")
}
