package appconfig

const (
	InstallmentGroup                 = "/transactions/installments"
	InstallmentCreate                = "/"
	InstallmentFindTrxById           = "/:id"
	InstallmentFindTrxMany           = "/"
	InstallmentFindTrxByUserId       = "/users/:userId"
	InstallmentFindTrxByUserAndTrxId = "/users/:userId/:trxId"
	InstallmentMidtransHook          = "/midtrans-hook"
)
