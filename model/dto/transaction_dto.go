package dto

type LoanTransactionRequestDto struct {
	UserId            string `json:"userId" binding:"required"`
	ProductId         string `json:"productId" binding:"required"`
	Amount            int    `json:"amount" binding:"required"`
	Purpose           string `json:"purpose" binding:"required"`
	InstallmentPeriod string `json:"installmentPeriod" binding:"required"`
}

type InstallmentTransactionRequestDto struct {
	UserId            string `json:"userId" binding:"required"`
	LoanId            string `json:"LoanId" binding:"required"`
	InstallmentAmount int    `json:"installmentAmount" binding:"required"`
	PaymentMethod     string `json:"paymentMethod" binding:"required"`
}
