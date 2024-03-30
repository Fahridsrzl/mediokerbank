package config

const (
	UserSesion = "user"

	UserAdmin = "adfef570-8c30-43e0-bf5d-3a34ab908fcf"

	AuthGroup        = "/auth"
	AuthRegister     = "/register"
	AuthLogin        = "/login"
	AuthRefreshToken = "/refresh-token"

	StockGroup    = "/stocks"
	StockCreate   = "/"
	StockFindById = "/:id"
	StockFindMany = "/"
	StockUpdate   = "/:id"
	StockDelete   = "/:id"
	
	LoanProductGroup = "/loan-product"
	LoanProductCreate  = "/"
	LoanProductFindByid = "/:id"
	LoanProductFindAll = "/"
	LoanProductUpdate = "/:id"
	LoanProductDelete = "/:id"
)
