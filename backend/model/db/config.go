package db

var (
	MongoLink = "mongodb://47.244.109.92:27017"
	Db        string
)

const (
	CollectionBlocks        = "blocks"
	CollectionTxs           = "txs"
	CollectionFlatTx        = "flatxs"
	CollectionAccount       = "accounts"
	CollectionAccountTx     = "accountTx"
	CollectionAccountPubkey = "accountPubkey"
	CollectionContract      = "contracts"
	CollectionContractTx    = "contractTx"
	CollectionTaskCursor    = "taskCursors"
	CollectionBlockPay      = "blockPays"
)
