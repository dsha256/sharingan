package dto

import "time"

const (
	BillingStatusOpen   BillStatus = "open"
	BillingStatusClosed BillStatus = "closed"
)

type BillStatus string

type Bill struct {
	Id        string    `json:"id"`
	Currency  string    `json:"currency"`
	Total     string    `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	ClosedAt  time.Time `json:"closed_at"`
	Items     []*Item   `json:"items"`
}

// Denilify checks if Items is nil and if so removes it with the appropriate value.
// USE CASE EXAMPLE: avoiding the JSON null (items: null => items: []) while marshaling
func (b *Bill) Denilify() *Bill {
	if b.Items == nil {
		b.Items = make([]*Item, 0)
	}

	return b
}
