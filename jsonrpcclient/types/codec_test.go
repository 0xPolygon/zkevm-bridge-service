package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResponseMarshal(t *testing.T) {
	testCases := []struct {
		Name    string
		JSONRPC string
		ID      interface{}
		Result  interface{}
		Error   Error

		ExpectedJSON string
	}{
		{
			Name:    "Error is nil",
			JSONRPC: "2.0",
			ID:      1,
			Result: struct {
				A string `json:"A"`
			}{"A"},
			Error: nil,

			ExpectedJSON: "{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":{\"A\":\"A\"}}",
		},
		{
			Name:    "Result is nil and Error is not nil",
			JSONRPC: "2.0",
			ID:      1,
			Result:  nil,
			Error:   NewRPCError(123, "m"),

			ExpectedJSON: "{\"jsonrpc\":\"2.0\",\"id\":1,\"error\":{\"code\":123,\"message\":\"m\"}}",
		},
		{
			Name:    "Result is not nil and Error is not nil",
			JSONRPC: "2.0",
			ID:      1,
			Result: struct {
				A string `json:"A"`
			}{"A"},
			Error: NewRPCError(123, "m"),

			ExpectedJSON: "{\"jsonrpc\":\"2.0\",\"id\":1,\"error\":{\"code\":123,\"message\":\"m\"}}",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			req := Request{
				JSONRPC: testCase.JSONRPC,
				ID:      testCase.ID,
			}
			var result []byte
			if testCase.Result != nil {
				r, err := json.Marshal(testCase.Result)
				require.NoError(t, err)
				result = r
			}

			res := NewResponse(req, result, testCase.Error)
			bytes, err := json.Marshal(res)
			require.NoError(t, err)
			assert.Equal(t, testCase.ExpectedJSON, string(bytes))
		})
	}
}
