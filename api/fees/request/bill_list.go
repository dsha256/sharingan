package request

import (
	"errors"

	"github.com/dsha256/sharingan/internal/dto"
)

var (
	ErrInvalidStatus = errors.New("invalid status, valid statuses: open|closed")
)

type ListBills struct {
	Status string `query:"status"`
}

func (req *ListBills) Validate() error {
	if dto.BillStatus(req.Status) != dto.BillingStatusOpen && dto.BillStatus(req.Status) != dto.BillingStatusClosed {
		return ErrInvalidStatus
	}

	return nil
}
