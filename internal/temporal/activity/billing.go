package activity

import (
	"github.com/shopspring/decimal"
)

type MoneyAdditionOptions struct {
	Target   string
	Addition string
	Times    int
}

func MoneyAddition(args MoneyAdditionOptions) (string, error) {
	total, err := decimal.NewFromString(args.Target)
	if err != nil {
		return "", err
	}
	addition, err := decimal.NewFromString(args.Addition)
	if err != nil {
		return "", err
	}

	total = total.Add(addition.Mul(decimal.NewFromInt(int64(args.Times))))

	return total.String(), nil
}
