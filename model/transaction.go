package model

import "time"

type LoanTransaction struct {
	Id        string                `json:"id"`
	TrxDate   time.Time             `json:"trxDate"`
	UserId    string                `json:"userId"`
	Status    string                `json:"status"`
	TrxDetail LoanTransactionDetail `json:"trxDetail"`
	CreatedAt time.Time             `json:"createdAt"`
	UpdatedAt time.Time             `json:"updatedAt"`
}

type LoanTransactionDetail struct {
	Id                string      `json:"id"`
	LoanProduct       LoanProduct `json:"loanProduct"`
	Amount            int         `json:"amount"`
	Purpose           string      `json:"purpose"`
	Interest          int         `json:"interest"`
	InstallmentPeriod int         `json:"installmentPeriod"`
	InstallmentUnit   string      `json:"installmentUnit"`
	InstallmentAmount int         `json:"installmentAmount"`
	TrxId             string      `json:"trxId"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
}

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
	PaymentMethod     int       `json:"paymentMethod"`
	TrxId             string    `json:"trxId"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
