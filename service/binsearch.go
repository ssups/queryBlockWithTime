package service

import (
	"context"
	"log"
	"math/big"
	"utils/queryBlockWithTime/util"

	"github.com/ethereum/go-ethereum/ethclient"
)

type QueryTool struct {
	client *ethclient.Client
}

func NewQueryTool(c *ethclient.Client) *QueryTool {
	return &QueryTool{client: c}
}

func (q *QueryTool) BinarySearch(target, low, high uint64) (uint64, uint64) {
	if high < low {
		log.Fatal("High must be bigger than low")
	}
	if high == low {
		return high, (util.SeperateFatal(q.client.BlockByNumber(context.Background(), new(big.Int).SetUint64(high)))).Time()
	}

	mid := ((low + high) / 2)

	blocktimestamp := (util.SeperateFatal(q.client.BlockByNumber(context.Background(), new(big.Int).SetUint64(mid)))).Time()

	switch {
	case blocktimestamp > target:
		return q.BinarySearch(target, low, mid)
	case blocktimestamp < target:
		return q.BinarySearch(target, mid+1, high)
	default:
		return mid, blocktimestamp
	}
}
