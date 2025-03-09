package client

import (
	"context"
	"encoding/json"

	"github.com/0xPolygonHermez/zkevm-bridge-service/jsonrpcclient/types"
	"github.com/ethereum/go-ethereum/common"
)

// ExitRootsByGER returns the exit roots accordingly to the provided Global Exit Root
func (c *Client) ExitRootsByGER(ctx context.Context, globalExitRoot common.Hash) (*types.ExitRoots, error) {
	response, err := JSONRPCCall(c.url, "zkevm_getExitRootsByGER", globalExitRoot.String())
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error.RPCError()
	}

	var result *types.ExitRoots
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetLatestGlobalExitRoot returns the latest global exit root
func (c *Client) GetLatestGlobalExitRoot(ctx context.Context) (common.Hash, error) {
	response, err := JSONRPCCall(c.url, "zkevm_getLatestGlobalExitRoot")
	if err != nil {
		return common.Hash{}, err
	}

	if response.Error != nil {
		return common.Hash{}, response.Error.RPCError()
	}

	var result string
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return common.Hash{}, err
	}

	return common.HexToHash(result), nil
}
