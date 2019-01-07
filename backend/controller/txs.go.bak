package controller

import (
	"github.com/iost-official/explorer/backend/model"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

const (
	TxEachPageNum = 25
	TxMaxPage     = 20
)

type TxsOutput struct {
	TxList   []*model.TxnDetail `json:"txsList"`
	Page     int64              `json:"page"`
	PagePrev int64              `json:"pagePrev"`
	PageNext int64              `json:"pageNext"`
	PageLast int64              `json:"pageLast"`
	TotalLen int                `json:"totalLen"`
}

func GetTxnDetail(c echo.Context) error {
	txHash := c.Param("id")

	if txHash == "" {
		return nil
	}

	txnOut, err := model.GetDetailTxn(txHash)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, FormatResponse(txnOut))
}

func GetIndexTxns(c echo.Context) error {
	topTxs, err := model.GetFlatTxnSlicePage(1, 15, -1, "")

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, FormatResponse(topTxs))
}

func GetTxs(c echo.Context) error {
	page := c.QueryParam("page")
	address := c.QueryParam("account")
	blk := c.QueryParam("block")

	pageInt64, err := strconv.ParseInt(page, 10, 64)

	if err != nil || pageInt64 <= 0 {
		pageInt64 = 1
	}

	blockInt64, err := strconv.ParseInt(blk, 10, 64)

	if err != nil {
		blockInt64 = -1
	}

	txList, err := model.GetFlatTxnSlicePage(pageInt64, TxEachPageNum, blockInt64, address)

	if err != nil {
		return err
	}

	var (
		lastPage int64
		totalLen int
	)

	if address != "" {
		// get total page count for specific account
		lastPage, _ = db.GetFlatTxPageCntWithAddress(TxEachPageNum, address)
		totalLen, _ = db.GetTotalFlatTxnLen(address, -1)
	} else if blk != "" {
		// get total page count for specific block
		lastPage, _ = db.GetFlatTxPageCntWithBlk(TxEachPageNum, blockInt64)
		totalLen, _ = db.GetTotalFlatTxnLen("", blockInt64)
	} else {
		// get all page count for all transactions
		lastPage, _ = db.GetFlatTxTotalPageCnt(TxEachPageNum, "", -1)
		totalLen, _ = db.GetTotalFlatTxnLen("", -1)
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

	return c.JSON(http.StatusOK, FormatResponse(output))
}
