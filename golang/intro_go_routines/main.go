package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/aerwin3/tutorials/golang/intro_routing_api/api"
	"github.com/aerwin3/tutorials/golang/intro_routing_api/db"
	"github.com/aerwin3/tutorials/golang/intro_routing_api/processor"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	// Show buffer chan capabilities
	gameChan := make(chan *db.Game)
	p := processor.NewGameProcessor(gameChan)
	apiServer := api.NewApiServer(gameChan)

	// Start the Api and Backend game processors
	p.Start()
	apiServer.Start()

	<-sigChan
	apiServer.Stop()
	p.Stop()
}
