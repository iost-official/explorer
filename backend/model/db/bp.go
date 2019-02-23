package db

import (
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
					"$lt": endTime.Nanosecond(),
					"$gt": startTime.Nanosecond(),
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

	return results, err
}
