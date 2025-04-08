package pushtask

import "github.com/0xPolygonHermez/zkevm-bridge-service/log"

func init() {
	log.Info("pushtask.minVerifyDuration = ", minVerifyDuration)
	log.Info("pushtask.defaultVerifyDuration = ", defaultVerifyDuration)
	log.Info("pushtask.maxVerifyDuration = ", maxVerifyDuration)
	log.Info("pushtask.verifyDurationListLen = ", verifyDurationListLen)

	log.Info("pushtask.minCommitDuration", minCommitDuration)
	log.Info("pushtask.defaultCommitDuration", defaultCommitDuration)
	log.Info("pushtask.commitDurationListLen", commitDurationListLen)
}
