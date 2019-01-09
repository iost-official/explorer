package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"fmt"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/labstack/echo"
	"regexp"
)

var (
	applyHttpClient *http.Client
	ipFreqMapMutex  sync.RWMutex
	applyErrReg     = regexp.MustCompile(`Uncaught exception:(.*)\n`)
	ipFreqMap       = make(map[string]int)
	ErrGreCaptcha   = errors.New("reCAPTCHA check failed")
	ErrApplyFreq    = errors.New("client ip create account exceeds limit")
)

func init() {
	applyHttpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 10,
		},
	}

	go cronUpdateIPFreqMap()
}

func ApplyIOST(c echo.Context) error {
	accountPubKey := c.FormValue("address")
	accountName := c.FormValue("account")
	email := c.FormValue("email")
	gcaptcha := c.FormValue("gcaptcha")

	if accountPubKey == "" || email == "" {
		log.Println("ApplyIOST nil params")
		return ErrInvalidInput
	}

	if len(accountPubKey) < 35 {
		log.Println("ApplyIOST invalid accountPubKey")
		return ErrInvalidInput
	}

	if !RegEmail.MatchString(email) {
		log.Println("ApplyIOST invaild regexp email")
		return ErrInvalidInput
	}

	remoteIP := c.Request().Header.Get("Iost_Remote_Addr")
	if !verifyGCAP(gcaptcha, remoteIP) {
		log.Println(ErrGreCaptcha.Error())
		return c.JSON(http.StatusOK, &CommOutput{1, ErrGreCaptcha.Error()})
	}

	if !checkIPFreq(remoteIP) {
		log.Println(ErrApplyFreq.Error())
		return c.JSON(http.StatusOK, &CommOutput{2, ErrApplyFreq.Error()})
	}

	if err := createAccount(accountPubKey, accountName); err != nil {
		log.Println("createAccount err: ", err.Error())
		return c.JSON(http.StatusOK, &CommOutput{2, err.Error()})
	}

	db.AddApply(accountPubKey, accountName, email, remoteIP)

	return c.JSON(http.StatusOK, FormatResponse("create success"))
}

// 简单的check逻辑，不落盘数据库，也就是说如果重启，单ip可以再申请ApplyAccountPerIPPerDay次。
func checkIPFreq(remoteIP string) bool {
	ipFreqMapMutex.RLock()
	if ipFreq := ipFreqMap[remoteIP]; ipFreq >= ApplyAccountPerIPPerDay {
		ipFreqMapMutex.RUnlock()
		return false
	}
	ipFreqMapMutex.RUnlock()

	ipFreqMapMutex.Lock()
	defer ipFreqMapMutex.Unlock()

	if ipFreq := ipFreqMap[remoteIP]; ipFreq >= ApplyAccountPerIPPerDay {
		return false
	}
	ipFreqMap[remoteIP]++

	return true
}

func cronUpdateIPFreqMap() {
	for {
		// 等到下一个零点
		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C

		log.Println("refresh IPFreqMap now...")
		ipFreqMapMutex.Lock()
		ipFreqMap = make(map[string]int)
		ipFreqMapMutex.Unlock()
	}
}

func createAccount(pubKey, name string) error {
	reqUrl := fmt.Sprintf(
		"http://%s/register?ID=%s&Name=%s&Net=%s&mk=%s",
		ApplyHost,
		pubKey,
		name,
		ApplyNet,
		ApplyMK,
	)
	req, err := http.NewRequest("IOST_BOOT", reqUrl, nil)
	if err != nil {
		return err
	}

	resp, err := applyHttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	bodyStr := string(body)
	if !strings.Contains(bodyStr, "Uncaught exception") {
		return nil
	}

	errFound := applyErrReg.FindStringSubmatch(bodyStr)
	if len(errFound) < 2 {
		return errors.New(bodyStr)
	}

	return errors.New(errFound[1])
}

type GCAPResponse struct {
	Success     bool   `json:"success"`
	ChallengeTs string `json:"challengeTs"`
	Hostname    string `json:"hostname"`
}

func verifyGCAP(gcap, remoteIP string) bool {
	postData := url.Values{}
	postData.Set("secret", GCAPSecretKey)
	postData.Set("response", gcap)
	postData.Set("remoteip", remoteIP)

	req, _ := http.NewRequest("POST", GCAPVerifyUrl, strings.NewReader(postData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := applyHttpClient.Do(req)
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
