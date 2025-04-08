package claimtxman

import "github.com/ethereum/go-ethereum/common"

func (tm *NonceCache) decr(from common.Address) {
	nonce, err := tm.l2Node.NonceAt(tm.ctx, from, nil)
	if err != nil {
		return
	}
	if tempNonce, found := tm.nonceCache.Get(from.Hex()); found {
		tempNonce--
		if tempNonce >= nonce {
			tm.nonceCache.Add(from.Hex(), tempNonce)
		} else {
			tm.nonceCache.Remove(from.Hex())
		}
	}
}
