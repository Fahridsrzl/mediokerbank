package dto

import "medioker-bank/model"

type LoanTransactionRequestDto struct {
	UserId            string `json:"userId" binding:"required"`
	ProductId         string `json:"productId" binding:"required"`
	Amount            int    `json:"amount" binding:"required"`
	Purpose           string `json:"purpose" binding:"required"`
	InstallmentPeriod string `json:"installmentPeriod" binding:"required"`
}

type InstallmentTransactionRequestDto struct {
	UserId        string `json:"userId" binding:"required"`
	LoanId        string `json:"loanId" binding:"required"`
	PaymentMethod string `json:"paymentMethod" binding:"required"`
}

type InstallmentTransactionResponseDto struct {
	Message     string
	PaymentLink string
	Transaction model.InstallmentTransaction
}

type InstallmentTransactionSearchDto struct {
	TrxDate string `json:"trxDate"`
}
