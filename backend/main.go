package main

import (
	"fmt"
	"github.com/iost-official/explorer/backend/controller"
	"github.com/iost-official/explorer/backend/middleware"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/labstack/echo"
	echoMiddle "github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("~/go/src/github.com/iost-official/explorer/backend")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	db.InitConfig()
	e := echo.New()
	e.Debug = true
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler
	e.Use(middleware.CorsHeader)
	e.Use(echoMiddle.Recover())
	e.Use(echoMiddle.Logger())

	// index
	e.GET("/api/market", controller.GetMarket)
	e.GET("/api/indexBlocks", controller.GetIndexBlocks)
	e.GET("/api/indexTxns", controller.GetIndexTxns)

	// blocks
	e.GET("/api/blocks", controller.GetBlocks)
	e.GET("/api/block/:id", controller.GetBlockDetail)

	// transactions
	e.GET("/api/txs", controller.GetTxs)
	e.GET("/api/tx/:id", controller.GetTxnDetail)

	// accounts
	e.GET("/api/accounts", controller.GetAccounts)
	e.GET("/api/account/:id", controller.GetAccountDetail)
	e.GET("/api/account/:id/txs", controller.GetAccountTxs)

	// search
	e.GET("/api/search/:id", controller.GetSearch)

	// applyIOST
	e.POST("/api/sendSMS", controller.SendSMS)
	e.POST("/api/applyIOST", controller.ApplyIOST)

	//e.POST("/api/applyIOSTBenchMark", controller.ApplyIOSTBenMark)

	// mail
	e.POST("/api/feedback", controller.SendMail)

	// Test api
	e.GET("/api/test", controller.TestPage)

	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))
}
