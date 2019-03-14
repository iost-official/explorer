package controller

import (
	"github.com/iost-official/explorer/backend/util/mail"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func SendMail(c echo.Context) error {
	to := c.FormValue("email")
	content := c.FormValue("content")

	err := mail.SendMail(to, content)

	if err != nil {
		log.Println("SendMail error:", err)
		return err
	}

	return c.JSON(http.StatusOK, FormatResponse("Mail send success"))
}
