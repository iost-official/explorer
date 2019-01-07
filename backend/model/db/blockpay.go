package db

import (
	"github.com/globalsign/mgo/bson"
	"log"
)

type BlockPay struct {
	Height        int64   `json:"height" bson:"_id,omitempty"`
	AvgGasPrice   float64 `json:"avg_gas_price"`
	TotalGasLimit int64   `json:"total_gas_limit"`
}


func GetBlockPayListByHeight(heightList []int64) (map[int64]*BlockPay, error) {
	blkPC, err := GetCollection(CollectionBlockPay)
	if err != nil {
		log.Println("GetBlockPayByHeight get blockpay collection error:", err)
		return nil, err
	}

	query := bson.M{
		"_id": bson.M{
			"$in": heightList,
		},
	}

	var payList []*BlockPay
	err = blkPC.Find(query).All(&payList)
	if err != nil {
		return nil, err
	}

	payMap := make(map[int64]*BlockPay)
	for _, pay := range payList {
		payMap[pay.Height] = pay
	}

	return payMap, nil
}

func GetBlockPayByHeight(height int64) (*BlockPay, error) {
	payList, err := GetBlockPayListByHeight([]int64{height})
	if err != nil {
		return nil, err
	}

	return payList[0], nil
}

func GetTopBlockPay() (*BlockPay, error) {
	blkPC, err := GetCollection(CollectionBlockPay)
	if err != nil {
		log.Println("GetTopBlockPay get blockpay collection error:", err)
		return nil, err
	}

	var (
		emptyQuery   interface{}
		topPayDetail *BlockPay
	)
	err = blkPC.Find(emptyQuery).Sort("-_id").Limit(1).One(&topPayDetail)
	if err != nil {
		log.Println("GetTopBlockPay error:", err)
		return nil, err
	}

	return topPayDetail, nil
}
