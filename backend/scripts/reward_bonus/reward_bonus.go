package main

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/GincoInc/iost-explorer/backend/config"
	"github.com/GincoInc/iost-explorer/backend/model/db"
	"github.com/GincoInc/iost-explorer/backend/util"
	"github.com/iost-official/go-iost/account"
	"github.com/iost-official/go-iost/common"
	"github.com/iost-official/go-iost/crypto"
	"github.com/iost-official/go-iost/sdk"

	iostpb "github.com/iost-official/go-iost/rpc/pb"
)


type UserAward struct {
	UserName          string `bson:"userName"`
	ProducerID          string  `bson:"producerID"`
	Vote        int64 `bson:"vote"`
	Award float64 `bson:"award"`
}

type UserAwardFailed struct {
	UserName          string `bson:"userName"`
}

func award() error {
	// config and constants
	acc := "admin"
	byteKey, err := util.ReadPassword("Enter Password:  ")
	if err != nil {
		return err
	}
	privKey := string(byteKey)
	actionsNumPerTx := 500
	//gasPledgeAmount := 1000

	// prepare config and db
	config.ReadConfig("")
	awardC := db.GetCollection(db.CollectionUserAward)
	var userAwards []*UserAward

	// get tasks
	err = awardC.Find(bson.M{}).Sort("userName").All(&userAwards)
	if err != nil {
		return err
	}
	fmt.Printf("reward %v users\n", len(userAwards))

	// init sdk
	s := sdk.NewIOSTDevSDK()
	err = s.Connect()
	if err != nil {
		return err
	}
	s.SetTxInfo(4000000, 1, 90, 0, nil)
	kp, err := account.NewKeyPair(common.Base58Decode(privKey), crypto.Ed25519)
	if err != nil {
		return err
	}
	s.SetAccount(acc, kp)

	// pledge 1000 IOST for gas usage
	//action := sdk.NewAction("gas.iost", "pledge", fmt.Sprintf("[\"%s\",\"%s\",\"%d\"]", acc, acc, gasPledgeAmount))
	//_, err = s.SendTxFromActions([]*iostpb.Action{action})
	//if err != nil {
	//	return err
	//}

	i := 0
	failCount := 0
	for i < len(userAwards) {
		actions := make([]*iostpb.Action, 0)
		j := 0
		for j < actionsNumPerTx && i + j < len(userAwards) {
			actions = append(actions, sdk.NewAction("token.iost", "transfer",
				fmt.Sprintf("[\"iost\",\"%s\",\"%s\",\"%.8f\",\"\"]", acc, userAwards[i + j].UserName, userAwards[i + j].Award)))
			j++
		}
		if len(actions) == 0 {
			break
		}
		_, err := s.SendTxFromActions(actions)
		if err == nil {
			fmt.Printf("handle user %v to %v done\n", i, i + j)
		} else {
			failCount += j
			fmt.Printf("Error! handle user %v to %v failed %v\n", i, i + j, err)
			failedC := db.GetCollection(db.CollectionFailedAward)
			failedB := failedC.Bulk()
			j := 0
			for j < actionsNumPerTx && i + j < len(userAwards) {
				failedB.Insert(&UserAwardFailed{UserName:userAwards[i + j].UserName})
				j++
			}
			_, err = failedB.Run()
			if err != nil {
				return fmt.Errorf("insert to fail table err %v", err)
			}
		}
		i += j
	}
	if failCount > 0 {
		return fmt.Errorf("error count: %v, please check collection %v manually", failCount, db.CollectionFailedAward)
	}
	fmt.Printf("Award done!\n")
	return nil
}

func initFakeData() {

	c := db.GetCollection(db.CollectionUserAward)
	bulk := c.Bulk()

	offset := 0
	i := 0
	for i < 10000 {
		i++
		uname := fmt.Sprintf("uname_%d", offset + i)
		bulk.Insert(&UserAward{UserName:uname, Award:float64(0.01)})
	}
	_, err := bulk.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	initFakeData()
	err := award()
	fmt.Println(err)
}
