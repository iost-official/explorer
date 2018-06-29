package controller

import (
	"net/http"

	"explorer/model"

	"github.com/labstack/echo"
)

func GetMarket(c echo.Context) error {
	marketInfo, err := model.GetMarketInfo()
	if err != nil {
		return err
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, marketInfo)
}
