package gerror

import "errors"

var (
	// XLayer
	// ErrInternalErrorForRpcCall bridge web call service, when occur internal error(db error, redis error, network error...), return this error
	ErrInternalErrorForRpcCall = errors.New("your request could not be processed, please try again later")
)
