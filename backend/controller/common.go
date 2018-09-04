package controller

import (
	"errors"
	"net/http"
	"regexp"
)

type CommOutput struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}


var (
	ErrInvalidInput     = errors.New("invalid input")
	ErrMobileVerfiy     = errors.New("mobile verify failed")
	ErrOutOfRetryTime   = errors.New("out of retry time")
	ErrOutOfCheckTxHash = errors.New("out of check txHash retry time")
	RegEmail            = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	gcapHttpClient      *http.Client
)

const (
	AccountEachPage = 25
	AccountMaxPage  = 20
	GCAPVerifyUrl   = "https://www.google.com/recaptcha/api/siteverify"
	GCAPSecretKey   = "6Lc1vF8UAAAAAGv1XihAK4XygBMn3UobipWMqBym"
)


type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func FormatResponse(data interface{}) Response {
	return Response{
		Code: 0,
		Data: data,
	}
}
