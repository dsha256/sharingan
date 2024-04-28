package activity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MoneyAddition(t *testing.T) {
	testCases := []struct {
		Args         MoneyAdditionOptions
		Expected     string
		RequireError bool
	}{
		{
			Args: MoneyAdditionOptions{
				Target:   "0.0",
				Addition: "1.0",
				Times:    2,
			},
			Expected:     "2",
			RequireError: false,
		},
		{
			Args: MoneyAdditionOptions{
				Target:   "10.0000000000000003",
				Addition: "1.00000000000000007",
				Times:    2,
			},
			Expected:     "12.00000000000000044",
			RequireError: false,
		},
		{
			Args: MoneyAdditionOptions{
				Target:   "10.0000000000000003",
				Addition: "1.00000000000000007",
				Times:    2,
			},
			Expected:     "12.00000000000000044",
			RequireError: false,
		},
		{
			Args: MoneyAdditionOptions{
				Target:   "10,0000000000000003",
				Addition: "1.00000000000000007",
				Times:    2,
			},
			Expected:     "12.00000000000000044",
			RequireError: true,
		},
		{
			Args: MoneyAdditionOptions{
				Target:   "10.0000000000000003",
				Addition: "1,00000000000000007",
				Times:    2,
			},
			Expected:     "12.00000000000000044",
			RequireError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Expected, func(t *testing.T) {
			got, err := MoneyAddition(testCase.Args)
			if testCase.RequireError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.Expected, got)
			}
		})
	}
}
