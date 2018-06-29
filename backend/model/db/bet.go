package db

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
	"explorer/model/cache"
)

type BetAddress struct {
	Address     string `json:"address"`
	Nonce       int    `json:"nonce"`
	PrivKey     string `json:"priv_key"`
	LuckyNumber int    `json:"lucky_number"`
	BetAmount   int    `json:"bet_amount"`
	BetTime     int64  `json:"bet_time"`
	ClientIp    string `json:"client_ip"`
}

type BetResultCommInfo struct {
	Round        int     `json:"round"`
	BlockHeight  int64   `json:"BlockHeight"`
	TotalUserNum int     `json:"TotalUserNumber"`
	WinUserNum   int     `json:"WinUserNumber"`
	TotalRewards float64 `json:"TotalRewards"`
	WinTime      int64   `json:"win_time"`
}

type BetWinUser struct {
	Round       int     `json:"round"`
	Address     string  `json:"Address"`
	Amount      float64 `json:"Amount"`
	BetAmount   float64 `json:"BetAmount"`
	Nonce       int     `json:"Nonce"`
	LuckyNumber int     `json:"LuckyNumber"`
	IsWin       bool    `json:"is_win"`
	OnChainTime int64 `json:"on_chain_time"`
}

type BetResult struct {
	*BetResultCommInfo
	WinUserList   map[string]*BetWinUser `json:"WinUserList"`
	UnWinUserList map[string]*BetWinUser `json:"UnWinUserList"`
}

type BetResultFormat struct {
	*BetResultCommInfo
	*BetWinUser
}

type DayBetWinner struct {
	Address       string  `json:"_id"`
	TotalWinIOST  float64 `json:"totalWinIOST"`
	TotalWinTimes int     `json:"totalWinTimes"`
}

var robotAddressList = []string{"23hJissnRLwMcGFcPwyDxDfj9FaB5Z7LkY13n5TGZ2gL5","iQ68dENWhAoAPKPa5oqvdyyhrMVJWKSAALM5wL6DaJw9","yAw8vVunPfKrxSFCJPaxSkj4uurYNyxGcegCvyNhLkUW","y5fZWCfJsV4kvc2RyEAxmiTg6VjX5kdsreYNnjVnshmu","2AdmgVyJ5ystHaqcPjcjCeyt1T2A5mHDuGkyF12gmegQk","mn8oJ62wmRQUiRhXiw9BJKu7Y9PErwdLH4fyemTF492h","p3mA75VB39xKdj1CEbdigmxCJrS3aN3QHv3LB9Bg7btv","yXFopFVsjiRnznaHz4DS94gGp7ESiBGZE9z6FSb4khTd","26RnmyAXcVYyLkQroLKxAPeE5ffnXfMiP7f5uV3UqxDhc","rgfg7KmMeQyiDzMCEGCc6qrpsfdwGC6MDSwAFPo5Knt1","kjy5343149WyYy4PYQPYavfXPNXqV8HKCvxcu7zeLwGC","28eRp4MEJXEoNYGqdnJ72zLdciRJgCTUjs82mchtcay3J","tfnSXCTb3iPukhGMR3eQ9FBK8FMnwSXZevi6zt33mXhv","26777e13XSygCdvBJfPqvvva93r6Q53qJF9gHixeZqy4B","c8Hu3GAyQk8FM6wHiNs737ithZaNmMrKcYx2iZC84T77","vpByWtccdakv3mGjq8YCxYkAx5mR11b39D2j4i2taNtH","2AjmPnrRFE2EEbMtV7Hoph6RQSPaLmoK8MdsxPGDFHXwi","22M7R5KWpe1zH2Q2qGu5iqJHNYLzsT7JtXrwtHDhChcg2","yhuuF5qdTxEv7Lezx2t8kWa58fzevrzpVRw4UobuFcSM","xmJdxGFQ7ErZh5EqFWp19np7YWju1Y2HETmDGmn5B4ei","218qJFa5aFsNnQMaD5a48727d2enGXDv3NXiTG8ENUB5E","p9sAXisoKXk2JMUwyhRND2AzRRPR7y5Xqo4EQTYU5Rhp","xxAb6wTvEn7bFKzAKwsYQ3LYKk7PA6ZD8ARJYSyS9qqQ","rKZnLc5ybyrzgpMVUQDN13QeApe45xDt5ZCXbxwVpPH5","234SUfSHrEze1Z1UN3YoJTxrNuziTf6fhHradQQRkncF2","w3d7bxCKheCuy7aKoYG9jzZybM8CERi91VxTqGRWisoD","hyyvK2mBsCxuSLnVMYkmxaUH3FedABdv73zJsdATL9Vc","c2gJ5v1xLpiqYqXvK8pe31ivFiuTs7qKjaMPcJqPmhwC","h1V7zyNqSeA4ykwUHaEMvfNJsLC8cE63H7PKuUAJPFFU","m9nNARjLZHVDRvPtgETo3hk3tPt8csdEuRWKUMebWnsJ","sgoLawSskcdNezEHSR1cjW9WBtSGXvjfeZnZb9mkekuc","kYfWGAdsYNhzfSBpyS6c29e5EG1AzdaSSgNWFra5yRm5","29orQxYpoiFZmcfoms29JvmTLCzKgoPmiU2kDf6NDpEpx","fccwWR5AvTMHKL94kUYPJfu8DHbvViaK9mupPQJJKHVc","vU72w1hBhJwZrpYv6qfpvj4n6sWtsL6wW1xK2eLAXUxo","j8NCUnrhsxmWhrQVQbyqFYdjKwQmBJ4VfjZEKyJ8cdUh","caBZKqf8BfcYjydmqsHVFXZVvLry5jzTXaAT13yQoCrB","hfNWNzLM1aDW5mNUj4YKupyiQ6oYViXA1zB9xZTS92zw","yFXHN3YdiXsZrEvWAPLDhbSSpHxycdHhgaaujADfKSga","nPGLceu49G1MEbTkPekWLmQbqDu9k2uJKCaSTUS524QS","qSGdkUyrRGZo6xWxAGwSBrKGd1sDtDnoxbyyxXUBSBua","x3NsEpH2dNSCYWKryYd2LmKXLJUXHbTcWnzWXUmRzSi9","tBCiLVGx4DJ53wc7x8e1vCFZhkcDyVfgAET9gECYJeyX","qPKSw34v9SeoBboNL4R1XCymXGNDoHfPuRZiDuNKrtc7","yyyhw6oRbv4CP1XCX9xoZPu9xdifbE25xZ8U86Cvvc57","mo6DqtLm2e6KB5Pfshjd4N9g4qnhw6Bbz2CeV2kUtFbU","oHcqfvShKxmbyHXD7xmRjVMw5EwCSUieYbWzp1MRcSfo","2BavaTX5fkPfu7wpK56F1aV9EekG2UGqUdKQmHvYiwiJY","gAT86FDPyocFe79m3JHCHxMHUCpi58jPEEm2wWs2PtpV","cCtr3P2wLn8PigtDMRfYhTXEgrbLVSbJeo5YWd8gJR88","xA7eoJRYTRjyubJhPcZSrerF58AfBVZX9TKATPP1wr6J","qctmYhARoYgwcTGjFgrwxReU7YYKgHCJm9XSimULgdNm","zu75fCFZcwt26kBfQwQV5ttUn29bBVNomQMVtkf51BzV","oPFqX3CcQ8e8gv5PSvRNDed6tapULLTLbvuEhVCYQq81","twseF7MZHXAVjDm9kgeeiRJQrn1NmLZMa5DoTWjFoUVq","pc6vrCCRr8PM4L8Se6KeretrNBsoQzpFkBwejWnEKy9R","pNNibqorNipxTLJobAnp7kyNgwymkR7VPgUXZHhQoNdu","zZ98M1S4sZ1MA9xpc3gLRy5U96F9VMzr25JmUbwGjuBv","qverDM4JBJT8ApkEGxSQTdYTbVP7XoN1G77e7Zh5yD4n","jKiiZmDBccMvhFchdC4jL3jPxmweiDQTSQTPSW7set4e","zbQN3MBTgEzn97xqmEjTiQbqG8EhyMGkR7pVWrEeMzs9","wD2ny3sNqttFgZnEVu7UWnJwLeNgwRcTPFPLQAAAAmVX","crQwpYr6csfuPHpibY7LWBfRwxkQTpnE8ijbA4Xgq3J8","22zgKXxBBdc85L5Jbg8X7vVfRdJGCB5xXtexSKAi8ooH8","bYYTDeLrmiQtZshLk6b87aq7yZFR6J198GJ4ruZ818bZ","caeR7puE7AkupTQwgbx1H14PbiiAEgb9f2Udz814tvc2","29kJXCK3DUQpz9aJ7fpf7D9vPUUfMhXxBuBWkA3hYa2zJ","242juxLU78QHdqWkesdDZSRxGj463UVyEtMiqCZq5EGSv","23oSdHjdLnrMrH2jRZRZ7fEGdKugyXDsbRnLdnmwWGu9r","22THVA7ScJcfFeFc8KCeEnUYbkft4Tk2puhjS3PaR3zEc","oz6mMhW7dHzfeEK3jWD35hX7v8iSSsrEaC6eCxgHbvv9","xEUMmLr3LYNN5xUNxPtisTtGfquKgRNVaJG9R2XyXp23","23jmzULwYTseJHCjWHx546xGKpXnWrqUQioGfjwM2nYto","jGMELeWRcFhquLN8YfUkvVxdPKjedfnjYha29YTybpja","gNQHtgG5qL27uuT8GmxfHEHR1deNwMZGyYg96YmmotTQ","kaV69PLsiE7a311CqcNLSQxbrzPSddk6zENz7VsEUbZb","jGbhDg4avV4eH8Q87GUanzpDFcAewSKhWsGEZbdkrvyE","fLyk2xZZhA9boXJUwoYWW8zxfQtiyGdShgBAkMn18DbV","kPFCfpmxv3k7zRJxMFB6UfHxr3jHjexCBQrZgZi7Y5UQ","zRXUubxyp4BbmmdjmYgMv2ceyVDSDhB39ePxLPdJuKgi","mrBtmmcL3Nepywyx9mH11Cr6CVrEoAKs9ik7HNnW244K","295PHddmC8h7n4GtAyp9knXvAecfsBRmEmN1qf7GgA2rV","23F8Gee9BNaRAuHmhhKXoip1WqsfUkns8TcPR5fniViyG","25eU1PAsL8AbU1NeqddeumbAKeFcVajBdpmW5REZz6tWi","dp5wxs1xKYmQGQX38VDoUM3SL1SBmCskYgfHuM9ZouoZ","258DESEygNhREdTr6Pviy6fRU3XHkkctKuHMBNujKcnxY","qE954c2p4BTihaJbWc5RLMa67S8JX2cAZbxv9p1jbqfu","vvdfSEzELRUo7KFBCRMSKWSMyr9LqVi4389uQGYZDoym","pTqiChe4jozF625FViyzE9wV34LhwoMxSF3MTHLhaBwQ","wN8NiTLcuCNtzSw46tUnVnYtJB6dsxkSbnbQvnwPx9Ku","bbxumz61HFdXe3dfY2cZnTwjmcb8uZsDxwNE67TXHqJX","xtFL5yhyNBt8egKb6SKJEGWm5hrq48bkqnmtBXqA3Mhh","22M5QtLcfhsy7eCT37diubz4NsbGBH8MtJzQBbsYean2D","2173azESS6CXJGgwwevm9QTP7vMoMgy2RkdSSodHHfeo2","wC4mY2Uh1iaKb6CZVGg2wAQZxneKnKTtwHkZJmJ3yPVt","27xqcPRLdCEqZuAqfPPMajVUoBtJ6eQ5tPMuVKpAEnb71","24aHmprwcBpNtzRHNoATzRyXcjPufok3x6qaYMeRxMkJ1","vT552dPv8Z9JEnq6cFfSdVR6kYNV6id3Gn8ySr1CwGgM","fbtaNAcwNXD6LWFP3jp4Ai89QYt2LHEiaLbysZUGbmVc","nyXbjCqEfUYFGNDb5RjvWFyCgLhG5FXrCvWwQZzwEZ23","h6F29p52q35u4Q3LKJmyiyqXLQLnCGpNtJ7fCKoKmiuN"}

func SaveAddressBet(bet *BetAddress) error {
	bAC, err := GetCollection("betAddress")
	if err != nil {
		log.Println("SaveAddressBet get collection error:", err)
		return err
	}

	return bAC.Insert(bet)
}

func GetAddressBet(address string, skip, limit int) ([]*BetAddress, error) {
	bAC, err := GetCollection("betAddress")
	if err != nil {
		log.Println("GetAddressBet get collection error:", err)
		return nil, err
	}

	query := bson.M{
		"address": address,
	}
	var betAddressList []*BetAddress
	err = bAC.Find(query).Sort("-bettime").Skip(skip).Limit(limit).All(&betAddressList)

	return betAddressList, err
}

func GetAddressBetDetail(address string, nonceList []int) (map[int]*BetWinUser, error) {
	resultAddressC, err := GetCollection("betResultAddress")
	if err != nil {
		log.Println("GetAddressWin get collection error:", err)
		return nil, err
	}

	query := bson.M{
		"address": address,
		"nonce": bson.M{
			"$in": nonceList,
		},
	}

	var betDetailList []*BetWinUser
	err = resultAddressC.Find(query).All(&betDetailList)
	if err != nil {
		log.Println("GetAddressWin get addres bet result error:", err)
		return nil, err
	}

	nonceMap := make(map[int]*BetWinUser)
	for _, v := range betDetailList {
		nonceMap[v.Nonce] = v
	}

	return nonceMap, nil
}

func GetTopRound() (*BetResultCommInfo, error) {
	resultC, err := GetCollection("betResults")
	if err != nil {
		log.Println("GetTopRound get collection error:", err)
		return nil, err
	}

	var query interface{}
	var rs *BetResultCommInfo
	err = resultC.Find(query).Sort("-round").Limit(1).One(&rs)

	return rs, err
}

func GetRound(round int) ([]interface{}, error) {
	resultAddressC, err := GetCollection("betResultAddress")
	if err != nil {
		log.Println("GetRound get collection error:", err)
		return nil, err
	}

	queryPip := []bson.M{
		bson.M{
			"$match": bson.M{
				"round": round,
				"iswin": true,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":           "$address",
				"totalWinIOST":  bson.M{"$sum": "$amount"},
				"totalWinTimes": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{
				"totalWinIOST": -1,
			},
		},
	}

	var betWinList []interface{}
	err = resultAddressC.Pipe(queryPip).All(&betWinList)

	return betWinList, err
}

func GetLatestRoundInfo(roundLen int) ([]*BetResultCommInfo, error) {
	topRound, err := GetTopRound()
	if err != nil {
		return nil, err
	}

	resultC, err := GetCollection("betResults")
	if err != nil {
		log.Println("GetTopRound get collection error:", err)
		return nil, err
	}

	query := bson.M{
		"round": bson.M{
			"$gt": topRound.Round - roundLen,
		},
	}

	var roundInfoList []*BetResultCommInfo
	err = resultC.Find(query).Sort("-round").All(&roundInfoList)
	if err != nil {
		return nil, err
	}

	return roundInfoList, nil
}

func GetTop10AddressWithDay(daytime int64) ([]interface{}, error) {
	if top10Interface, ok := cache.GlobalCache.Get("top10DayBets"); ok {
		if top10, ok := top10Interface.([]interface{}); ok {
			return top10, nil
		}
	}

	resultAddressC, err := GetCollection("betResultAddress")
	if err != nil {
		log.Println("GetRound get collection error:", err)
		return nil, err
	}

	t := time.Unix(daytime, 0)
	dayBegin := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	dayEnd := dayBegin + 24*3600

	queryPip := []bson.M{
		bson.M{
			"$project": bson.M{
				"onchaintime": 1,
				"address": 1,
				"amount":  1,
				"betamount": 1,
			},
		},
		bson.M{
			"$match": bson.M{
				"onchaintime": bson.M{
					"$gte": dayBegin,
					"$lt":  dayEnd,
				},
				"address": bson.M{
					"$nin": robotAddressList,
				},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":           "$address",
				"totalWinIOST":  bson.M{"$sum": "$amount"},
				"totalBet": bson.M{"$sum": "$betamount"},
				"totalWinTimes": bson.M{"$sum": 1},
				//"netEarn": bson.M{"$subtract": []bson.M{bson.M{"$sum": "$amount"}, bson.M{"$sum": "$betamount"}}},
			},
		},
		bson.M{
			"$addFields": bson.M{
				"netEarn": bson.M{"$subtract": []string{"$totalWinIOST", "$totalBet"}},
			},
		},
		bson.M{
			"$sort": bson.M{
				"netEarn": -1,
			},
		},
		bson.M{
			"$limit": 10,
		},
	}

	var top10DayBetWinners []interface{}
	err = resultAddressC.Pipe(queryPip).All(&top10DayBetWinners)

	if err == nil {
		cache.GlobalCache.Set("top10DayBets", top10DayBetWinners, time.Minute * 2)
	}

	return top10DayBetWinners, err
}

func GetAddressBetTimes(address string) (int, error) {
	bAC, err := GetCollection("betAddress")
	if err != nil {
		log.Println("GetAddressBet get collection error:", err)
		return 0, err
	}

	query := bson.M{
		"address": address,
	}
	return bAC.Find(query).Count()
}