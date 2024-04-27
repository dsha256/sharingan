package request

import (
	"errors"
	"math"

	"github.com/dsha256/sharingan/internal/util"
)

var (
	ErrInvalidName     = errors.New("name length must range from 1 to 50 chars")
	ErrInvalidQuantity = errors.New("quantity must be a positive integer")
	ErrInvalidPrice    = errors.New("price must be a positive float64")
)

type AddItem struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}

func (req *AddItem) Validate() error {
	if len(req.Name) < 1 || len(req.Name) > 50 {
		return ErrInvalidName
	}
	if req.Quantity == 0 {
		return ErrInvalidQuantity
	}

	p := util.ParseFloat64(req.Price)
	if p == 0 {
		return ErrInvalidPrice
	}

	if p < math.SmallestNonzeroFloat64 {
		return ErrInvalidPrice
	}

	req.Price = util.FormatFloat64(p)

	return nil
}
