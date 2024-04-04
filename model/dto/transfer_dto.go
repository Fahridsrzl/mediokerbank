package dto

import "time"

type TransferDto struct {
	SenderID   string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
	Amount     int    `json:"amount"`
}

type ResponseTransfer struct {
	ID         string    `json:"id"`
	TrxDate    time.Time `json:"trxDate"`
	SenderID   string    `json:"senderId"`
	ReceiverID string    `json:"receiverId"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
