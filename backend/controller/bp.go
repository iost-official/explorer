package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/iost-official/explorer/backend/model/blockchain"

	"github.com/iost-official/explorer/backend/model/db"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

const (
	chainInfoUrl        = "https://api.iost.io/getChainInfo"
	chainStorageUrl     = "http://api.iost.io/getContractStorage"
	awardInterval       = 172800
	producerOnlineLimit = 2100000
	ActionVote          = iota
	ActionUnvote
	ActionRegister
	ActionUnRegister
)

var (
	totalAward      int64 = 10000
	lastBlockNumber int64 = 2100000
)

type ProducerOnlineTime struct {
	Start int64
	End   int64
}
type AidPidPair struct {
	aid string
	pid string
}
type VoteAction struct {
	ActionType  int
	From        string
	To          string
	Amount      int64
	BlockNumber int64
}
type chainInfo struct {
	WitnessList    []string `json:"witness_list"`
	LibWitnessList []string `json:"lib_witness_list"`
}

type accountInfo struct {
	Name   string `json:"data"`
	PubKey string `json:"pub_key"`
}

type BPAccountDetail struct {
	Name       string `json:"name"`
	Pubkey     string `json:"pubkey"`
	Loc        string `json:"loc"`
	Url        string `json:"url"`
	NetId      string `json:"netId"`
	IsProducer bool   `json:"isProducer"`
	Status     int    `json:"status"`
	Online     bool   `json:"online"`
}
type bpRegist struct {
	Id             string `json:"id"`
	Key            string `json:"key"`
	Field          string `json:"field"`
	ByLongestChain bool   `json:"by_longest_chain"`
}
type bpAccountQuery struct {
	Id             string `json:"id"`
	Key            string `json:"key"`
	Field          string `json:"field"`
	ByLongestChain bool   `json:"by_longest_chain"`
}

var (
	bpClient           *http.Client
	errAccountNotExist = errors.New("account not exists")

	bpAccountList     = make(map[string]*BPAccountDetail)
	bpLibAccountList  []string
	bpLibStartTime    = time.Now()
	bpAccountListLock sync.RWMutex
)

func init() {
	bpClient = &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			MaxIdleConns: 5,
		},
	}
}

func SyncBP() {
	accountInfoCh := make(chan string)
	accountDetailCh := make(chan *accountInfo)

	syncBP := viper.GetBool("syncBP")

	go func() {
		ticker := time.NewTicker(time.Second * 20)
		for _ = range ticker.C {
			witnessList, libWitnessList, err := getWitnessList()
			if err != nil {
				log.Println("getWitnessList error:", err)
				continue
			}

			if !reflect.DeepEqual(libWitnessList, bpLibAccountList) && syncBP {
				if len(bpLibAccountList) == 0 {
					bpLibStartTime = time.Now()
					bpLibAccountList = libWitnessList
				} else {
					go func(libWitnessList []string) {
						endTime := time.Now()
						_, err := db.GetBPProduce(bpLibAccountList, bpLibStartTime, endTime)
						if err != nil {
							log.Println("GetBPProduce error: ", err)
						}
						bpLibAccountList = libWitnessList
						bpLibStartTime = endTime
					}(libWitnessList)
				}
			}

			witnessMap := make(map[string]bool)
			for _, witness := range witnessList {
				witnessMap[witness] = true
			}

			var accountToQuery []string
			bpAccountListLock.Lock()

			for account := range bpAccountList {
				if _, ok := witnessMap[account]; !ok {
					delete(bpAccountList, account)
				}
			}

			for witness := range witnessMap {
				if _, ok := bpAccountList[witness]; !ok {
					accountToQuery = append(accountToQuery, witness)
				}
			}
			bpAccountListLock.Unlock()

			for _, account := range accountToQuery {
				accountInfoCh <- account
			}
		}
	}()

	go func() {
		for {
			select {
			case pubkey := <-accountInfoCh:
				ai, err := getBPAccount(pubkey)
				if err != nil {
					log.Println("getBPAccount error:", err)
					continue
				}

				accountDetailCh <- ai
			}
		}
	}()

	go func() {
		for {
			select {
			case ai := <-accountDetailCh:
				ad, err := getBPAccountDetail(ai.Name)
				if err != nil {
					log.Println("getBPAccountDetail error:", err)
					continue
				}

				bpAccountListLock.Lock()
				bpAccountList[ai.PubKey] = ad
				bpAccountListLock.Unlock()
			}
		}
	}()
}

func SetVoteAwardInfo(c echo.Context) (err error) {
	aInfo := new(db.AwardInfo)
	if err = c.Bind(aInfo); err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	aInfo.Aid = strconv.FormatInt(aInfo.StartTime, 10) + strconv.FormatInt(aInfo.EndTime, 10) + strconv.FormatInt(aInfo.TotalAmount, 10)

	if err = db.SaveAwardInfo(*aInfo); err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, FormatResponse(aInfo.Aid))
}

func GetUserAward(c echo.Context) error {
	id := c.QueryParam("aid")

	userAward, err := db.GetUserAward(id)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(nil))
	}
	return c.JSON(http.StatusOK, FormatResponse(userAward))
}

func GetVoteAwardList(c echo.Context) (err error) {
	voteAwardList, err := db.GetVoteAwardList()
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(nil))
	}
	return c.JSON(http.StatusOK, FormatResponse(voteAwardList))
}

func GetVoteAwardInfo(c echo.Context) (err error) {
	id := c.QueryParam("aid")
	voteAward, err := db.GetVoteAwardInfo(id)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, FormatResponse(voteAward))
}

func GetProducerAward(c echo.Context) (err error) {
	id := c.QueryParam("aid")
	producerAward, err := db.GetProducerAward(id)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, FormatResponse(producerAward))
}

func CalculateAward(c echo.Context) (err error) {
	currentAid := c.QueryParam("aid")
	ainfo, err := db.GetAwardInfo(currentAid)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	totalAward = ainfo.TotalAmount
	lastBlockNumber, err = db.GetLastBlockNumberBefore(ainfo.EndTime)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	var firstBlockNumber int64
	firstBlockNumber, err = db.GetFirstBlockNumberAfter(ainfo.StartTime)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}

	voteTxs, err := db.GetVoteTxs(firstBlockNumber, lastBlockNumber)
	var producerTxs map[string][]VoteAction
	var userVote map[AidPidPair][]VoteAction
	for _, vTx := range voteTxs {
		receiptSucc := false
		for _, receipt := range vTx.TxReceipt.Receipts {
			switch receipt.FuncName {
			case "vote_producer.iost/vote":
				receiptSucc = true
				var params []string
				err := json.Unmarshal([]byte(receipt.Content), &params)
				if err == nil && len(params) == 3 {
					var amount int64
					amount, err = strconv.ParseInt(params[2], 10, 64)
					if err != nil {
						break
					}
					vAction := VoteAction{
						ActionType:  ActionVote,
						From:        params[0],
						To:          params[1],
						Amount:      amount,
						BlockNumber: vTx.BlockNumber,
					}
					producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
					aidPidPair := AidPidPair{aid: params[0], pid: params[1]}
					userVote[aidPidPair] = append(userVote[aidPidPair], vAction)

				}
				break
			case "vote_producer.iost/unvote":
				receiptSucc = true
				var params []string
				err := json.Unmarshal([]byte(receipt.Content), &params)
				if err == nil && len(params) == 3 {
					var amount int64
					amount, err = strconv.ParseInt(params[2], 10, 64)
					if err != nil {
						break
					}
					vAction := VoteAction{
						ActionType:  ActionUnvote,
						From:        params[0],
						To:          params[1],
						Amount:      amount,
						BlockNumber: vTx.BlockNumber,
					}
					producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
					aidPidPair := AidPidPair{aid: params[0], pid: params[1]}
					userVote[aidPidPair] = append(userVote[aidPidPair], vAction)
				}
				break
			case "vote_producer.iost/unregister":
				receiptSucc = true
				var params []string
				err := json.Unmarshal([]byte(receipt.Content), &params)
				if err == nil && len(params) == 1 {
					vAction := VoteAction{
						ActionType: ActionUnRegister,
					}
					producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
				}
				break
			case "vote_producer.iost/applyRegister":
				receiptSucc = true
				var params []string
				err := json.Unmarshal([]byte(receipt.Content), &params)
				if err == nil && len(params) == 6 {
					vAction := VoteAction{
						ActionType: ActionRegister,
					}
					producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
				}
				break
			default:
				break

			}
		}
		if !receiptSucc {
			for _, action := range vTx.Actions {
				if action.Contract == "vote_producer.iost" {
					switch action.ActionName {
					case "vote":
						var params []string
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 3 {
							var amount int64
							amount, err = strconv.ParseInt(params[2], 10, 64)
							if err != nil {
								break
							}
							vAction := VoteAction{
								ActionType:  ActionVote,
								From:        params[0],
								To:          params[1],
								Amount:      amount,
								BlockNumber: vTx.BlockNumber,
							}
							producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
							aidPidPair := AidPidPair{aid: params[0], pid: params[1]}
							userVote[aidPidPair] = append(userVote[aidPidPair], vAction)
						}
						break
					case "unvote":
						var params []string
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 3 {
							var amount int64
							amount, err = strconv.ParseInt(params[2], 10, 64)
							if err != nil {
								break
							}
							vAction := VoteAction{
								ActionType:  ActionUnvote,
								From:        params[0],
								To:          params[1],
								Amount:      amount,
								BlockNumber: vTx.BlockNumber,
							}
							producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
							aidPidPair := AidPidPair{aid: params[0], pid: params[1]}
							userVote[aidPidPair] = append(userVote[aidPidPair], vAction)
						}
						break
					case "unregister":
						var params []string
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 1 {
							vAction := VoteAction{
								ActionType: ActionUnRegister,
							}
							producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
						}
						break
					case "applyRegister":
						var params []string
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 6 {
							vAction := VoteAction{
								ActionType: ActionRegister,
							}
							producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
						}
						break
					default:
						break
					}
				}
			}
		}
	}
	var producerOnlineList map[string][]ProducerOnlineTime
	var producerVoteTotals map[string]int64
	var totalVotes int64
	//Calculate Producer award and ValidTime
	for pid, voteActions := range producerTxs {
		var producerVote int64
		var producerVoteTotal int64
		producerRegistered := true
		producerOnline := false
		var producerOnlineStart int64
		var producerVoteChangeLast int64
		for _, voteAction := range voteActions {
			switch voteAction.ActionType {
			case ActionVote:
				if producerRegistered && producerOnline {
					currentBlockNumber := voteAction.BlockNumber
					producerVoteTotal += (currentBlockNumber/awardInterval - producerVoteChangeLast/awardInterval) * producerVote
					producerVoteChangeLast = currentBlockNumber
				}
				producerVote += voteAction.Amount
				if producerRegistered && !producerOnline && producerVote > producerOnlineLimit {
					producerOnline = true
					producerOnlineStart = voteAction.BlockNumber
					producerVoteChangeLast = voteAction.BlockNumber
				}

			case ActionUnvote:
				producerVote -= voteAction.Amount
				if producerRegistered && producerOnline && producerVote < producerOnlineLimit {
					producerOnlineList[pid] = append(producerOnlineList[pid], ProducerOnlineTime{Start: producerOnlineStart, End: voteAction.BlockNumber})
					producerOnline = false
				}
			case ActionRegister:
				producerRegistered = true
			case ActionUnRegister:
				if producerRegistered && producerOnline {
					producerOnlineList[pid] = append(producerOnlineList[pid], ProducerOnlineTime{Start: producerOnlineStart, End: voteAction.BlockNumber})
				}
				producerRegistered = false
				producerOnline = false
			}
		}
		//Currently Still online
		if producerRegistered && producerOnline {
			producerVoteTotal += (lastBlockNumber/awardInterval - producerVoteChangeLast/awardInterval) * producerVote
			producerOnlineList[pid] = append(producerOnlineList[pid], ProducerOnlineTime{Start: producerOnlineStart, End: lastBlockNumber})
		}
		producerVoteTotals[pid] = producerVoteTotal
		totalVotes += producerVoteTotal
	}
	awardPerVote := float64(totalAward) / float64(totalVotes)
	var producerAwards []db.ProducerAward
	for k, v := range producerVoteTotals {
		producerAwards = append(producerAwards, db.ProducerAward{
			Aid:   currentAid,
			Pid:   k,
			Vote:  v,
			Award: float64(v) * awardPerVote,
		})

	}

	var userVotes map[AidPidPair]int64
	for aidPidPair, voteActions := range userVote {
		var validVoteTime int64
		var voterLastVoteAmount int64
		var voterLastVote int64
		for _, voteAction := range voteActions {
			switch voteAction.ActionType {
			case ActionVote:
				currentBlock := voteAction.BlockNumber
				var timeInter int64
				for _, o := range producerOnlineList[aidPidPair.pid] {
					if o.Start < voterLastVote && o.End > voterLastVote {
						var endBlock int64
						if currentBlock > o.End {
							endBlock = o.End
						} else {
							endBlock = currentBlock
						}
						timeInter += endBlock/awardInterval - voterLastVote/awardInterval
					} else if o.Start < currentBlock && o.End > currentBlock {
						var startBlock int64
						if voterLastVote < o.Start {
							startBlock = o.Start
						} else {
							startBlock = voterLastVote
						}
						timeInter += currentBlock/awardInterval - startBlock/awardInterval
					}
				}
				validVoteTime += voterLastVoteAmount * timeInter
				voterLastVote = currentBlock
				voterLastVoteAmount += voteAction.Amount
			case ActionUnvote:
				currentBlock := voteAction.BlockNumber
				var timeInter int64
				for _, o := range producerOnlineList[aidPidPair.pid] {
					if o.Start < voterLastVote && o.End > voterLastVote {
						var endBlock int64
						if currentBlock > o.End {
							endBlock = o.End
						} else {
							endBlock = currentBlock
						}
						timeInter += endBlock/awardInterval - voterLastVote/awardInterval
					} else if o.Start < currentBlock && o.End > currentBlock {
						var startBlock int64
						if voterLastVote < o.Start {
							startBlock = o.Start
						} else {
							startBlock = voterLastVote
						}
						timeInter += currentBlock/awardInterval - startBlock/awardInterval
					}
				}
				validVoteTime += voterLastVoteAmount * timeInter
				voterLastVote = currentBlock
				voterLastVoteAmount -= voteAction.Amount
			}
		}
		userVotes[aidPidPair] = validVoteTime
	}
	var userAwards []db.UserAward
	for k, v := range userVotes {
		userAwards = append(userAwards, db.UserAward{
			Aid:      currentAid,
			Username: k.aid,
			Pid:      k.pid,
			Vote:     v,
			Award:    float64(v) * awardPerVote,
		})
	}
	err = db.SaveProducerAward(producerAwards)
	err = db.SaveUserAward(userAwards)

	return c.JSON(http.StatusOK, FormatResponse(""))
}

func RegistBP(c echo.Context) (err error) {
	accountName := c.QueryParam("accountName")
	bpURL := c.QueryParam("bpURL")
	region := c.QueryParam("region")
	password := c.QueryParam("password")
	usingAccount := "wtf123"
	targetIP := blockchain.RPCAddress
	if password != blockchain.BPRegisterPassword {
		return c.JSON(http.StatusOK, FormatResponseFailed(struct {
			Message string
			TxHash  string
		}{
			"Password Incorrect!",
			"",
		}))
	}

	cmd := exec.Command("bash", "-c", "iwallet -s "+targetIP+" sys producer-register "+accountName+" --target "+accountName+" --url "+bpURL+" --location "+region+" --partner --account "+usingAccount)
	fmt.Println(cmd.Args)

	var b bytes.Buffer
	var outputB []byte
	b.Write(blockchain.PasswordKey)
	cmd.Stdin = &b
	outputB, err = cmd.Output()
	txHashReg := regexp.MustCompile(`The transaction hash is: (.*)`)
	errorReg := regexp.MustCompile(`(?ms)throw new Error(.*)`)
	txHash := txHashReg.FindStringSubmatch(string(outputB))
	errorString := errorReg.FindString(string(outputB))
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(struct {
			Message string
			TxHash  string
		}{
			errorString,
			txHash[1],
		}))
	}
	return c.JSON(http.StatusOK, FormatResponse(struct {
		Message string
		TxHash  string
	}{
		errorString,
		txHash[1],
	}))
}

func GetBPLastProducer(c echo.Context) (err error) {
	startTimeS := c.QueryParam("startTime")
	var bp []db.BPProduce

	if startTimeS != "" {
		startTime, _ := strconv.ParseInt(startTimeS, 10, 64)
		bp, err = db.GetBPProduceByStartTime(startTime)
	} else {
		bp, err = db.GetLastBPProduce()
	}

	sort.Slice(bp, func(i, j int) bool {
		return bp[i].Witness < bp[j].Witness
	})

	return c.JSON(http.StatusOK, FormatResponse(bp))
}

func GetBPList(c echo.Context) error {
	var bpList []*BPAccountDetail

	bpAccountListLock.RLock()
	for _, bpAccount := range bpAccountList {
		bpList = append(bpList, bpAccount)
	}
	bpAccountListLock.RUnlock()

	sort.Slice(bpList, func(i, j int) bool {
		return bpList[i].Name < bpList[j].Name
	})

	return c.JSON(http.StatusOK, FormatResponse(bpList))
}

func getWitnessList() (witnessList, libWitnessList []string, err error) {
	resp, err := bpClient.Get(chainInfoUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var ci *chainInfo
	err = json.Unmarshal(body, &ci)
	if err != nil {
		return
	}

	return ci.WitnessList, ci.LibWitnessList, nil
}

func getBPAccount(pubKey string) (*accountInfo, error) {
	query := &bpAccountQuery{
		Id:    "vote_producer.iost",
		Key:   "producerKeyToId",
		Field: pubKey,
	}

	queryStr, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", chainStorageUrl, bytes.NewBuffer(queryStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := bpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ai *accountInfo
	err = json.Unmarshal(body, &ai)
	if err != nil {
		return nil, err
	}

	ai.Name = strings.Trim(ai.Name, "\"")
	ai.PubKey = pubKey

	return ai, nil
}

func getBPAccountDetail(account string) (*BPAccountDetail, error) {
	query := &bpAccountQuery{
		Id:    "vote_producer.iost",
		Key:   "producerTable",
		Field: account,
	}

	queryStr, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", chainStorageUrl, bytes.NewBuffer(queryStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := bpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ai map[string]interface{}
	err = json.Unmarshal(body, &ai)
	if err != nil {
		return nil, err
	}

	data, ok := ai["data"]
	if !ok {
		return nil, errAccountNotExist
	}
	dataStr, ok := data.(string)
	if !ok || len(dataStr) == 0 {
		return nil, errAccountNotExist
	}

	var ad *BPAccountDetail
	err = json.Unmarshal([]byte(dataStr), &ad)
	if err != nil {
		return nil, err
	}

	ad.Name = account

	return ad, nil
}
