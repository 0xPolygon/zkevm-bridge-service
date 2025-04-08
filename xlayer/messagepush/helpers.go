package messagepush

import (
	"encoding/json"
	"math/rand"

	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/pkg/errors"
)

func convertMsgToString(msg interface{}) (string, error) {
	var msgString string
	switch v := msg.(type) {
	case string:
		// If message is a string, just send it
		msgString = v
	default:
		// If message is an object, encode to json
		b, err := json.Marshal(msg)
		if err != nil {
			log.Errorf("msg cannot be encoded to json: msg[%v] err[%v]", msg, err)
			return "", errors.Wrap(err, "kafka produce: JSON marshal error")
		}
		msgString = string(b)
	}
	return msgString, nil
}

func GenerateTraceID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))] //nolint:gosec
	}
	return string(b)
}
