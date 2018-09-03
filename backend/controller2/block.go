package controller2

import (
	"github.com/iost-official/explorer/backend/model2"
	"github.com/iost-official/explorer/backend/model2/db"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

const (
	BlockEachPageNum = 30
)

type BlockListOutput struct {
	BlockList []*model2.BlockOutput `json:"block_list"`
	Page      int64                 `json:"page"`
	PagePrev  int64                 `json:"page_prev"`
	PageNext  int64                 `json:"page_next"`
	PageLast  int64                 `json:"page_last"`
}

func GetIndexBlocks(c echo.Context) error {
	top10Blks, err := model2.GetBlock(1, 10)
	if err != nil {
		return err
	}

	for _, v := range top10Blks {
		v.TxList = nil
	}

	return c.JSON(http.StatusOK, top10Blks)
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

	blkList, err := model2.GetBlock(pageInt64, 30)
	if err != nil {
		return c.String(http.StatusOK, "error: "+err.Error())
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

	return c.JSON(http.StatusOK, output)
}

func GetBlockDetail(c echo.Context) error {
	blkId := c.Param("id")
	blkIdInt, err := strconv.Atoi(blkId)
	if err != nil {
		return err
	}

	blkInfo, err := db.GetBlockByHeight(int64(blkIdInt))

	return c.JSON(http.StatusOK, model2.GenerateBlockOutput(blkInfo))
}
