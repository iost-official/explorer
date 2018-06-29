package transport

import (
	"math/rand"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const EachGRPCClientNum = 5

var (
	grpcClientMap 		= make(map[string][]*grpc.ClientConn)
	grpcClientMapLock 	sync.RWMutex
)

func GetGRPCClient(address string) (*grpc.ClientConn, error) {
	grpcClientMapLock.RLock()
	if connList, ok := grpcClientMap[address]; ok {
		grpcClientMapLock.RUnlock()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rIndex := r.Intn(EachGRPCClientNum)
		return connList[rIndex], nil
	}
	grpcClientMapLock.RUnlock()

	grpcClientMapLock.Lock()
	defer grpcClientMapLock.Unlock()

	//  need improvement.
	connList, err := generateClientsWithNum(address, EachGRPCClientNum)
	if err != nil {
		return nil, err
	}
	grpcClientMap[address] = connList
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rIndex := r.Intn(EachGRPCClientNum)
	return connList[rIndex], nil
}

func generateClientsWithNum(address string, num int) ([]*grpc.ClientConn, error) {
	var (
		clientList []*grpc.ClientConn
		i int
	)

	for i < num {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		clientList = append(clientList, conn)
		i++
	}

	return clientList, nil
}