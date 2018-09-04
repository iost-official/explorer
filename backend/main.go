package main

import (
	"github.com/iost-official/explorer/backend/controller2"
	"github.com/iost-official/explorer/backend/middleware"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler
	e.Use(middleware.CorsHeader)

	// index
	e.GET("/api/market", controller2.GetMarket)
	e.GET("/api/indexBlocks", controller2.GetIndexBlocks)
	//e.GET("/api/indexTxns", controller.GetIndexTxns)

	// blocks
	e.GET("/api/blocks", controller2.GetBlocks)
	e.GET("/api/block/:id", controller2.GetBlockDetail)

	// transactions
	//e.GET("/api/txs", controller.GetTxs)
	e.GET("/api/tx/:id", controller2.GetTxsDetail)

	// accounts
	e.GET("/api/accounts", controller2.GetAccounts)
	e.GET("/api/account/:id", controller2.GetAccountDetail)
	e.GET("/api/account/:id/txs", controller2.GetAccountTxs)

	// search
	//e.GET("/api/search/:id", controller.GetSearch)

	// applyIOST
	e.POST("/api/sendSMS", controller2.SendSMS)
	//e.POST("/api/applyIOST", controller2.ApplyIOST)

	//e.POST("/api/applyIOSTBenchMark", controller.ApplyIOSTBenMark)

	// mail
	e.POST("/api/feedback", controller2.SendMail)

	// Test api
	e.GET("/api/test", controller2.TestPage)

	e.Logger.Fatal(e.Start(":8080"))
}
