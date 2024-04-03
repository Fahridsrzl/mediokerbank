package appconfig

const (
	LoanTransactionGroup                = "/transactions/loans"
	LoanTransactionCreate               = "/"
	LoanTransactionFindAll              = "/"
	LoanTransactionFindById             = "/:id"
	LoanTransactionFindByUserId         = "/users/:userId"
	LoanTransactionFindByUserIdAndTrxId = "users/:userId/:trxId"
	// LoanProductFindAll = "/"
	// LoanProductUpdate = "/:id"
	// LoanProductDelete = "/:id"
)
