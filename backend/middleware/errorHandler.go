package middleware

import (
	"github.com/iost-official/explorer/backend/controller2"
	"github.com/labstack/echo"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	c.JSON(http.StatusOK, controller2.ErrorResponse{
		Code: 1,
		Message: err.Error(),
	})
}
