package events

import "time"

type OrderCreated struct {
	OrderID   string    `json:"order_id`
	Amount    float64   `json:"amount`
	CreatedAt time.Time `json:"created_at`
}
