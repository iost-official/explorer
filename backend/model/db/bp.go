package db

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
)

type bpStore struct {
	Id    string `bson:"_id" json:"_id"`
	Count int64  `bson:"count" json:"count"`
}

func GetBPProduce(pubkeys []string, startTime time.Time, endTime time.Time) ([]bpStore, error) {
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
	var results []bpStore
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
