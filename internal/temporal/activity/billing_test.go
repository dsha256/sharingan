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
				Target:   "10.1728312712816297126317231712938712937",
				Addition: "0.1728312712816297126317231712938712937",
				Times:    100,
			},
			Expected:     "27.4559583994446009758040403006810006637",
			RequireError: false,
		},
		{
			Args: MoneyAdditionOptions{
				Target:   "10.000000000000000000000000000000000000000007",
				Addition: "00.000000000000000000000000000000000000000001",
				Times:    3,
			},
			Expected:     "10.00000000000000000000000000000000000000001",
			RequireError: false,
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
		{
			Args: MoneyAdditionOptions{
				Target:   "10,0000000000000003",
				Addition: "1.00000000000000007",
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
