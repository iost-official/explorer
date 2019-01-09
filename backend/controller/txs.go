package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/iost-official/explorer/backend/model"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/labstack/echo"
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
	topTxs, err := model.GetFlatTxnSlicePage(1, 15, -1)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, FormatResponse(topTxs))
}

func getTxsOfAddress(c echo.Context, address string, page int64) error {
	start := (int(page) - 1) * AccountEachPage
	var txHashes []string
	var totalLen int
	if isContract(address) {
		conTxs, err := db.GetContractTxByID(address, start, AccountEachPage)
		if err != nil {
			return err
		}
		for _, conTx := range conTxs {
			txHashes = append(txHashes, conTx.TxHash)
		}

		totalLen, err = db.GetContractTxNumber(address)
		if err != nil {
			log.Printf("get contract tx number failed. contract=%v, err=%v", address, err)
		}
	} else {
		accTxs, err := db.GetAccountTxByName(address, start, AccountEachPage)
		if err != nil {
			return err
		}
		for _, accTx := range accTxs {
			txHashes = append(txHashes, accTx.TxHash)
		}

		totalLen, err = db.GetAccountTxNumber(address)
		if err != nil {
			log.Printf("get account tx number failed. account=%v, err=%v", address, err)
		}

	}

	txList, err := db.GetTxsByHash(txHashes)
	if err != nil {
		return err
	}

	lastPage := calLastPage(totalLen)
	if lastPage > TxMaxPage {
		lastPage = TxMaxPage
	}

	output := &TxsOutput{
		TxList:   model.ConvertTxsOutputs(txList),
		Page:     page,
		PagePrev: page - 1,
		PageNext: page + 1,
		PageLast: int64(lastPage),
		TotalLen: totalLen,
	}

	return c.JSON(http.StatusOK, FormatResponse(output))

}

func GetTxs(c echo.Context) error {
	page := c.QueryParam("page")
	address := c.QueryParam("account")
	blk := c.QueryParam("block")

	pageInt64, err := strconv.ParseInt(page, 10, 64)

	if err != nil || pageInt64 <= 0 {
		pageInt64 = 1
	}

	if address != "" {
		return getTxsOfAddress(c, address, pageInt64)
	}

	blockInt64, err := strconv.ParseInt(blk, 10, 64)

	if err != nil {
		blockInt64 = -1
	}

	txList, err := model.GetFlatTxnSlicePage(pageInt64, TxEachPageNum, blockInt64)

	if err != nil {
		return err
	}

	lastPage, err := db.GetTxTotalPageCnt(TxEachPageNum, blockInt64)
	if err != nil {
		log.Printf("GetTxTotalPageCnt failed. blockInt64=%v, err=%v", blockInt64, err)
	}
	txCount, err := db.GetTxCountByNumber(blockInt64)
	if err != nil {
		log.Printf("GetTxCountByNumber failed. blockInt64=%v, err=%v", blockInt64, err)
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
		TotalLen: txCount,
	}

	return c.JSON(http.StatusOK, FormatResponse(output))
}
