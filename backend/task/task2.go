package main

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/db"
)

func main () {
	col, err := db.GetCollection("testCol")
	if err != nil {
		fmt.Println("error", err)
	}

	err = col.Insert(bson.M{"hello": "world"})
	fmt.Println("insert errror")
}
