package processor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aerwin3/tutorials/golang/intro_routing_api/db"
)

const (
	workerCount = 5
)

type Processor interface {
	Start()
	Stop()
}

type gameProcessor struct {
	gameChan  chan *db.Game
	wg        *sync.WaitGroup
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func NewGameProcessor(gameChan chan *db.Game) Processor {
	ctx, cancel := context.WithCancel(context.Background())
	return &gameProcessor{
		gameChan:  gameChan,
		wg:        &sync.WaitGroup{},
		ctx:       ctx,
		ctxCancel: cancel,
	}
}

func (gp *gameProcessor) Start() {
	for i := 0; i < workerCount; i++ {
		gp.wg.Add(1)
		go func() {
			defer gp.wg.Done()
			for {
				select {
				case game, ok := <-gp.gameChan:
					if !ok {
						return
					}
					runGame(gp.ctx, game)
				}
			}
		}()
	}
	return
}

func (gp *gameProcessor) Stop() {
	fmt.Printf("Stopping game processor.")
	close(gp.gameChan)
	gp.ctxCancel()
	gp.wg.Wait()
}

func runGame(ctx context.Context, game *db.Game) {
	for game.Winner == "" {
		select {
		case <-time.After(2 * time.Second):
			game.RunPlay()
		case <-ctx.Done():
			return
		}
	}
}
