package controller2

import (
	"github.com/iost-official/explorer/backend/model2/db"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type AccountsOutput struct {
	AccountList []*db.Account `json:"accountList"`
	Page        int64         `json:"page"`
	PagePrev    int64         `json:"pagePrev"`
	PageNext    int64         `json:"pageNext"`
	PageLast    int           `json:"pageLast"`
	TotalLen    int           `json:"totalLen"`
}

/*type AccountDetailOutput struct {
	Account *db.Account                `json:"account"`
	TxnList []*model.TransactionOutput `json:"txn_list"`
	TxnLen  int64                      `json:"txn_len"`
}*/

type AccountTxs struct {
	Address  string       `json:"address"`
	TxnList  []*db.FlatTx `json:"txnList"`
	TxnLen   int64        `json:"txnLen"`
	PageLast int64        `json:"pageLast"`
}

func init() {
	gcapHttpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 10,
		},
	}
}

func GetAccounts(c echo.Context) error {
	page := c.QueryParam("p")

	pageInt64, _ := strconv.ParseInt(page, 10, 64)
	if pageInt64 <= 0 {
		pageInt64 = 1
	}

	start := int((pageInt64 - 1) * AccountEachPage)
	accountList, err := db.GetAccounts(start, AccountEachPage)
	if err != nil {
		return err
	}

	accountTotalLen, err := db.GetAccountsTotalLen()
	if err != nil {
		return err
	}

	var lastPage int
	if accountTotalLen/AccountEachPage == 0 {
		lastPage = accountTotalLen / AccountEachPage
	} else {
		lastPage = accountTotalLen/AccountEachPage + 1
	}

	if lastPage > AccountMaxPage {
		lastPage = AccountMaxPage
	}

	output := &AccountsOutput{
		AccountList: accountList,
		Page:        pageInt64,
		PagePrev:    pageInt64 - 1,
		PageNext:    pageInt64 + 1,
		PageLast:    lastPage,
		TotalLen:    accountTotalLen,
	}

	return c.JSON(http.StatusOK, FormatResponse(output))
}

func GetAccountDetail(c echo.Context) error {
	// TODO 实时获取数据
	address := c.Param("id")

	account, err := db.GetAccountByAddress(address)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, FormatResponse(account))
}

func GetAccountTxs(c echo.Context) error {
	address := c.Param("id")
	if address == "" {
		return nil
	}

	page := c.QueryParam("p")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	if pageInt <= 0 {
		pageInt = 1
	}

	eachPage := 25
	start := (pageInt - 1) * eachPage
	txnList, err := db.GetTxnListByAccount(address, start, eachPage)
	if err != nil {
		return err
	}


	totalLen, _ := db.GetFlatTxnLenByAccount(address)

	txsInt64Len := int64(totalLen)

	var pageLast int64
	if txsInt64Len%30 == 0 {
		pageLast = txsInt64Len / 30
	} else {
		pageLast = txsInt64Len / 30
	}

	output := &AccountTxs{
		address,
		txnList,
		int64(totalLen),
		pageLast,
	}

	return c.JSON(http.StatusOK, FormatResponse(output))
}

/*func ApplyIOST(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	address := c.FormValue("address")
	email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	vc := c.FormValue("verify")

	sess, _ := session.GlobalSessions.SessionStart(c.Response(), c.Request())
	defer sess.SessionRelease(c.Response())

	if sess.SessionID() == "" {
		log.Println("ApplyIOST nil session id")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if address == "" || email == "" || mobile == "" || vc == "" {
		log.Println("ApplyIOST nil params")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if len(address) != 44 && len(address) != 45 {
		log.Println("ApplyIOST invalid address")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if len(mobile) < 10 || mobile[0] != '+' {
		log.Println("ApplyIOST invalid mobile")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if len(vc) != 6 {
		log.Println("ApplyIOST invalid vc")
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

	vcInterface := sess.Get("verification")
	vcInSession, _ := vcInterface.(string)

	log.Println("ApplyIOST:", sess.SessionID(), "vc:", vc)

	if strings.ToLower(vcInSession) != strings.ToLower(vc) {
		log.Println("ApplyIOST", ErrMobileVerfiy.Error())
		return c.JSON(http.StatusOK, &CommOutput{2, ErrMobileVerfiy.Error()})
	}

	// send to blockChain
	var (
		txHash        []byte
		err           error
		transferIndex int
	)
	for transferIndex < 3 {
		txHash, err = blockchain.TransferIOSTToAddress(address, 10.1)
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

	txHashEncoded := common.Base58Encode(txHash)

	// check BlocakChain
	var checkIndex int
	for checkIndex < 30 {
		time.Sleep(time.Second * 2)
		if _, err := blockchain.GetTxnByHash(txHash); err != nil {
			log.Printf("ApplyIOST GetTxnByHash error: %v, retry...\n", err)
		} else {
			log.Println("ApplyIOST blockChain Hash: ", txHashEncoded)
			break
		}
		checkIndex++
	}

	if checkIndex == 30 {
		log.Println("ApplyIOST checkTxHash error:", ErrOutOfCheckTxHash)
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

	return c.JSON(http.StatusOK, &CommOutput{0, txHashEncoded})
}*/

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
	return c.JSON(http.StatusOK, FormatResponse([]string{"hello world"}))
}
