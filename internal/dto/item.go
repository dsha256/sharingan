package dto

import "time"

type Item struct {
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"added_at"`
}
