package main

import (
	"fmt"
	"github.com/iost-official/Go-IOS-Protocol/common"
	"github.com/iost-official/explorer/backend/model/blkchain"
	"github.com/iost-official/explorer/backend/model/dbv2"
	"log"
)


func main () {
	num, err := blkchain.GetCurrentBlockHeight()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chain current height", num)

	//fmt.Println("block info", block.Txs)
	//fmt.Println("block info", block.Txhash)
	lastNum, err := dbv2.GetLastBlockNumber()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db last number", lastNum)

	block1, err := blkchain.GetBlockByNum(num, false)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(block1.Head)
	//fmt.Println(block1.Head.ParentHash)
	//fmt.Println(block1.Hash)

	txHash := common.Base58Encode(block1.Txhash[0])
	fmt.Println(txHash)
	txRes, err := blkchain.GetTxByHash(txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txRes.Hash)
	fmt.Println(txRes.TxRaw.Signs)
	fmt.Println(txRes.TxRaw.Actions)

	//fmt.Println(block1.Txs)
	tx, err := dbv2.RpcGetTxByHash(txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("自定义TX: ", tx.ToFlatTx()[0])


	//return
	//collection, err := dbv2.GetCollection("block")
	//for i := lastNum; i <= num ; i++ {
	//	block, err := blkchain.GetBlockByNum(i, true)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	_block := dbv2.Block{
	//		ParentHash: common.Base58Encode(block.Head.ParentHash),
	//		Number: i,
	//	}
	//
	//	err = collection.Insert(_block)
	//	if err != nil {
	//		fmt.Println("insert error")
	//	}
	//	fmt.Println("current", common.Base58Encode(block.Head.ParentHash))
	//}
}
