package controller2

import (
	"github.com/iost-official/explorer/backend/model2"
	"github.com/labstack/echo"
	"net/http"
)

func GetIndexBlocks(c echo.Context) error {
	top10Blks, err := model2.GetBlock(1, 10)
	if err != nil {
		return err
	}

	for _, v := range top10Blks {
		v.TxList = nil
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, top10Blks)
}

//func GetIndexTxns(c echo.Context) error {
//	top10Txs, err := model2.GetTransaction(1, 15, -1, "")
//	if err != nil {
//		return err
//	}
//
//	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
//	return c.JSON(http.StatusOK, top10Txs)
//}
