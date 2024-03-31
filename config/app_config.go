package config

const (
	UserSesion = "user"

	UserAdmin = "adfef570-8c30-43e0-bf5d-3a34ab908fcf"

	AuthGroup        = "/auth"
	AuthLoginUser    = "/admins/login"
	AuthRegisterUser = "/users/register"
	AuthVerifyUser   = "/users/register/verify"
	AuthLoginAdmin   = "/users/login"
	AuthRefreshToken = "/refresh-token"

	StockGroup    = "/stocks"
	StockCreate   = "/"
	StockFindById = "/:id"
	StockFindMany = "/"
	StockUpdate   = "/:id"
	StockDelete   = "/:id"
)
