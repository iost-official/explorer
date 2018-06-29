package controller

import (
	"net/http"

	"explorer/model"

	"github.com/labstack/echo"
)

func GetIndexBlocks(c echo.Context) error {
	top10Blks, err := model.GetBlock(1, 10)
	if err != nil {
		return err
	}

	for _, v := range top10Blks {
		v.TxList = nil
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, top10Blks)
}

func GetIndexTxns(c echo.Context) error {
	top10Txs, err := model.GetTransaction(1, 15, -1, "")
	if err != nil {
		return err
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, top10Txs)
}
