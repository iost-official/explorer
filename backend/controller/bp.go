package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo"
)

const (
	chainInfoUrl    = "https://api.iost.io/getChainInfo"
	chainStorageUrl = "http://api.iost.io/getContractStorage"
)

type chainInfo struct {
	WitnessList []string `json:"witness_list"`
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
	bpAccountListLock sync.RWMutex
)

func init() {
	bpClient = &http.Client{
		Timeout: time.Second * 2,
		Transport: &http.Transport{
			MaxIdleConns: 5,
		},
	}

	accountInfoCh := make(chan string)
	accountDetailCh := make(chan *accountInfo)

	go func() {
		ticker := time.NewTicker(time.Second * 1)
		for _ = range ticker.C {
			witnessList, err := getWitnessList()
			if err != nil {
				log.Println("getWitnessList error:", err)
				continue
			}
			bpAccountListLock.RLock()
			for _, witness := range witnessList {
				if _, ok := bpAccountList[witness]; !ok {
					accountInfoCh <- witness
				}
			}
			bpAccountListLock.RUnlock()
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

func GetBPList(c echo.Context) error {
	accountInfo, err := json.Marshal(bpAccountList)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, FormatResponse(string(accountInfo)))
}

func getWitnessList() ([]string, error) {
	resp, err := bpClient.Get(chainInfoUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ci *chainInfo
	err = json.Unmarshal(body, &ci)
	if err != nil {
		return nil, err
	}

	return ci.WitnessList, nil
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
	fmt.Println("queryStr:", string(queryStr))

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
	fmt.Println("body: ", string(body))
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

	fmt.Println("data: ", dataStr)
	fmt.Println("ad: ", ad)
	fmt.Println("err: ", err)

	ad.Name = account

	return ad, nil
}
