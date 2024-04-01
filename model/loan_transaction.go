package model

import "time"

type LoanTransaction struct {
	Id                      string                  `json:"id"`
	TrxDate                 time.Time               `json:"trxDate"`
	User                    User                    `json:"user"`
	Status                  string                  `json:"status"`
	LoanTransactionDetaills []LoanTransactionDetail `json:"loanTransactionDetails"`
	CreatedAt               time.Time               `json:"createdAt"`
	UpdatedAt               time.Time               `json:"updatedAt"`
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
