package main

import (
	"github.com/iost-official/explorer/backend/config"
	"github.com/iost-official/explorer/backend/task/cron"
	"sync"
)

var ws = new(sync.WaitGroup)

func main()  {
	config.ReadConfig()

	// start tasks
	ws.Add(7)
	go cron.UpdateBlocks(ws)
	go cron.ProcessFailedSyncBlocks(ws)
	go cron.UpdateTxns(ws, 0)
	go cron.UpdateTxns(ws, 1)
	go cron.UpdateRpcErrTxns(ws)
	go cron.UpdateBlockPay(ws)
	go cron.UpdateAccounts(ws)
	ws.Wait()
}
