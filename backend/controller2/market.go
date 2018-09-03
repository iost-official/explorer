package controller2

import (
	"net/http"

	"github.com/iost-official/explorer/backend/model2"
	"github.com/labstack/echo"
)

func GetMarket(c echo.Context) error {
	marketInfo, err := model2.GetMarketInfo()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, marketInfo)
}
