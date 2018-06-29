package controller

import (
	"strconv"

	"explorer/model"
	"explorer/model/blockchain"
	"explorer/model/cache"
	"explorer/model/db"
	"fmt"
	"github.com/iost-official/prototype/common"
	"github.com/iost-official/prototype/core/tx"
	"github.com/labstack/echo"
	"net/http"
)

const (
	TxEachPageNum = 25
	TxMaxPage     = 20
)

type TxsOutput struct {
	TxList   []*model.TransactionOutput `json:"txs_list"`
	Page     int64                      `json:"page"`
	PagePrev int64                      `json:"page_prev"`
	PageNext int64                      `json:"page_next"`
	PageLast int64                      `json:"page_last"`
	TotalLen int                        `json:"total_len"`
}

type BetTxsOutput struct {
	From string `json:"from"`
	Time int64  `json:"time"`
	Code string `json:"code"`
}

// e.GET("/txs", getBlock)
func GetTxs(c echo.Context) error {
	page := c.QueryParam("p")
	address := c.QueryParam("a")
	blk := c.QueryParam("b")

	if address == blockchain.BetHash {
		address = "Bet"
	}

	pageInt64, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		pageInt64 = 1
	}

	if pageInt64 <= 0 {
		pageInt64 = 1
	}

	blockInt64, err := strconv.ParseInt(blk, 10, 64)
	if err != nil {
		blockInt64 = -1
	}

	txList, err := model.GetTransaction(pageInt64, TxEachPageNum, blockInt64, address)
	if err != nil {
		return c.String(http.StatusOK, "error: "+err.Error())
	}

	var (
		lastPage int64
		totalLen int
	)
	if address != "" {
		lastPage, _ = db.GetTxDetailLastPageWithAddress(TxEachPageNum, address)
		totalLen, _ = db.GetTotalTxnDetailLen(address, -1)
	} else if blk != "" {
		lastPage, _ = db.GetTxDetailLastPageWithBlk(TxEachPageNum, blockInt64)
		totalLen, _ = db.GetTotalTxnDetailLen("", blockInt64)
	} else {
		lastPage, _ = db.GetTxDetailLastPage(TxEachPageNum)
		totalLen, _ = db.GetTotalTxnDetailLen("", -1)
	}
	if lastPage > TxMaxPage {
		lastPage = TxMaxPage
	}

	output := &TxsOutput{
		TxList:   txList,
		Page:     pageInt64,
		PagePrev: pageInt64 - 1,
		PageNext: pageInt64 + 1,
		PageLast: lastPage,
		TotalLen: totalLen,
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, output)
}

func GetTxsDetail(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	txHash := c.Param("id")
	if txHash == "" {
		return nil
	}

	if txHash == blockchain.BetHash {
		txnInterface, _ := cache.GlobalCache.Get("betMainCode")
		fmt.Println(txnInterface)
		txn, _ := txnInterface.(*tx.Tx)
		if txn == nil {
			return nil
		}

		return c.JSON(http.StatusOK, &BetTxsOutput{
			Time: txn.Time,
			From: common.Base58Encode(txn.Publisher.Pubkey),
			Code: txn.Contract.Code(),
		})
	}

	txn, err := db.GetTxnDetailByHash(txHash)
	if err != nil {
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	txnOutput := model.GenerateTxnOutput(txn)

	return c.JSON(http.StatusOK, txnOutput)
}
