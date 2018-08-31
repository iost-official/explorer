package db

//func GetTopBlock2() (*rpc.BlockInfo, error) {
//	collection, err := GetCollection("blocks")
//	if err != nil {
//		return nil, err
//	}
//
//	var emptyQuery interface{}
//	var topBlk *rpc.BlockInfo
//	err = collection.Find(emptyQuery).Sort("-head.blockNumber").Limit(1).One(&topBlk)
//	if err != nil {
//		log.Println("getTopBlock error:", err)
//		return nil, err
//	}
//
//	return topBlk, nil
//}
