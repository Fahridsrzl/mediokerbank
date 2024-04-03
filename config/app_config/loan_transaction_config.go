package appconfig

const (
	LoanTransactionGroup                = "/transactions/loans"
	LoanTransactionCreate               = "/"
	LoanTransactionFindAll              = "/"
	LoanTransactionFindById             = "/:Id"
	LoanTransactionFindByUserId         = "/users/:userId"
	LoanTransactionFindByUserIdAndTrxId = "users/:userId/:trxId"
	// LoanProductFindAll = "/"
	// LoanProductUpdate = "/:id"
	// LoanProductDelete = "/:id"
)
