package blockchain

import "testing"

func TestGetBlockByNumAndHash(t *testing.T) {
	blockRes, err := GetBlockByNum(1, true)
	if err != nil {
		t.Fatalf("GetBlockByNum error: %v\n", err)
	}
	if blockRes.Block.TxCount <= 0 {
		t.Fatalf("GetBlockByNum invalid tx count")
	}

	blockRes, err = GetBlockByHash(blockRes.Block.Hash, true)
	if err != nil {
		t.Fatalf("GetBlockByHash error: %v\n", err)
	}
	if blockRes.Block.TxCount <= 0 {
		t.Fatalf("GetBlockByHash invalid tx count")
	}
}
