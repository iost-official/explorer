package main

import (
	"github.com/iost-official/explorer/backend/task/cron"
	"sync"
)

var ws = new(sync.WaitGroup)

func main()  {
	ws.Add(1)
	cron.UpdateBlocks2(ws)
	ws.Wait()
}