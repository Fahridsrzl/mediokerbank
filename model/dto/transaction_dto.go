package dto

import "medioker-bank/model"

type InstallmentTransactionResponseDto struct {
	Message     string
	PaymentLink string
	Transaction model.InstallmentTransaction
}

type InstallmentTransactionSearchDto struct {
	TrxDate string `json:"trxDate"`
}

type LoanTransactionRequestDto struct {
	UserId                string                            `json:"userId" binding:"required"`
	LoanTransactionDetail []LoanTransactionDetailRequestDto `json:"loanTransactionDetail"`
}

type LoanTransactionDetailRequestDto struct {
	ProductId         string `json:"productId" binding:"required"`
	Amount            int    `json:"amount" binding:"required"`
	Purpose           string `json:"purpose" binding:"required"`
	InstallmentPeriod int    `json:"installmentPeriod" binding:"required"`
}

type InstallmentTransactionRequestDto struct {
	UserId        string `json:"userId" binding:"required"`
	LoanId        string `json:"loanId" binding:"required"`
	PaymentMethod string `json:"paymentMethod" binding:"required"`
}
