package controller

import (
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/labstack/echo"
	"net/http"
)

func DropDatabase(c echo.Context) error {

	db, err := db.GetDb()
	if err != nil {
		return err
	}

	err = db.DropDatabase()
	if err != nil {
		return err
	}

	// TODO: 清数据库前需要先停掉同步任务
	return c.JSON(http.StatusOK, FormatResponse("Success"))
}