package controller

import (
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"explorer/util/session"

	"github.com/labstack/echo"
	"encoding/json"
	"explorer/model/db"
)

const (
	VerifyCodeLen     = 6
	MobileMaxSendTime = 1

	AccountSid        = "AC47b8c0b922a3eb016f263869ac0d2951"
	AuthToken         = "0daee011527a806c76792d46cd71dd13"
	TwilioSmsUrl      = "https://api.twilio.com/2010-04-01/Accounts/" + AccountSid + "/Messages.json"
)

type CommOutput struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

var (
	httpClient     *http.Client
	verifySeed     = rand.NewSource(time.Now().UnixNano())
	verifyCodeList = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	ErrEmptySSID = errors.New("empty session id")
	ErrGreCaptcha = errors.New("reCAPTCHA check failed")
	ErrMobileApplyExceed = errors.New("mobile applied earlier today")
)

func init() {
	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 10,
		},
	}
}

func SendSMS(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	mobile := c.FormValue("mobile")
	gcaptcha := c.FormValue("gcaptcha")

	remoteip := c.Request().Header.Get("Iost_Remote_Addr")
	if !verifyGCAP(gcaptcha, remoteip) {
		log.Println(ErrGreCaptcha.Error())
		return c.JSON(http.StatusOK, &CommOutput{1, ErrGreCaptcha.Error()})
	}

	if len(mobile) < 10 || mobile[0] != '+' {
		return c.JSON(http.StatusOK, &CommOutput{2, ErrInvalidInput.Error()})
	}

	mobileSendNum, err := db.GetApplyNumTodayByMobile(mobile)
	if err != nil {
		return c.JSON(http.StatusOK, &CommOutput{3, err.Error()})
	}

	if mobileSendNum >= MobileMaxSendTime {
		return c.JSON(http.StatusOK, &CommOutput{4, ErrMobileApplyExceed.Error()})
	}

	sess, _ := session.GlobalSessions.SessionStart(c.Response(), c.Request())
	defer sess.SessionRelease(c.Response())

	vc, err := sendSMS(mobile)
	if err != nil {
		return ErrEmptySSID
	}

	sess.Set("verification", vc)
	log.Printf("sendSMS ssid: %s, vc: %s\n", sess.SessionID(), vc)

	mobileSendNum++
	

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, &CommOutput{0, "ok"})
}

func sendSMS(number string) (string, error) {
	vc := generateVerifyCode(VerifyCodeLen)

	postData := url.Values{}
	postData.Set("To", number)
	postData.Set("From", "+13192642988")
	postData.Set("Body", "[IOST] verification code: "+vc)

	req, _ := http.NewRequest("POST", TwilioSmsUrl, strings.NewReader(postData.Encode()))
	req.SetBasicAuth(AccountSid, AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("sendSMS error:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("sendSMS error:", err)
		return "", err
	}

	log.Println("sendSMS body:", string(body))

	return vc, nil
}

func generateVerifyCode(codeLen int) string {
	r := rand.New(verifySeed)
	rList := r.Perm(len(verifyCodeList))

	var vc []byte
	for i := 0; i < codeLen; i++ {
		vc = append(vc, verifyCodeList[rList[i]])
	}

	return string(vc)
}

func verifyGCAP(gcap, remoteip string) bool {
	postData := url.Values{}
	postData.Set("secret", GCAPSecretKey)
	postData.Set("response", gcap)
	postData.Set("remoteip", remoteip)

	req, _ := http.NewRequest("POST", GCAPVerifyUrl, strings.NewReader(postData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := gcapHttpClient.Do(req)
	if err != nil {
		log.Println("verifyGCAP error:", err)
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("verifyGCAP error:", err)
		return false
	}

	log.Println("verifyGCAP result:", string(body))
	var result *GCAPResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("verifyGCAP json unmaral error:", err)
		return false
	}

	return result.Success
}
