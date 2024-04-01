package model

import "time"

type LoanProduct struct {
	Id                    string    `json:"id"`
	Name                  string    `json:"name"`
	MaxAmount             int       `json:"maxAmount"`
	MinInstallmentPeriod  int       `json:"minInstallmentPeriod"`
	MaxInstallmentPeriod  int       `json:"maxInstallmentPeriod"`
	InstallmentPeriodUnit string    `json:"installmentPeriodUnit"`
	AdminFee              int       `json:"adminFee"`
	MinCreditScore        int       `json:"minCreditScore"`
	MinMonthlyIncome      int       `json:"minMothlyIncome"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

type Loan struct {
	Id                string      `json:"id"`
	UserId            string      `json:"userId"`
	LoanProduct       LoanProduct `json:"loanProduct"`
	Amount            int         `json:"amount"`
	Interest          int         `json:"interest"`
	InstallmentAmount int         `json:"installmentAmount"`
	InstallmentPeriod int         `json:"installmentPeriod"`
	InstallmentUnit   string      `json:"installmentUnit"`
	PeriodLeft        int         `json:"periodLeft"`
	Status            string      `json:"status"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
}
