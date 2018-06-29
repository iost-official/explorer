package cron

import (
	"sync"
	"explorer/model/db"
	"log"
	"time"
	"explorer/model/blockchain"
	"encoding/json"
	"strings"
	"fmt"
)

func UpdateBet(wg *sync.WaitGroup)  {
	defer wg.Done()

	resultC, err := db.GetCollection("betResults")
	if err != nil {
		log.Println("updateBet get collection error:", err)
		return
	}

	resultAddressC, err := db.GetCollection("betResultAddress")
	if err != nil {
		log.Println("updateBet get collection error:", err)
		return
	}

	ticker := time.NewTicker(time.Second * 2)
	for _ = range ticker.C {
		var topRoundInMgo int
		topRoundInfoInMgo, err := db.GetTopRound()
		if err != nil {
			if err.Error() != "not found" {
				log.Println("UpdateBet getTopRound error:", err)
				continue
			}
		} else {
			topRoundInMgo = topRoundInfoInMgo.Round
		}

		topRoundInStat, err := blockchain.GetTotalRounds(blockchain.BetHash)
		if err != nil {
			log.Println("UpdateBet blockchain getTotalRounds error:", err)
			continue
		}
		log.Println("UpdateBet topRoundInMgo:", topRoundInMgo, "topRoundInStat:", topRoundInStat)

		if topRoundInStat == 0 {
			log.Println("UpdateBet no new round found, continue...")
			continue
		}

		if topRoundInMgo == int(topRoundInStat) {
			log.Println("UpdateBet no new round found, continue...")
			continue
		}

		for topRoundInMgo < int(topRoundInStat) {
			topRoundInMgo++

			roundStrInfo, err := blockchain.GetRoundWithNumber(blockchain.BetHash, topRoundInMgo)
			if err != nil {
				log.Println("UpdateBet GetRoundWithNumber error:", err)
				continue
			}

			fmt.Println("roundInfo:", roundStrInfo)
			var betResult *db.BetResult
			err = json.Unmarshal([]byte(strings.TrimLeft(roundStrInfo, "s")), &betResult)
			if err != nil {
				log.Println("UpdateBet json Unmarshal error:", err, "str:", roundStrInfo)
				continue
			}

			betResult.Round = topRoundInMgo
			betResult.BetResultCommInfo.WinTime = time.Now().Unix()

			err = resultC.Insert(betResult.BetResultCommInfo)
			if err != nil {
				log.Println("UpdateBet insert BetResultCommInfo error:", err)
				continue
			}

			var userList []interface{}
			if len(betResult.WinUserList) > 0 {
				for _, v := range betResult.WinUserList {
					v.Round = topRoundInMgo
					v.IsWin = true
					v.OnChainTime = time.Now().Unix()
					userList = append(userList, v)
				}
			}
			if len(betResult.UnWinUserList) > 0 {
				for _, v := range betResult.UnWinUserList {
					v.Round = topRoundInMgo
					v.IsWin = false
					v.OnChainTime = time.Now().Unix()
					userList = append(userList, v)
				}
			}

			err = resultAddressC.Insert(userList...)
			if err != nil {
				log.Println("UpdateBet insert resultBet error:", err)
				continue
			}
			log.Println("UpdateBet insert success, len:", len(userList))
		}
	}
}
