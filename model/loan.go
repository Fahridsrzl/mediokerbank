package model

import "time"

type LoanProduct struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	MaxAmount        int       `json:"maxAmount"`
	PeriodUnit       string    `json:"periodUnit"`
	MinCreditScore   int       `json:"minCreditScore"`
	MinMonthlyIncome int       `json:"minMonthlyIncome"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
