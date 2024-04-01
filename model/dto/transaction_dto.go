package dto

type LoanTransactionRequestDto struct {
	UserId                string `json:"userId" binding:"required"`
	LoanTransactionDetail []LoanTransactionDetailRequestDto `json:"loanTransactionDetail"`
}

type LoanTransactionDetailRequestDto struct{
	ProductId             string `json:"productId" binding:"required"`
	Amount                int    `json:"amount" binding:"required"`
	Purpose               string `json:"purpose" binding:"required"`
	InstallmentPeriod     int `json:"installmentPeriod" binding:"required"`
}
