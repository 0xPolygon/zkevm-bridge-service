package types

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/0xPolygonHermez/zkevm-bridge-service/hex"
	"github.com/ethereum/go-ethereum/common"
)

// ArgUint64 helps to marshal uint64 values provided in the RPC requests
type ArgUint64 uint64

// MarshalText marshals into text
func (b ArgUint64) MarshalText() ([]byte, error) {
	buf := make([]byte, 2) //nolint:mnd
	copy(buf, `0x`)
	buf = strconv.AppendUint(buf, uint64(b), hex.Base)
	return buf, nil
}

// UnmarshalText unmarshals from text
func (b *ArgUint64) UnmarshalText(input []byte) error {
	str := strings.TrimPrefix(string(input), "0x")
	num, err := strconv.ParseUint(str, hex.Base, hex.BitSize64)
	if err != nil {
		return err
	}
	*b = ArgUint64(num)
	return nil
}

// Hex returns a hexadecimal representation
func (b ArgUint64) Hex() string {
	bb, _ := b.MarshalText()
	return string(bb)
}

// ArgUint64Ptr returns the pointer of the provided ArgUint64
func ArgUint64Ptr(a ArgUint64) *ArgUint64 {
	return &a
}

// ArgBytes helps to marshal byte array values provided in the RPC requests
type ArgBytes []byte

// MarshalText marshals into text
func (b ArgBytes) MarshalText() ([]byte, error) {
	return encodeToHex(b), nil
}

// UnmarshalText unmarshals from text
func (b *ArgBytes) UnmarshalText(input []byte) error {
	hh, err := decodeToHex(input)
	if err != nil {
		return nil
	}
	aux := make([]byte, len(hh))
	copy(aux, hh)
	*b = aux
	return nil
}

// Hex returns a hexadecimal representation
func (b ArgBytes) Hex() string {
	bb, _ := b.MarshalText()
	return string(bb)
}

// ArgBytesPtr helps to marshal byte array values provided in the RPC requests
func ArgBytesPtr(b []byte) *ArgBytes {
	bb := ArgBytes(b)

	return &bb
}

// ArgBig helps to marshal big number values provided in the RPC requests
type ArgBig big.Int

// UnmarshalText unmarshals an instance of ArgBig into an array of bytes
func (a *ArgBig) UnmarshalText(input []byte) error {
	buf, err := decodeToHex(input)
	if err != nil {
		return err
	}

	b := new(big.Int)
	b.SetBytes(buf)
	*a = ArgBig(*b)

	return nil
}

// MarshalText marshals an array of bytes into an instance of ArgBig
func (a ArgBig) MarshalText() ([]byte, error) {
	b := (*big.Int)(&a)

	return []byte("0x" + b.Text(hex.Base)), nil
}

// Hex returns a hexadecimal representation
func (b ArgBig) Hex() string {
	bb, _ := b.MarshalText()
	return string(bb)
}

func decodeToHex(b []byte) ([]byte, error) {
	str := string(b)
	str = strings.TrimPrefix(str, "0x")
	if len(str)%2 != 0 {
		str = "0" + str
	}
	return hex.DecodeString(str)
}

func encodeToHex(b []byte) []byte {
	str := hex.EncodeToString(b)
	if len(str)%2 != 0 {
		str = "0" + str
	}
	return []byte("0x" + str)
}

// ExitRoots structure
type ExitRoots struct {
	BlockNumber     ArgUint64   `json:"blockNumber"`
	Timestamp       ArgUint64   `json:"timestamp"`
	MainnetExitRoot common.Hash `json:"mainnetExitRoot"`
	RollupExitRoot  common.Hash `json:"rollupExitRoot"`
}
