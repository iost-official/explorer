package main

import (
	"github.com/iost-official/explorer/backend/task/cron"
	"sync"
)

var ws = new(sync.WaitGroup)

func main()  {
	ws.Add(4)
	go cron.UpdateBlocks2(ws)
	go cron.ProcessFailedSyncBlocks(ws)
	go cron.UpdateTxns(ws)
	go cron.UpdateBlockPay(ws)
	ws.Wait()
}
