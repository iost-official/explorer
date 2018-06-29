package controller

import (
	"explorer/model/db"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type SearchOutput struct {
	Search string `json:"data"`
	Type   string `json:"type"`
	Text   string `json:"text,omitempty"`
}

func GetSearch(c echo.Context) error {
	search := c.Param("id")
	if search == "" {
		return nil
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

	blkHash, _ := db.GetBlockByHash(search)
	if blkHash != nil {
		output.Type = "block"
		output.Text = strconv.FormatInt(blkHash.Head.Number, 10)
	}

	if searchInt64, _ := strconv.ParseInt(search, 10, 64); searchInt64 > 0 {
		block, _ := db.GetBlockByHeight(searchInt64)
		if block != nil {
			output.Type = "block"
		}
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, output)
}
