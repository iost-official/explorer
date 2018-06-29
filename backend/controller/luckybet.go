package controller

import (
	"bytes"
	"encoding/hex"
	"explorer/model"
	"explorer/model/blockchain"
	"explorer/model/db"
	"github.com/iost-official/prototype/common"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

type BetBlockInfo struct {
	Height    int64 `json:"height"`
	Timestamp int64 `json:"timestamp"`
}

type BetInfo struct {
	Ret     int             `json:"ret"`
	Top6Blk []*BetBlockInfo `json:"top_6_blk"`
}

type SendBetInfo struct {
	Ret    int    `json:"ret"`
	Msg    string `json:"msg"`
	TxHash string `json:"tx_hash"`
	Nonce  int    `json:"nonce"`
}

type AddressBet struct {
	Address     string         `json:"address"`
	Nonce       int            `json:"nonce"`
	LuckyNumber int            `json:"lucky_number"`
	BetAmount   int            `json:"bet_amount"`
	BetTime     int64          `json:"bet_time"`
	Result      *db.BetWinUser `json:"result"`
}

type AddressBetOutput struct {
	AddressBetList []*AddressBet `json:"address_bet_list"`
	Page           int           `json:"page"`
	LastPage       int           `json:"last_page"`
}

type LatestBetOutput struct {
	BlockHeight   int64                   `json:"block_height"`
	BlockTime     int64                   `json:"block_time"`
	LatestBetList []*db.BetResultCommInfo `json:"latest_bet_list"`
}

type InsufficientBalanceOutput struct {
	Ret     int     `json:"ret"`
	Msg     string  `json:"msg"`
	Balance float64 `json:"balance"`
}

const BetResultEachPage = 10

var ErrInsufficientBalance = errors.New("Insufficient balance")

func GetBetInfo(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	top10Blks, err := model.GetBlock(1, 6)
	if err != nil {
		return err
	}

	output := &BetInfo{
		Ret: 0,
	}

	betTop10Blk := make([]*BetBlockInfo, len(top10Blks))
	for k, blk := range top10Blks {
		betTop10Blk[k] = &BetBlockInfo{blk.Height, blk.Timestamp}
	}

	output.Top6Blk = betTop10Blk

	return c.JSON(http.StatusOK, output)
}

func GetLuckyBet(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	address := c.FormValue("address")
	betAmount := c.FormValue("betAmount")
	luckyNumber := c.FormValue("luckyNumber")
	privKey := c.FormValue("privKey")
	gcaptcha := c.FormValue("gcaptcha")

	remoteip := c.Request().Header.Get("Iost_Remote_Addr")
	if !verifyGCAP(gcaptcha, remoteip) {
		log.Println(ErrGreCaptcha.Error())
		return c.JSON(http.StatusOK, &CommOutput{1, ErrGreCaptcha.Error()})
	}

	if address == "" || betAmount == "" || privKey == "" || luckyNumber == "" {
		log.Println("GetLuckyBet nil params")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	luckyNumberInt, err := strconv.Atoi(luckyNumber)
	if err != nil || (luckyNumberInt < 0 || luckyNumberInt > 9) {
		log.Println("GetLuckyBet invalud lucky number")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	betAmountInt, err := strconv.Atoi(betAmount)
	if err != nil || (betAmountInt <= 0 || betAmountInt > 5) {
		log.Println("GetLuckyBet invalud bet amount")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if len(address) != 44 && len(address) != 45 {
		log.Println("GetLuckyBet invalid address")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	//if len(common.Base58Decode(address)) != 33 {
	//	log.Println("GetLuckyBet invalid decode address")
	//	return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	//}

	//if len(common.Base58Decode(privKey)) != 32 {
	//	log.Println("GetLuckyBet invalid decode privKey")
	//	return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	//}

	balance, err := blockchain.GetBalanceByKey(address)
	if err != nil {
		log.Println("GetLuckyBet GetBalanceByKey error:", err)
	}
	if float64(betAmountInt) > balance {
		return c.JSON(http.StatusOK, &InsufficientBalanceOutput{6, ErrInsufficientBalance.Error(), balance})
	}

	// send to blockChain
	var (
		txHash               []byte
		transferIndex, nonce int
	)
	for transferIndex < 3 {
		txHash, nonce, err = blockchain.SendBet(address, privKey, luckyNumberInt, betAmountInt)
		if err != nil {
			log.Println("GetLuckyBet SendBet error:", err)
		}
		if txHash != nil {
			break
		}
		transferIndex++
		time.Sleep(time.Second)
	}
	txHashEncoded := common.Base58Encode(txHash)
	if transferIndex == 3 {
		log.Println("GetLuckyBet SendBet error:", ErrOutOfRetryTime)
		return c.JSON(http.StatusOK, &SendBetInfo{3, ErrOutOfRetryTime.Error(), txHashEncoded, nonce})
	}

	// check BlocakChain
	var checkIndex int
	for checkIndex < 30 {
		time.Sleep(time.Second * 2)
		if _, err := blockchain.GetTxnByHash(txHash); err != nil {
			//log.Printf("GetLuckyBet SendBet error: %v, retry...\n", err)
		} else {
			log.Println("GetLuckyBet blockChain Hash: ", txHashEncoded)
			break
		}
		checkIndex++
	}

	if checkIndex == 30 {
		log.Println("GetLuckyBet checkTxHash error:", ErrOutOfCheckTxHash)
		return c.JSON(http.StatusOK, &SendBetInfo{4, ErrOutOfCheckTxHash.Error(), txHashEncoded, nonce})
	}
	log.Println("GetLuckyBet checkTxHash success.")

	ba := &db.BetAddress{
		Address:     address,
		Nonce:       nonce,
		PrivKey:     privKey,
		LuckyNumber: luckyNumberInt,
		BetAmount:   betAmountInt,
		BetTime:     time.Now().Unix(),
		ClientIp:    remoteip,
	}
	db.SaveAddressBet(ba)

	return c.JSON(http.StatusOK, &SendBetInfo{0, hex.EncodeToString(txHash), txHashEncoded, nonce})
}

func GetLuckyBetBenchMark(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	address := c.FormValue("address")
	betAmount := c.FormValue("betAmount")
	luckyNumber := c.FormValue("luckyNumber")
	privKey := c.FormValue("privKey")
	//gcaptcha := c.FormValue("gcaptcha")

	remoteip := c.Request().Header.Get("Iost_Remote_Addr")
	//if !verifyGCAP(gcaptcha, remoteip) {
	//	log.Println(ErrGreCaptcha.Error())
	//	return c.JSON(http.StatusOK, &CommOutput{1, ErrGreCaptcha.Error()})
	//}

	if address == "" || betAmount == "" || privKey == "" || luckyNumber == "" {
		log.Println("GetLuckyBet nil params")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	luckyNumberInt, err := strconv.Atoi(luckyNumber)
	if err != nil || (luckyNumberInt < 0 || luckyNumberInt > 9) {
		log.Println("GetLuckyBet invalud lucky number")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	betAmountInt, err := strconv.Atoi(betAmount)
	if err != nil || (betAmountInt <= 0 || betAmountInt > 5) {
		log.Println("GetLuckyBet invalud bet amount")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if len(address) != 44 && len(address) != 45 {
		log.Println("GetLuckyBet invalid address")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if len(common.Base58Decode(address)) != 33 {
		log.Println("GetLuckyBet invalid decode address")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	if len(common.Base58Decode(privKey)) != 32 {
		log.Println("GetLuckyBet invalid decode privKey")
		return c.JSON(http.StatusOK, &CommOutput{1, ErrInvalidInput.Error()})
	}

	// send to blockChain
	var (
		txHash               []byte
		transferIndex, nonce int
	)
	for transferIndex < 3 {
		txHash, nonce, err = blockchain.SendBet(address, privKey, luckyNumberInt, betAmountInt)
		if err != nil {
			log.Println("GetLuckyBet SendBet error:", err)
		}
		if txHash != nil {
			break
		}
		transferIndex++
		time.Sleep(time.Second)
	}
	if transferIndex == 3 {
		log.Println("GetLuckyBet SendBet error:", ErrOutOfRetryTime)
		return c.JSON(http.StatusOK, &SendBetInfo{3, ErrOutOfRetryTime.Error(), common.Base58Encode(txHash), nonce})
	}

	// check BlocakChain
	var checkIndex int
	for checkIndex < 30 {
		time.Sleep(time.Second * 2)
		if txx, err := blockchain.GetTxnByHash(txHash); err != nil {
			//log.Printf("GetLuckyBet SendBet error: %v, retry...\n", err)
		} else {
			log.Println("GetLuckyBet blockChain Hash: ", bytes.Equal(txx.Hash(), txHash), hex.EncodeToString(txx.Hash()), hex.EncodeToString(txHash))
			break
		}
		checkIndex++
	}

	if checkIndex == 30 {
		log.Println("GetLuckyBet checkTxHash error:", ErrOutOfCheckTxHash)
		return c.JSON(http.StatusOK, &SendBetInfo{4, ErrOutOfCheckTxHash.Error(), common.Base58Encode(txHash), nonce})
	}
	log.Println("GetLuckyBet checkTxHash success.")

	ba := &db.BetAddress{
		Address:     address,
		Nonce:       nonce,
		PrivKey:     privKey,
		LuckyNumber: luckyNumberInt,
		BetAmount:   betAmountInt,
		BetTime:     time.Now().Unix(),
		ClientIp:    remoteip,
	}
	db.SaveAddressBet(ba)

	return c.JSON(http.StatusOK, &SendBetInfo{0, hex.EncodeToString(txHash), common.Base58Encode(txHash), nonce})
}

func GetAddressBet(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	address := c.Param("id")
	page := c.QueryParam("p")
	tag := c.QueryParam("t")

	pageInt, _ := strconv.Atoi(page)
	if pageInt <= 0 {
		pageInt = 1
	}

	limit := BetResultEachPage
	if tag == "t" {
		limit = 5
	}
	skip := (pageInt - 1) * limit
	top5Bet, err := db.GetAddressBet(address, skip, limit)
	if err != nil {
		return err
	}

	var nonceList []int
	for _, v := range top5Bet {
		nonceList = append(nonceList, v.Nonce)
	}

	var addressBetList []*AddressBet
	addressWinMap, err := db.GetAddressBetDetail(address, nonceList)

	for _, betInfo := range top5Bet {
		if err == nil && len(addressWinMap) > 0 {
			addressBetList = append(addressBetList, &AddressBet{
				Address:     betInfo.Address,
				Nonce:       betInfo.Nonce,
				LuckyNumber: betInfo.LuckyNumber,
				BetAmount:   betInfo.BetAmount,
				BetTime:     betInfo.BetTime,
				Result:      addressWinMap[betInfo.Nonce],
			})
		} else {
			addressBetList = append(addressBetList, &AddressBet{
				Address:     betInfo.Address,
				Nonce:       betInfo.Nonce,
				LuckyNumber: betInfo.LuckyNumber,
				BetAmount:   betInfo.BetAmount,
				BetTime:     betInfo.BetTime,
				Result:      nil})
		}
	}

	totalTimes, err := db.GetAddressBetTimes(address)
	if err != nil {
		log.Println("GetAddressBet GetAddressBetTimes error:", err)
	}

	var lastPage int
	if totalTimes%BetResultEachPage == 0 {
		lastPage = totalTimes / BetResultEachPage
	} else {
		lastPage = totalTimes/BetResultEachPage + 1
	}

	return c.JSON(http.StatusOK, &AddressBetOutput{addressBetList, pageInt, lastPage})
}

func GetLatestBetInfo(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	latestList, err := db.GetLatestRoundInfo(5)
	if err != nil {
		return err
	}

	output := new(LatestBetOutput)
	output.LatestBetList = latestList

	if blk, err := db.GetTopBlock(); err == nil && blk != nil {
		output.BlockHeight = blk.Head.Number
		output.BlockTime = model.ConvertSlotTimeToTimeStamp(blk.Head.Time)
	}

	return c.JSON(http.StatusOK, output)
}

func GetBetRound(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	round := c.Param("id")

	roundInt, err := strconv.Atoi(round)
	if err != nil || roundInt <= 0 {
		return nil
	}

	roundInfo, err := db.GetRound(roundInt)

	if err != nil {
		log.Println("GetBetRound error:", err)
		return err
	}

	return c.JSON(http.StatusOK, roundInfo)
}

func GetTodayTop10Address(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	winnerList, err := db.GetTop10AddressWithDay(time.Now().Unix())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, winnerList)
}
