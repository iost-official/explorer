package main

import (
	"sync"

	"github.com/GincoInc/iost-explorer/backend/config"
	"github.com/GincoInc/iost-explorer/backend/task/cron"
)

var ws = new(sync.WaitGroup)

func main() {
	config.ReadConfig("")

	// start tasks
	ws.Add(1)
	// download block
	go cron.UpdateBlocks(ws)
	ws.Wait()
}
