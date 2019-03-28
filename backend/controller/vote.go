package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/iost-official/explorer/backend/model/db"
	"github.com/labstack/echo"
)

const (
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

func SetVoteAwardInfo(c echo.Context) (err error) {
	aInfo := new(db.AwardInfo)
	if err = c.Bind(aInfo); err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	aInfo.Aid = strconv.FormatInt(aInfo.StartTime, 10) + strconv.FormatInt(aInfo.EndTime, 10) + strconv.FormatInt(aInfo.TotalAmount, 10)

	if err = db.SaveAwardInfo(*aInfo); err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}

	return CalculateAward(c, *aInfo)
}

func GetUserAward(c echo.Context) error {
	id := c.QueryParam("aid")
	format := c.QueryParam("format")

	userAward, err := db.GetUserAward(id)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(nil))
	}
	var userAwardsOutput []interface{}
	for _, ua := range userAward {
		userAwardOutput := []interface{}{ua.Username, ua.Pid, ua.Vote, ua.Award}
		userAwardsOutput = append(userAwardsOutput, userAwardOutput)
	}
	if format == "1" {
		return c.JSON(http.StatusOK, FormatResponse(userAwardsOutput))
	} else {
		return c.JSON(http.StatusOK, FormatResponse(userAward))
	}
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
	format := c.QueryParam("format")
	producerAward, err := db.GetProducerAward(id)
	if err != nil {
		return c.JSON(http.StatusOK, FormatResponseFailed(err.Error()))
	}
	var producerAwardsOutput []interface{}
	for _, pa := range producerAward {
		producerAwardOutput := []interface{}{pa.Pid, pa.Vote, pa.Award}
		producerAwardsOutput = append(producerAwardsOutput, producerAwardOutput)
	}
	if format == "1" {
		return c.JSON(http.StatusOK, FormatResponse(producerAwardsOutput))
	} else {
		return c.JSON(http.StatusOK, FormatResponse(producerAward))
	}
}

func CalculateAward(c echo.Context, ainfo db.AwardInfo) (err error) {
	currentAid := ainfo.Aid
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
	fmt.Println("fristBlock: ", firstBlockNumber, " lastBlock: ", lastBlockNumber)

	voteTxs, err := db.GetVoteTxs(lastBlockNumber)
	producerTxs := map[string][]VoteAction{}
	userVote := map[AidPidPair][]VoteAction{}
	fmt.Println("fristBlock: ", firstBlockNumber, " lastBlock: ", lastBlockNumber)
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
	producerOnlineList := map[string][]ProducerOnlineTime{}
	producerVoteTotals := map[string]int64{}
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
					producerVoteTotal += calculateVotes(producerVoteChangeLast, currentBlockNumber, producerVote, firstBlockNumber)
					producerVoteChangeLast = currentBlockNumber
				}
				producerVote += voteAction.Amount
				if producerRegistered && !producerOnline && producerVote > producerOnlineLimit {
					producerOnline = true
					producerOnlineStart = voteAction.BlockNumber
					producerVoteChangeLast = voteAction.BlockNumber
				}

			case ActionUnvote:
				if producerRegistered && producerOnline {
					currentBlockNumber := voteAction.BlockNumber
					producerVoteTotal += calculateVotes(producerVoteChangeLast, currentBlockNumber, producerVote, firstBlockNumber)
					producerVoteChangeLast = currentBlockNumber
				}
				producerVote -= voteAction.Amount
				if producerRegistered && producerOnline && producerVote < producerOnlineLimit {
					producerOnlineList[pid] = append(producerOnlineList[pid], ProducerOnlineTime{Start: producerOnlineStart, End: voteAction.BlockNumber})
					producerOnline = false
				}

			case ActionRegister:
				producerRegistered = true
				if producerVote > producerOnlineLimit {
					producerOnline = true
					producerOnlineStart = voteAction.BlockNumber
					producerVoteChangeLast = voteAction.BlockNumber
				}
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
			producerVoteTotal += calculateVotes(producerVoteChangeLast, lastBlockNumber, producerVote, firstBlockNumber)
			producerOnlineList[pid] = append(producerOnlineList[pid], ProducerOnlineTime{Start: producerOnlineStart, End: lastBlockNumber})
		}
		producerVoteTotals[pid] = producerVoteTotal
		totalVotes += producerVoteTotal
	}
	awardPerVote := float64(totalAward) / float64(totalVotes)
	var producerAwards []db.ProducerAward
	for k, v := range producerVoteTotals {
		var producerAward db.ProducerAward
		producerAward = db.ProducerAward{
			Aid:   currentAid,
			Pid:   k,
			Vote:  v,
			Award: float64(v) * awardPerVote,
		}
		fmt.Println("producerAward: ", producerAward)

		producerAwards = append(producerAwards, producerAward)

	}

	userVotes := map[AidPidPair]int64{}
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
		// Add last part
		{
			currentBlock := lastBlockNumber
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
				} else if o.Start > voterLastVote && o.End < currentBlock {
					timeInter += o.End/awardInterval - o.Start/awardInterval
				}
			}
			validVoteTime += voterLastVoteAmount * timeInter
			voterLastVote = currentBlock
		}
		userVotes[aidPidPair] = validVoteTime
	}
	var userAwards []db.UserAward
	for k, v := range userVotes {
		var userAward db.UserAward
		userAward = db.UserAward{
			Aid:      currentAid,
			Username: k.aid,
			Pid:      k.pid,
			Vote:     v,
			Award:    float64(v) * awardPerVote,
		}
		fmt.Println("userAward: ", userAward)
		userAwards = append(userAwards, userAward)
	}
	err = db.SaveProducerAward(producerAwards)
	err = db.SaveUserAward(userAwards)

	return c.JSON(http.StatusOK, FormatResponse(currentAid))
}

func calculateVotes(voteStart, voteEnd, voteAmount, blockStart int64) int64 {
	if voteEnd < blockStart {
		return 0
	}
	if voteStart < blockStart {
		voteStart = blockStart
	}
	return (voteEnd/awardInterval - voteStart/awardInterval) * voteAmount
}
