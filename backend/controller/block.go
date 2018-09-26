package controller

import (
	"github.com/iost-official/explorer/backend/model"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

const (
	BlockEachPageNum = 30
)

type BlockListOutput struct {
	BlockList []*model.BlockOutput `json:"blockList"`
	Page      int64                `json:"page"`
	PagePrev  int64                `json:"pagePrev"`
	PageNext  int64                `json:"pageNext"`
	PageLast  int64                `json:"pageLast"`
}

func GetIndexBlocks(c echo.Context) error {
	top10Blks, err := model.GetBlock(1, 10)
	if err != nil {
		return err
	}

	for _, v := range top10Blks {
		v.TxList = nil
	}

	return c.JSON(http.StatusOK, FormatResponse(top10Blks))
}

func GetBlocks(c echo.Context) error {
	page := c.QueryParam("p")

	pageInt64, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		pageInt64 = 1
	}

	if pageInt64 <= 0 {
		pageInt64 = 1
	}

	blkList, err := model.GetBlock(pageInt64, 30)
	if err != nil {
		return err
	}

	for _, v := range blkList {
		v.TxList = nil
	}

	output := &BlockListOutput{
		blkList,
		pageInt64,
		pageInt64 - 1,
		pageInt64 + 1,
		db.GetBlockLastPage(BlockEachPageNum),
	}

	return c.JSON(http.StatusOK, FormatResponse(output))
}

func GetBlockDetail(c echo.Context) error {
	blkId := c.Param("id")
	blkIdInt, err := strconv.Atoi(blkId)
	if err != nil {
		return err
	}

	blkInfo, err := db.GetBlockByHeight(int64(blkIdInt))

	if nil != err {
		return err
	}

	blkOutput := model.GenerateBlockOutput(blkInfo)

	return c.JSON(http.StatusOK, FormatResponse(blkOutput))
}
