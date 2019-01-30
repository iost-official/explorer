package main

import (
	"context"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/iost-official/explorer/backend/config"
	"github.com/iost-official/explorer/backend/controller"
	"github.com/iost-official/explorer/backend/middleware"
	"github.com/iost-official/explorer/backend/util/boot"
	"github.com/labstack/echo"
	echoMiddle "github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {

	//
	// boot service
	//

	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = v.Unmarshal(&boot.C)
	if err != nil {
		panic(err)
	}

	// hot reload
	ctx, _ := context.WithCancel(context.Background())
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config change detected, updating config ...")

		err := v.ReadInConfig()
		if err != nil {
			panic(err)
		}

		err = v.Unmarshal(&boot.C)
		if err != nil {
			panic(err)
		}
	})

	go func() {
		v.WatchConfig()
		<-ctx.Done()
	}()


	//
	// original explorer
	//

	config.ReadConfig("")
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

	/* // blocks */
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

	// mail
	e.POST("/api/feedback", controller.SendMail)

	// applyIOST
	e.POST("/api/applyIOST", controller.ApplyIOST)

	e.Logger.Fatal(e.Start(":8080"))
}
