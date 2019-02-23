package db

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
)

type BPStore struct {
	Id    string `bson:"_id" json:"_id"`
	Count int64  `bson:"count" json:"count"`
}
type BPProduce struct {
	Witness   string `bson:"witness" json:"witness"`
	StartTime int64  `bson:"startTime" json:"startTime"`
	EndTime   int64  `bson:"endTime" json:"endTime"`
	Count     int64  `bson:"count" json:"count"`
}

func GetBPProduceByStartTime(startTime int64) ([]BPProduce, error) {
	BPC := GetCollection(CollectionBP)
	query := bson.M{
		"startTime": startTime,
	}
	var bpProduces []BPProduce
	err := BPC.Find(query).All(&bpProduces)
	return bpProduces, err
}

func GetLastBPProduce() ([]BPProduce, error) {
	BPC := GetCollection(CollectionBP)
	var bpProduce BPProduce
	err := BPC.Find(nil).Sort("-startTime").One(&bpProduce)
	if err != nil {
		return nil, err
	}
	return GetBPProduceByStartTime(bpProduce.StartTime)
}

func GetBPProduce(pubkeys []string, startTime time.Time, endTime time.Time) ([]BPStore, error) {
	blockC := GetCollection(CollectionBlocks)
	query := []bson.M{
		bson.M{
			"$match": bson.M{
				"witness": bson.M{
					"$in": pubkeys,
				},
				"time": bson.M{
					"$lt": endTime.UnixNano(),
					"$gt": startTime.UnixNano(),
				},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":   "$witness",
				"count": bson.M{"$sum": 1},
			},
		},
	}
	var results []BPStore
	err := blockC.Pipe(query).All(&results)
	fmt.Println(results)
	if err != nil {
		return results, err
	}
	BPC := GetCollection(CollectionBP)
	for _, result := range results {
		bp := bson.M{
			"witness":   result.Id,
			"startTime": startTime.UnixNano(),
			"endTime":   endTime.UnixNano(),
			"count":     result.Count,
		}
		err = BPC.Insert(bp)
	}

	return results, err
}
