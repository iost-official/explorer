package common

import (
	"encoding/json"
	"fmt"
	"github.com/GincoInc/iost-explorer/backend/model/db"
	"strconv"
)

func CalculateAward(ainfo db.AwardInfo) (err error) {
	currentAid := ainfo.Aid
	totalAward := ainfo.TotalAmount
	lastBlockNumber, err := db.GetLastBlockNumberBefore(ainfo.EndTime)
	if err != nil {
		return err
	}
	var firstBlockNumber int64
	firstBlockNumber, err = db.GetFirstBlockNumberAfter(ainfo.StartTime)
	if err != nil {
		return err
	}
	fmt.Println("fristBlock: ", firstBlockNumber, " lastBlock: ", lastBlockNumber)

	// Calculate producerTxs and userVote
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
					var amount float64
					amount, err = strconv.ParseFloat(params[2], 64)
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
					aidPidPair := AidPidPair{Aid: params[0], Pid: params[1]}
					userVote[aidPidPair] = append(userVote[aidPidPair], vAction)

				}
				break
			case "vote_producer.iost/unvote":
				receiptSucc = true
				var params []string
				err := json.Unmarshal([]byte(receipt.Content), &params)
				if err == nil && len(params) == 3 {
					var amount float64
					amount, err = strconv.ParseFloat(params[2], 64)
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
					aidPidPair := AidPidPair{Aid: params[0], Pid: params[1]}
					userVote[aidPidPair] = append(userVote[aidPidPair], vAction)
				}
				break
			case "vote_producer.iost/unregister":
				receiptSucc = true
				var params []string
				err := json.Unmarshal([]byte(receipt.Content), &params)
				if err == nil && len(params) == 1 {
					vAction := VoteAction{
						ActionType:  ActionUnRegister,
						To:          params[0],
						BlockNumber: vTx.BlockNumber,
					}
					producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
				}
				break
			case "vote_producer.iost/applyRegister":
				receiptSucc = true
				var params []interface{}
				err := json.Unmarshal([]byte(receipt.Content), &params)
				if err == nil && len(params) == 6 {
					vAction := VoteAction{
						ActionType:  ActionRegister,
						To:          params[0].(string),
						BlockNumber: vTx.BlockNumber,
					}
					producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
				} else {
					fmt.Println(err.Error())
				}
				break
			default:
				break

			}
		}
		if !receiptSucc {
			for _, action := range vTx.Actions {
				// ContractF3tLtxdXwYmKsDiUtTmaQztwJQLPVf9VyWDqufMZHP5p is LieBi VoteContract
				if action.Contract == "vote_producer.iost" || action.Contract == "ContractF3tLtxdXwYmKsDiUtTmaQztwJQLPVf9VyWDqufMZHP5p" {
					switch action.ActionName {
					case "vote":
						var params []string
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 3 {
							var amount float64
							amount, err = strconv.ParseFloat(params[2], 64)
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
							aidPidPair := AidPidPair{Aid: params[0], Pid: params[1]}
							userVote[aidPidPair] = append(userVote[aidPidPair], vAction)
						}
						break
					case "unvote":
						var params []string
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 3 {
							var amount float64
							amount, err = strconv.ParseFloat(params[2], 64)
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
							aidPidPair := AidPidPair{Aid: params[0], Pid: params[1]}
							userVote[aidPidPair] = append(userVote[aidPidPair], vAction)
						}
						break
					case "unregister":
						var params []string
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 1 {
							vAction := VoteAction{
								ActionType:  ActionUnRegister,
								To:          params[0],
								BlockNumber: vTx.BlockNumber,
							}
							producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
						}
						break
					case "applyRegister":
						var params []interface{}
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 6 {
							vAction := VoteAction{
								ActionType:  ActionRegister,
								To:          params[0].(string),
								BlockNumber: vTx.BlockNumber,
							}
							producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
						} else {
							fmt.Println(err.Error())
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
	producerVoteTotals := map[string]float64{}
	var totalVotes float64
	//Calculate Producer award and ValidTime
	for pid, voteActions := range producerTxs {
		var producerVote float64
		var producerVoteTotal float64
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
				if producerRegistered && !producerOnline && producerVote >= ProducerOnlineLimit {
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
				if producerRegistered && producerOnline && producerVote < ProducerOnlineLimit {
					producerOnlineList[pid] = appendProducerOnline(producerOnlineList[pid], producerOnlineStart, voteAction.BlockNumber, firstBlockNumber)
					producerOnline = false
				}

			case ActionRegister:
				producerRegistered = true
				if producerVote >= ProducerOnlineLimit {
					producerOnline = true
					producerOnlineStart = voteAction.BlockNumber
					producerVoteChangeLast = voteAction.BlockNumber
				}
			case ActionUnRegister:
				if producerRegistered && producerOnline {
					currentBlockNumber := voteAction.BlockNumber
					producerVoteTotal += calculateVotes(producerVoteChangeLast, currentBlockNumber, producerVote, firstBlockNumber)
					producerVoteChangeLast = currentBlockNumber
					producerOnlineList[pid] = appendProducerOnline(producerOnlineList[pid], producerOnlineStart, voteAction.BlockNumber, firstBlockNumber)
				}
				producerRegistered = false
				producerOnline = false
			}
		}
		//Currently Still online
		if producerRegistered && producerOnline {
			producerVoteTotal += calculateVotes(producerVoteChangeLast, lastBlockNumber, producerVote, firstBlockNumber)
			producerOnlineList[pid] = appendProducerOnline(producerOnlineList[pid], producerOnlineStart, lastBlockNumber, firstBlockNumber)
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

	//Calculate userVotes that counts
	userVotes := map[AidPidPair]float64{}
	for aidPidPair, voteActions := range userVote {
		var validVoteTime float64
		var voterLastVoteAmount float64
		var voterLastVote int64
		for _, voteAction := range voteActions {
			switch voteAction.ActionType {
			case ActionVote:
				currentBlock := voteAction.BlockNumber
				var timeInter int64
				for _, o := range producerOnlineList[aidPidPair.Pid] {
					if o.Start <= voterLastVote && o.End >= voterLastVote {
						var endBlock int64
						if currentBlock > o.End {
							endBlock = o.End
						} else {
							endBlock = currentBlock
						}
						timeInter += endBlock/AwardInterval - voterLastVote/AwardInterval
					} else if o.Start <= currentBlock && o.End >= currentBlock {
						var startBlock int64
						if voterLastVote < o.Start {
							startBlock = o.Start
						} else {
							startBlock = voterLastVote
						}
						timeInter += currentBlock/AwardInterval - startBlock/AwardInterval
					} else if o.Start >= voterLastVote && o.End <= currentBlock {
						timeInter += o.End/AwardInterval - o.Start/AwardInterval
					}

				}
				validVoteTime += voterLastVoteAmount * float64(timeInter)
				voterLastVote = currentBlock
				voterLastVoteAmount += voteAction.Amount
			case ActionUnvote:
				currentBlock := voteAction.BlockNumber
				var timeInter int64
				for _, o := range producerOnlineList[aidPidPair.Pid] {
					if o.Start <= voterLastVote && o.End >= voterLastVote {
						var endBlock int64
						if currentBlock > o.End {
							endBlock = o.End
						} else {
							endBlock = currentBlock
						}
						timeInter += endBlock/AwardInterval - voterLastVote/AwardInterval
					} else if o.Start <= currentBlock && o.End >= currentBlock {
						var startBlock int64
						if voterLastVote < o.Start {
							startBlock = o.Start
						} else {
							startBlock = voterLastVote
						}
						timeInter += currentBlock/AwardInterval - startBlock/AwardInterval
					} else if o.Start >= voterLastVote && o.End <= currentBlock {
						timeInter += o.End/AwardInterval - o.Start/AwardInterval
					}

				}
				validVoteTime += voterLastVoteAmount * float64(timeInter)
				voterLastVote = currentBlock
				voterLastVoteAmount -= voteAction.Amount
				if timeInter < 0 {
					fmt.Println("WTF! Inter<0: ", timeInter, ", aid: ", aidPidPair.Aid, "ProducerOnlne: ", producerOnlineList[aidPidPair.Pid])
				}
				if voterLastVoteAmount < 0 {
					fmt.Println("WTF! voteAmound <0 : ", voterLastVoteAmount, ", aid: ", aidPidPair.Aid, " voteAmound: ", voteAction.Amount, " pid: ", aidPidPair.Pid)
				}
			}
		}
		// Add last part
		{
			currentBlock := lastBlockNumber
			var timeInter int64
			for _, o := range producerOnlineList[aidPidPair.Pid] {
				if o.Start <= voterLastVote && o.End >= voterLastVote {
					var endBlock int64
					if currentBlock > o.End {
						endBlock = o.End
					} else {
						endBlock = currentBlock
					}
					timeInter += endBlock/AwardInterval - voterLastVote/AwardInterval
				} else if o.Start <= currentBlock && o.End >= currentBlock {
					var startBlock int64
					if voterLastVote < o.Start {
						startBlock = o.Start
					} else {
						startBlock = voterLastVote
					}
					timeInter += currentBlock/AwardInterval - startBlock/AwardInterval
				} else if o.Start >= voterLastVote && o.End <= currentBlock {
					timeInter += o.End/AwardInterval - o.Start/AwardInterval
				}
			}
			validVoteTime += voterLastVoteAmount * float64(timeInter)
			voterLastVote = currentBlock
		}
		userVotes[aidPidPair] = validVoteTime
	}
	var userAwards []db.UserAward
	for k, v := range userVotes {
		var userAward db.UserAward
		userAward = db.UserAward{
			Aid:      currentAid,
			Username: k.Aid,
			Pid:      k.Pid,
			Vote:     v,
			Award:    float64(v) * awardPerVote,
		}
		fmt.Println("userAward: ", userAward)
		userAwards = append(userAwards, userAward)
	}
	err = db.SaveProducerAward(producerAwards)
	err = db.SaveUserAward(userAwards)
	return nil
}


func CalculateProducerContributions(pInfo db.ProducerLevelInfo) (err error) {
	producerLevels := map[string]int{}
	var levelProducerCount [6]int
	for _, v := range pInfo.ProducerLevels {
		producerLevels[v.Pid] = v.Level
		switch v.Level {
		case 1:
			levelProducerCount[0]++
		case 2:
			levelProducerCount[1]++
		case 3:
			levelProducerCount[2]++
		case 4:
			levelProducerCount[3]++
		case 5:
			levelProducerCount[4]++
		case 6:
			levelProducerCount[5]++
		default:
			fmt.Println("haha:", v)
			return fmt.Errorf("ProducerLevel Error!")
		}
	}
	fmt.Println("ProducerLevels: ", producerLevels)

	ainfo, err := db.GetAwardInfo(pInfo.Aid)
	currentAid := ainfo.Aid
	lastBlockNumber, err := db.GetLastBlockNumberBefore(ainfo.EndTime)
	if err != nil {
		return err
	}
	var firstBlockNumber int64
	firstBlockNumber, err = db.GetFirstBlockNumberAfter(ainfo.StartTime)
	if err != nil {
		return err
	}
	//lastBlockNumber = 1000000000
	fmt.Println("fristBlock: ", firstBlockNumber, " lastBlock: ", lastBlockNumber)

	voteTxs, err := db.GetVoteTxs(lastBlockNumber)
	producerTxs := map[string][]VoteAction{}
	userVote := map[AidPidPair][]VoteAction{}
	for _, vTx := range voteTxs {
		receiptSucc := false
		for _, receipt := range vTx.TxReceipt.Receipts {
			switch receipt.FuncName {
			case "vote_producer.iost/vote":
				receiptSucc = true
				var params []string
				err := json.Unmarshal([]byte(receipt.Content), &params)
				if err == nil && len(params) == 3 {
					var amount float64
					amount, err = strconv.ParseFloat(params[2], 64)
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
					aidPidPair := AidPidPair{Aid: params[0], Pid: params[1]}
					userVote[aidPidPair] = append(userVote[aidPidPair], vAction)

				}
				break
			case "vote_producer.iost/unvote":
				receiptSucc = true
				var params []string
				err := json.Unmarshal([]byte(receipt.Content), &params)
				if err == nil && len(params) == 3 {
					var amount float64
					amount, err = strconv.ParseFloat(params[2], 64)
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
					aidPidPair := AidPidPair{Aid: params[0], Pid: params[1]}
					userVote[aidPidPair] = append(userVote[aidPidPair], vAction)
				}
				break
			case "vote_producer.iost/unregister":
				receiptSucc = true
				var params []string
				err := json.Unmarshal([]byte(receipt.Content), &params)
				if err == nil && len(params) == 1 {
					vAction := VoteAction{
						ActionType:  ActionUnRegister,
						To:          params[0],
						BlockNumber: vTx.BlockNumber,
					}
					producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
				}
				break
			case "vote_producer.iost/applyRegister":
				receiptSucc = true
				var params []interface{}
				err := json.Unmarshal([]byte(receipt.Content), &params)
				if err == nil && len(params) == 6 {
					vAction := VoteAction{
						ActionType:  ActionRegister,
						To:          params[0].(string),
						BlockNumber: vTx.BlockNumber,
					}
					producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
				} else {
					fmt.Println(err.Error())
				}
				break
			default:
				break

			}
		}
		if !receiptSucc {
			for _, action := range vTx.Actions {
				if action.Contract == "vote_producer.iost" || action.Contract == "ContractF3tLtxdXwYmKsDiUtTmaQztwJQLPVf9VyWDqufMZHP5p" {
					switch action.ActionName {
					case "vote":
						var params []string
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 3 {
							var amount float64
							amount, err = strconv.ParseFloat(params[2], 64)
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
							aidPidPair := AidPidPair{Aid: params[0], Pid: params[1]}
							userVote[aidPidPair] = append(userVote[aidPidPair], vAction)
						}
						break
					case "unvote":
						var params []string
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 3 {
							var amount float64
							amount, err = strconv.ParseFloat(params[2], 64)
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
							aidPidPair := AidPidPair{Aid: params[0], Pid: params[1]}
							userVote[aidPidPair] = append(userVote[aidPidPair], vAction)
						}
						break
					case "unregister":
						var params []string
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 1 {
							vAction := VoteAction{
								ActionType:  ActionUnRegister,
								To:          params[0],
								BlockNumber: vTx.BlockNumber,
							}
							producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
						}
						break
					case "applyRegister":
						var params []interface{}
						err := json.Unmarshal([]byte(action.Data), &params)
						if err == nil && len(params) == 6 {
							vAction := VoteAction{
								ActionType:  ActionRegister,
								To:          params[0].(string),
								BlockNumber: vTx.BlockNumber,
							}
							producerTxs[vAction.To] = append(producerTxs[vAction.To], vAction)
						} else {
							fmt.Println(err.Error())
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
	producerVoteTotals := map[string]float64{}
	var totalVotes float64
	//Calculate Producer award and ValidTime
	for pid, voteActions := range producerTxs {
		var producerVote float64
		var producerVoteTotal float64
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
				if producerRegistered && !producerOnline && producerVote >= ProducerOnlineLimit {
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
				if producerRegistered && producerOnline && producerVote < ProducerOnlineLimit {
					producerOnlineList[pid] = appendProducerOnline(producerOnlineList[pid], producerOnlineStart, voteAction.BlockNumber, firstBlockNumber)
					producerOnline = false
				}

			case ActionRegister:
				producerRegistered = true
				if producerVote >= ProducerOnlineLimit {
					producerOnline = true
					producerOnlineStart = voteAction.BlockNumber
					producerVoteChangeLast = voteAction.BlockNumber
				}
			case ActionUnRegister:
				if producerRegistered && producerOnline {
					currentBlockNumber := voteAction.BlockNumber
					producerVoteTotal += calculateVotes(producerVoteChangeLast, currentBlockNumber, producerVote, firstBlockNumber)
					producerVoteChangeLast = currentBlockNumber
					producerOnlineList[pid] = appendProducerOnline(producerOnlineList[pid], producerOnlineStart, voteAction.BlockNumber, firstBlockNumber)
				}
				producerRegistered = false
				producerOnline = false
			}
		}
		//Currently Still online
		if producerRegistered && producerOnline {
			producerVoteTotal += calculateVotes(producerVoteChangeLast, lastBlockNumber, producerVote, firstBlockNumber)
			producerOnlineList[pid] = appendProducerOnline(producerOnlineList[pid], producerOnlineStart, lastBlockNumber, firstBlockNumber)
		}
		producerVoteTotals[pid] = producerVoteTotal
		totalVotes += producerVoteTotal
	}
	var producerAwards []db.ProducerAward
	for k, v := range producerVoteTotals {
		var producerAward db.ProducerAward

		var pLevel int
		var ok bool
		var award float64
		fmt.Println("producerName: ", k)
		if pLevel, ok = producerLevels[k]; ok {
			pLevel--
			award = float64(levelAward[pLevel]) / float64(levelProducerCount[pLevel])
		} else {
			award = 0
		}
		producerAward = db.ProducerAward{
			Aid:   currentAid,
			Pid:   k,
			Vote:  v,
			Award: award,
		}
		fmt.Println("producerAward: ", producerAward)

		producerAwards = append(producerAwards, producerAward)
	}

	userVotes := map[AidPidPair]float64{}
	for aidPidPair, voteActions := range userVote {
		var validVoteTime float64
		var voterLastVoteAmount float64
		var voterLastVote int64
		for _, voteAction := range voteActions {
			switch voteAction.ActionType {
			case ActionVote:
				currentBlock := voteAction.BlockNumber
				var timeInter int64
				for _, o := range producerOnlineList[aidPidPair.Pid] {
					if o.Start <= voterLastVote && o.End >= voterLastVote {
						var endBlock int64
						if currentBlock > o.End {
							endBlock = o.End
						} else {
							endBlock = currentBlock
						}
						timeInter += endBlock/AwardInterval - voterLastVote/AwardInterval
					} else if o.Start <= currentBlock && o.End >= currentBlock {
						var startBlock int64
						if voterLastVote < o.Start {
							startBlock = o.Start
						} else {
							startBlock = voterLastVote
						}
						timeInter += currentBlock/AwardInterval - startBlock/AwardInterval
					} else if o.Start >= voterLastVote && o.End <= currentBlock {
						timeInter += o.End/AwardInterval - o.Start/AwardInterval
					}

				}
				validVoteTime += voterLastVoteAmount * float64(timeInter)
				voterLastVote = currentBlock
				voterLastVoteAmount += voteAction.Amount
			case ActionUnvote:
				currentBlock := voteAction.BlockNumber
				var timeInter int64
				for _, o := range producerOnlineList[aidPidPair.Pid] {
					if o.Start <= voterLastVote && o.End >= voterLastVote {
						var endBlock int64
						if currentBlock > o.End {
							endBlock = o.End
						} else {
							endBlock = currentBlock
						}
						timeInter += endBlock/AwardInterval - voterLastVote/AwardInterval
					} else if o.Start <= currentBlock && o.End >= currentBlock {
						var startBlock int64
						if voterLastVote < o.Start {
							startBlock = o.Start
						} else {
							startBlock = voterLastVote
						}
						timeInter += currentBlock/AwardInterval - startBlock/AwardInterval
					} else if o.Start >= voterLastVote && o.End <= currentBlock {
						timeInter += o.End/AwardInterval - o.Start/AwardInterval
					}

				}
				validVoteTime += voterLastVoteAmount * float64(timeInter)
				voterLastVote = currentBlock
				voterLastVoteAmount -= voteAction.Amount
			}
		}
		// Add last part
		{
			currentBlock := lastBlockNumber
			var timeInter int64
			for _, o := range producerOnlineList[aidPidPair.Pid] {
				if o.Start <= voterLastVote && o.End >= voterLastVote {
					var endBlock int64
					if currentBlock > o.End {
						endBlock = o.End
					} else {
						endBlock = currentBlock
					}
					timeInter += endBlock/AwardInterval - voterLastVote/AwardInterval
				} else if o.Start <= currentBlock && o.End >= currentBlock {
					var startBlock int64
					if voterLastVote < o.Start {
						startBlock = o.Start
					} else {
						startBlock = voterLastVote
					}
					timeInter += currentBlock/AwardInterval - startBlock/AwardInterval
				} else if o.Start >= voterLastVote && o.End <= currentBlock {
					timeInter += o.End/AwardInterval - o.Start/AwardInterval
				}
			}
			validVoteTime += voterLastVoteAmount * float64(timeInter)
			voterLastVote = currentBlock
		}
		userVotes[aidPidPair] = validVoteTime
	}
	var userAwards []db.UserAward
	for k, v := range userVotes {
		var userAward db.UserAward
		var pLevel int
		var ok bool
		var award float64
		if pLevel, ok = producerLevels[k.Pid]; ok {
			pLevel--
			award = (float64(v) / float64(producerVoteTotals[k.Pid])) * (float64(levelAward[pLevel]) / float64(levelProducerCount[pLevel]))
		} else {
			award = 0
		}
		if producerVoteTotals[k.Pid] == 0 {
			award = 0
		}
		userAward = db.UserAward{
			Aid:      currentAid,
			Username: k.Aid,
			Pid:      k.Pid,
			Vote:     v,
			Award:    award,
		}
		fmt.Println("userAward: ", userAward)
		userAwards = append(userAwards, userAward)
	}
	err = db.SaveProducerContributionAward(producerAwards)
	err = db.SaveUserContributionAward(userAwards)

	return nil
}

type VoteAction struct {
	ActionType  int
	From        string
	To          string
	Amount      float64
	BlockNumber int64
}

type AidPidPair struct {
	Aid string
	Pid string
}

const (
	ActionVote = iota
	ActionUnvote
	ActionRegister
	ActionUnRegister
)

const (
	AwardInterval       = 172800
	ProducerOnlineLimit = 2100000
)

type ProducerOnlineTime struct {
	Start int64
	End   int64
}

func calculateVotes(voteStart, voteEnd int64, voteAmount float64, blockStart int64) float64 {
	if voteEnd < blockStart {
		return 0
	}
	if voteStart < blockStart {
		voteStart = blockStart
	}
	return float64(voteEnd/AwardInterval-voteStart/AwardInterval) * voteAmount
}

func appendProducerOnline(currentList []ProducerOnlineTime, onlineStart, onlineEnd, blockStart int64) []ProducerOnlineTime {
	if onlineEnd < blockStart {
		return currentList
	}
	if onlineStart < blockStart {
		onlineStart = blockStart
	}
	if onlineStart > onlineEnd {
		fmt.Println("Error: ShouldNotBeGreater!")
	}
	return append(currentList, ProducerOnlineTime{Start: onlineStart, End: onlineEnd})
}

// old:
// 账号贡献奖励:共1%，用户一半，节点一半
// 一等：50%, 二等: 25%，类推
// x = 2.1e8 / 365 * 92 / 2
// x * 0.5, x * 0.25, x * 0.125, x*0.075, x*0.005 
//var levelAward            = []float64{26465753/2.0, 13232877/2.0, 6616438/2.0, 3969863/2.0, 2646575/2.0, 0}

// new:

var levelAwardTotal = 2.1e8 / 365 * 92 / 2
var levelAward = []float64{levelAwardTotal*0.5, levelAwardTotal*0.3, levelAwardTotal*0.15, levelAwardTotal*0.04, levelAwardTotal*0.01}
