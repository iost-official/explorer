package controller

import (
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type SearchOutput struct {
	Search string `json:"search"`
	Type   string `json:"type"`
	Text   string `json:"text,omitempty"`
}

func GetSearch(c echo.Context) error {
	search := c.Param("id")
	if search == "" {
		return errors.New("Nothing to search")
	}

	output := &SearchOutput{
		Search: search,
	}

	account, _ := db.GetAccountByAddress(search)
	if account != nil {
		output.Type = "account"
	}

	tx, _ := db.GetTxnDetailByHash(search)
	if tx != nil {
		output.Type = "tx"
	}

	blkHash, _, _ := db.GetBlockByHash(search)
	if blkHash != nil {
		output.Type = "block"
		output.Text = strconv.FormatInt(blkHash.BlockNumber, 10)
	}

	if searchInt64, _ := strconv.ParseInt(search, 10, 64); searchInt64 > 0 {
		block, _ := db.GetBlockByHeight(searchInt64)
		if block != nil {
			output.Type = "block"
		}
	}
	return c.JSON(http.StatusOK, FormatResponse(output))
}
