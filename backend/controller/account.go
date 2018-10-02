package controller

import (
	"errors"
	"fmt"
	"github.com/iost-official/explorer/backend/model"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/blkchain"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/iost-official/explorer/backend/util/session"
	"github.com/iost-official/go-iost/common"
	"github.com/labstack/echo"
)

type AccountsOutput struct {
	AccountList []*db.Account `json:"accountList"`
	Page        int           `json:"page"`
	PagePrev    int           `json:"pagePrev"`
	PageNext    int           `json:"pageNext"`
	PageLast    int           `json:"pageLast"`
	TotalLen    int           `json:"totalLen"`
}

type AccountTxsOutput struct {
	Address  string           `json:"address"`
	TxnList  []*db.JsonFlatTx `json:"txnList"`
	TxnLen   int              `json:"txnLen"`
	PageLast int              `json:"pageLast"`
}

func init() {
	gcapHttpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 10,
		},
	}
}

func calLastPage(total int) int {
	var lastPage int
	if total%AccountEachPage == 0 {
		lastPage = total / AccountEachPage
	} else {
		lastPage = total/AccountEachPage + 1
	}

	if lastPage > AccountMaxPage { // ?
		lastPage = AccountMaxPage
	}
	return lastPage
}

func GetAccounts(c echo.Context) error {
	page := c.QueryParam("p")

	pageInt, _ := strconv.Atoi(page)
	if pageInt <= 0 {
		pageInt = 1
	}

	start := (pageInt - 1) * AccountEachPage
	accountList, err := db.GetAccounts(start, AccountEachPage)
	if err != nil {
		return err
	}

	accountTotalLen, err := db.GetAccountsTotalLen()
	if err != nil {
		return err
	}
	lastPage := calLastPage(accountTotalLen)

	output := &AccountsOutput{
		AccountList: accountList,
		Page:        pageInt,
		PagePrev:    pageInt - 1,
		PageNext:    pageInt + 1,
		PageLast:    lastPage,
		TotalLen:    accountTotalLen,
	}

	return c.JSON(http.StatusOK, FormatResponse(output))
}

func GetAccountDetail(c echo.Context) error {
	// TODO 检查地址格式
	address := c.Param("id")
	if address == "" {
		return errors.New("Address required")
	}

	col, err := db.GetCollection(db.CollectionAccount)
	if err != nil {
		return err
	}

	if !(address[0:4] != "IOST" || address[0:8] != "Contract") {
		return errors.New("Invalid address")
	}

	account, err := db.GetAccountByAddress(address)
	// 如果记录不存在创建记录
	if err != nil && err.Error() == "not found" {
		account = &db.Account{
			address,
			0,
			0,
			0,
		}
		err = col.Insert(*account)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	toUpdate := bson.M{}
	txCount, err := db.GetAccountTxCount(address)
	if err == nil {
		account.TxCount = txCount
		toUpdate["tx_count"] = txCount
	}
	if address[0:4] == "IOST" { // IOST 地址获取余额
		balance, err := blkchain.GetBalance(address)
		if err == nil {
			account.Balance = float64(balance)
			toUpdate["balance"] = balance
		}
	}
	err = col.Update(bson.M{"address": address}, bson.M{"$set": toUpdate})

	if err != nil {
		return err
	}

	// 计算iost对应的美元价值
	marketInfo, err := model.GetMarketInfo()
	price, _ := strconv.ParseFloat(marketInfo.Price, 32)
	value := account.Balance / 100000000 * price

	// 合约地址， 获取合约代码
	code := ""
	if address[0:8] == "Contract" {
		txhash := address[8:]
		txDetail, _ := db.GetFlatTxnDetailByHash(txhash);
		code = txDetail.Action.Data
	}


	return c.JSON(http.StatusOK, FormatResponse(struct{
		db.Account
		Value float64 `json:"value"`
		Code string `json:"code"`
	}{
		*account,
		value,
		code,
	}))
}

func GetAccountTxs(c echo.Context) error {
	address := c.Param("id")
	if address == "" {
		return errors.New("address requied")
	}

	page := c.QueryParam("p")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		pageInt = 1
	}

	start := (pageInt - 1) * AccountEachPage
	txnList, err := db.GetTxnListByAccount(address, start, AccountEachPage)
	if err != nil {
		return err
	}

	totalLen, _ := db.GetFlatTxnLenByAccount(address)
	pageLast := calLastPage(totalLen)

	output := &AccountTxsOutput{
		address,
		txnList,
		totalLen,
		pageLast,
	}

	return c.JSON(http.StatusOK, FormatResponse(output))
}

func ApplyIOST(c echo.Context) error {

	address := c.FormValue("address")
	email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	vc := c.FormValue("verify")

	sess, _ := session.GlobalSessions.SessionStart(c.Response(), c.Request())
	defer sess.SessionRelease(c.Response())

	if sess.SessionID() == "" {
		log.Println("ApplyIOST nil session id")
		return ErrInvalidInput
	}

	if address == "" || email == "" || mobile == "" || vc == "" {
		log.Println("ApplyIOST nil params")
		return ErrInvalidInput
	}

	//if len(address) != 44 && len(address) != 45 {
	//	log.Println("ApplyIOST invalid address")
	//	return ErrInvalidInput
	//}

	if len(mobile) < 10 || mobile[0] != '+' {
		log.Println("ApplyIOST invalid mobile")
		return ErrInvalidInput
	}

	if len(vc) != 6 {
		log.Println("ApplyIOST invalid vc")
		return ErrInvalidInput
	}

	//if len(common.Base58Decode(address)) != 33 {
	//	log.Println("ApplyIOST invalid decode address")
	//	return ErrInvalidInput
	//}

	if !RegEmail.MatchString(email) {
		log.Println("ApplyIOST invaild regexp email")
		return ErrInvalidInput
	}

	vcInterface := sess.Get("verification")
	vcInSession, _ := vcInterface.(string)

	log.Println("ApplyIOST:", sess.SessionID(), "vc:", vc)

	if strings.ToLower(vcInSession) != strings.ToLower(vc) {
		log.Println("ApplyIOST", ErrMobileVerfiy.Error())
		return ErrMobileVerfiy
	}

	// send to blockChain
	var (
		txHash        []byte
		err           error
		transferIndex int
	)
	for transferIndex < 3 { // 尝试 3 次
		txHash, err = db.TransferIOSTToAddress(address, 10.1*100000000)
		if err != nil {
			log.Println("ApplyIOST TransferIOSTToAddress error:", err)
		}
		if txHash != nil {
			break
		}
		transferIndex++
		time.Sleep(time.Second)
	}
	if transferIndex == 3 {
		log.Println("ApplyIOST TransferIOSTToAddress error:", ErrOutOfRetryTime)
		return ErrOutOfRetryTime
	}

	txHashEncoded := common.Base58Encode(txHash)

	// check BlocakChain
	var checkIndex int
	for checkIndex < 30 {
		time.Sleep(time.Second * 2)
		if _, err := blkchain.GetTxByHash(txHashEncoded); err != nil {
			log.Printf("ApplyIOST GetTxnByHash error: %v, retry...\n", err)
		} else {
			log.Println("ApplyIOST blockChain Hash: ", txHashEncoded)
			break
		}
		checkIndex++
	}

	if checkIndex == 30 {
		log.Println("ApplyIOST checkTxHash error:", ErrOutOfCheckTxHash)
		return ErrOutOfCheckTxHash
	}
	log.Println("ApplyIOST checkTxHash success.")

	ai := &db.ApplyTestIOST{
		Address:   address,
		Amount:    10,
		Email:     email,
		Mobile:    mobile,
		ApplyTime: time.Now().Unix(),
	}
	db.SaveApplyTestIOST(ai)

	return c.JSON(http.StatusOK, FormatResponse(txHashEncoded))
}

/*func ApplyIOSTBenMark(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	address := c.FormValue("address")
	email := c.FormValue("email")
	mobile := c.FormValue("mobile")

	if address == "" || email == "" || mobile == "" {
		log.Println("ApplyIOST nil params")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if len(address) != 44 && len(address) != 45 {
		log.Println("ApplyIOST invalid address")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if len(mobile) != 14 || mobile[0] != '+' {
		log.Println("ApplyIOST invalid mobile")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if len(common.Base58Decode(address)) != 33 {
		log.Println("ApplyIOST invalid decode address")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if !RegEmail.MatchString(email) {
		log.Println("ApplyIOST invaild regexp email")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	log.Println("ApplyIOST address:", address)

	// send to blockChain
	var (
		txHash        []byte
		err           error
		transferIndex int
	)
	for transferIndex < 3 {
		txHash, err = blockchain.TransferIOSTToAddress(address, 10)
		if err != nil {
			log.Println("ApplyIOST TransferIOSTToAddress error:", err)
		}
		if txHash != nil {
			break
		}
		transferIndex++
		time.Sleep(time.Second)
	}
	if transferIndex == 3 {
		log.Println("ApplyIOST TransferIOSTToAddress error:", ErrOutOfRetryTime)
		return c.JSON(http.StatusOK, &CommOutput{3, ErrOutOfRetryTime.Error()})
	}

	// check BlocakChain
	var checkIndex int
	for checkIndex < 30 {
		time.Sleep(time.Second * 2)
		if txx, err := blockchain.GetTxnByHash(txHash); err != nil {
			log.Printf("ApplyIOST GetTxnByHash error: %v, retry...\n", err)
		} else {
			log.Println("ApplyIOST blockChain Hash: ", bytes.Equal(txx.Hash(), txHash), hex.EncodeToString(txx.Hash()), hex.EncodeToString(txHash))
			break
		}
		checkIndex++
	}

	if checkIndex == 30 {
		log.Printf("ApplyIOST checkTxHash error: %v, address: %s\n", ErrOutOfCheckTxHash, address)
		return c.JSON(http.StatusOK, &CommOutput{4, ErrOutOfCheckTxHash.Error()})
	}
	log.Println("ApplyIOST checkTxHash success.")

	ai := &db.ApplyTestIOST{
		Address:   address,
		Amount:    10,
		Email:     email,
		Mobile:    mobile,
		ApplyTime: time.Now().Unix(),
	}
	db.SaveApplyTestIOST(ai)

	return c.JSON(http.StatusOK, &CommOutput{0, hex.EncodeToString(txHash)})
}*/

func TestPage(c echo.Context) error {
	sess, _ := session.GlobalSessions.SessionStart(c.Response(), c.Request())
	defer sess.SessionRelease(c.Response())

	sess.Set("test-session", "hello-123-456")
	return c.JSON(http.StatusOK, FormatResponse([]string{"hello world"}))
}

func TestPage2(c echo.Context) error {
	sess, _ := session.GlobalSessions.SessionStart(c.Response(), c.Request())
	defer sess.SessionRelease(c.Response())

	info := sess.Get("test-session")
	vcInSession, _ := info.(string)
	fmt.Println(vcInSession)
	return c.JSON(http.StatusOK, FormatResponse([]string{"hello world2"}))
}
