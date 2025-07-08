package types

import (
	"encoding/json"
)

const (
	// Earliest contains the string to represent the earliest block known.
	Earliest = "earliest"
	// Latest contains the string to represent the latest block known.
	Latest = "latest"
	// Pending contains the string to represent the pending block known.
	Pending = "pending"
	// Safe contains the string to represent the last virtualized block known.
	Safe = "safe"
	// Finalized contains the string to represent the last verified block known.
	Finalized = "finalized"

	// EIP-1898: https://eips.ethereum.org/EIPS/eip-1898 //

	// BlockNumberKey is the key for the block number for EIP-1898
	BlockNumberKey = "blockNumber"
	// BlockHashKey is the key for the block hash for EIP-1898
	BlockHashKey = "blockHash"
	// RequireCanonicalKey is the key for the require canonical for EIP-1898
	RequireCanonicalKey = "requireCanonical"
)

// Request is a jsonrpc request
type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

// Response is a jsonrpc  success response
type Response struct {
	JSONRPC string
	ID      interface{}
	Result  json.RawMessage
	Error   *ErrorObject
}

// ErrorObject is a jsonrpc error
type ErrorObject struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    *ArgBytes `json:"data,omitempty"`
}

// RPCError returns an instance of RPCError from the
// data available in the ErrorObject instance
func (e *ErrorObject) RPCError() RPCError {
	var data []byte
	if e.Data != nil {
		data = *e.Data
	}
	rpcError := NewRPCErrorWithData(e.Code, e.Message, data)
	return *rpcError
}

// NewResponse returns Success/Error response object
func NewResponse(req Request, reply []byte, err Error) Response {
	var result json.RawMessage
	if reply != nil {
		result = reply
	}

	var errorObj *ErrorObject
	if err != nil {
		errorObj = &ErrorObject{
			Code:    err.ErrorCode(),
			Message: err.Error(),
		}
		if err.ErrorData() != nil {
			errorObj.Data = ArgBytesPtr(err.ErrorData())
		}
	}

	return Response{
		JSONRPC: req.JSONRPC,
		ID:      req.ID,
		Result:  result,
		Error:   errorObj,
	}
}

// MarshalJSON customizes the JSON representation of the response.
func (r Response) MarshalJSON() ([]byte, error) {
	if r.Error != nil {
		return json.Marshal(struct {
			JSONRPC string       `json:"jsonrpc"`
			ID      interface{}  `json:"id"`
			Error   *ErrorObject `json:"error"`
		}{
			JSONRPC: r.JSONRPC,
			ID:      r.ID,
			Error:   r.Error,
		})
	}

	return json.Marshal(struct {
		JSONRPC string          `json:"jsonrpc"`
		ID      interface{}     `json:"id"`
		Result  json.RawMessage `json:"result"`
	}{
		JSONRPC: r.JSONRPC,
		ID:      r.ID,
		Result:  r.Result,
	})
}

// Bytes return the serialized response
func (s Response) Bytes() ([]byte, error) {
	return json.Marshal(s)
}
