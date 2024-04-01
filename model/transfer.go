package model

import (
	"time"
)

type TransferTransaction struct {
	ID         string    `json:"id"`
	TrxDate    time.Time `json:"trxDate"`
	SenderID   string    `json:"senderId"`
	ReceiverID string    `json:"receiverId"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
