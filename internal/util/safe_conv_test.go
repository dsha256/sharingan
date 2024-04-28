package util

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseFloat64_FormatFloat64(t *testing.T) {
	testCases := []struct {
		InStr  string
		OutF64 float64
	}{
		{
			InStr:  "0.0000000000001",
			OutF64: 0.0000000000001,
		},
		{
			InStr:  "12123.12312312312312313212321",
			OutF64: 12123.12312312312312313212321,
		},
		{
			InStr:  "0.0000000000001",
			OutF64: 0.0000000000001,
		},
		{
			InStr:  "0,0000000000001",
			OutF64: 0,
		},
		{
			InStr:  "okay",
			OutF64: 0,
		},
		{
			InStr:  FormatFloat64(math.MaxFloat64),
			OutF64: math.MaxFloat64,
		},
		{
			InStr:  FormatFloat64(math.SmallestNonzeroFloat64),
			OutF64: math.SmallestNonzeroFloat64,
		},
		{
			InStr:  FormatFloat64(math.MaxFloat64) + "0.1",
			OutF64: 0,
		},
		{
			InStr:  FormatFloat64(math.SmallestNonzeroFloat64) + "0.1",
			OutF64: 0,
		},
	}

	for _, tc := range testCases {
		got := ParseFloat64(tc.InStr)
		require.Equal(t, tc.OutF64, got)
	}
}
