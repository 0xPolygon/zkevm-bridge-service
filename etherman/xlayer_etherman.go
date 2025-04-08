package etherman

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
)

func (etherMan *Client) getFullBlockTime(vLog types.Log) time.Time {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fullBlock, err := etherMan.EtherClient.BlockByHash(ctx, vLog.BlockHash)
	if err != nil {
		log.Warnf("error getting hashParent. BlockNumber: %d. Error: %v", vLog.BlockNumber, err)
		return time.Now()
	}
	return time.Unix(int64(fullBlock.Time()), 0)
}
