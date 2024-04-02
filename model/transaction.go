package model

import "time"

type InstallmentTransaction struct {
	Id        string                       `json:"id"`
	TrxDate   time.Time                    `json:"trxDate"`
	UserId    string                       `json:"userId"`
	Status    string                       `json:"status"`
	TrxDetail InstallmentTransactionDetail `json:"trxDetail"`
	CreatedAt time.Time                    `json:"createdAt"`
	UpdatedAt time.Time                    `json:"updatedAt"`
}

type InstallmentTransactionDetail struct {
	Id                string    `json:"id"`
	Loan              Loan      `json:"loan"`
	InstallmentAmount int       `json:"installmentAmount"`
	PaymentMethod     string    `json:"paymentMethod"`
	TrxId             string    `json:"trxId"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
