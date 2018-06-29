package main

import (
	"sync"
	"explorer/task/cron"
)

var wg = new(sync.WaitGroup)

func main()  {
	wg.Add(5)
	go cron.UpdateBlocks(wg)
	go cron.UpdateTxs(wg)
	go cron.UpdateTxnDetail(wg)
	go cron.UpdateBlockPay(wg)
	go cron.UpdateAccounts(wg)
	go cron.UpdateBet(wg)
	wg.Wait()
}
