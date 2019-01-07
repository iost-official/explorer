package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/iost-official/explorer/backend/model/db"
	"github.com/iost-official/explorer/backend/util/session"
	"github.com/labstack/echo"
)

const (
	VerifyCodeLen     = 6
	MobileMaxSendTime = 1

	AccountSid   = "ACbb9c8973309348ffca81bb71291b3a4c"
	AuthToken    = "1e224c4c3166d94b0578a967ff1dd0ac"
	TwilioSmsUrl = "https://api.twilio.com/2010-04-01/Accounts/" + AccountSid + "/Messages.json"
)

var (
	httpClient     *http.Client
	verifySeed     = rand.NewSource(time.Now().UnixNano())
	verifyCodeList = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	ErrEmptySSID         = errors.New("empty session id")
	ErrGreCaptcha        = errors.New("reCAPTCHA check failed")
	ErrMobileApplyExceed = errors.New("mobile applied earlier today")
)

type GCAPResponse struct {
	Success     bool   `json:"success"`
	ChallengeTs string `json:"challengeTs"`
	Hostname    string `json:"hostname"`
}

func init() {
	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 10,
		},
	}
}

func SendSMS(c echo.Context) error {
	mobile := c.FormValue("mobile")
	gcaptcha := c.FormValue("gcaptcha")

	remoteip := c.Request().Header.Get("Iost_Remote_Addr")
	if !verifyGCAP(gcaptcha, remoteip) {
		return ErrGreCaptcha
	}

	if len(mobile) < 10 || mobile[0] != '+' {
		return ErrInvalidInput
	}

	mobileSendNum, err := db.GetApplyNumTodayByMobile(mobile)
	if err != nil {
		return err
	}

	if mobileSendNum >= MobileMaxSendTime {
		return ErrMobileApplyExceed
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

	return c.JSON(http.StatusOK, FormatResponse(&CommOutput{0, "ok"}))
}

func sendSMS(number string) (string, error) {
	vc := generateVerifyCode(VerifyCodeLen)

	postData := url.Values{}
	postData.Set("To", number)
	postData.Set("From", "+12568183697")
	postData.Set("Body", "[IOST] verification code: "+vc)

	req, _ := http.NewRequest("POST", TwilioSmsUrl, strings.NewReader(postData.Encode()))
	req.SetBasicAuth(AccountSid, AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
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
