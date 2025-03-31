package sqlitestorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertArrayUint64ToString(t *testing.T) {
	numbers := []uint64{1, 2, 4, 12, 78, 169, 200006712}
	result := convertArrayUint64ToString(numbers)
	assert.Equal(t, "1,2,4,12,78,169,200006712", result)
}

func TestConvertStringToArrayUint64(t *testing.T) {
	expectedResult := []uint64{1, 2, 4, 12, 78, 169, 200006712}
	input := "1,2,4,12,78,169,200006712"
	result, err := convertStringToArrayUint64(input)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

