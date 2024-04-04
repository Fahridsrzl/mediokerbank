package model

import (
	"time"
)

type TopupTransaction struct {
	ID        string    `json:"id"`
	TrxDate   time.Time `json:"trxDate"`
	UserID    string    `json:"userId"`
	Amount    int       `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
