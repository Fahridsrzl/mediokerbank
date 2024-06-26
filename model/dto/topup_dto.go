package dto

import "time"

type TopupDto struct {
	UserID string `json:"userId" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}

type ResponseTopUp struct {
	ID        string    `json:"id"`
	TrxDate   time.Time `json:"trxDate"`
	UserID    string    `json:"userId"`
	Amount    int       `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
