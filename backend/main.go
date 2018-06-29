package main

import (
	"explorer/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Debug = true

	// index
	e.GET("/api/market", controller.GetMarket)
	e.GET("/api/indexBlocks", controller.GetIndexBlocks)
	e.GET("/api/indexTxns", controller.GetIndexTxns)

	// blocks
	e.GET("/api/blocks", controller.GetBlocks)
	e.GET("/api/block/:id", controller.GetBlockDetail)

	// transactions
	e.GET("/api/txs", controller.GetTxs)
	e.GET("/api/tx/:id", controller.GetTxsDetail)

	// accounts
	e.GET("/api/accounts", controller.GetAccounts)
	e.GET("/api/account/:id", controller.GetAccountDetail)
	e.GET("/api/account/:id/txs", controller.GetAccountTxs)

	// search
	e.GET("/api/search/:id", controller.GetSearch)

	// applyIOST
	e.POST("/api/sendSMS", controller.SendSMS)
	e.POST("/api/applyIOST", controller.ApplyIOST)

	// lucky bet
	e.GET("/api/luckyBetBlockInfo", controller.GetBetInfo)
	e.POST("/api/luckyBet", controller.GetLuckyBet)
	e.POST("/api/luckyBetBenchMark", controller.GetLuckyBetBenchMark)
	e.GET("/api/luckyBet/round/:id", controller.GetBetRound)
	e.GET("/api/luckyBet/addressBet/:id", controller.GetAddressBet)
	e.GET("/api/luckyBet/latestBetInfo", controller.GetLatestBetInfo)
	e.GET("/api/luckyBet/todayRanking", controller.GetTodayTop10Address)


	e.POST("/api/applyIOSTBenchMark", controller.ApplyIOSTBenMark)

	// mail
	e.POST("/api/feedback", controller.SendMail)

	e.Logger.Fatal(e.Start(":8080"))
}
