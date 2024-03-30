package common

const (
	CreateStockProduct             = `INSERT INTO stock_products (company_name, rating, risk, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *`
	FindStockProductById           = `SELECT * FROM stock_products WHERE id = $1`
	FindAllStockProducts           = `SELECT * FROM stock_products`
	FindStockProductsByQueryRating = `SELECT * FROM stock_products WHERE rating = $1`
	FindStockProductsByQueryRisk   = `SELECT * FROM stock_products WHERE risk = $1`
	FindStockProductsByQueryBoth   = `SELECT * FROM stock_products WHERE rating = $1 AND risk = $2`
	UpdateStockProduct             = `UPDATE stock_products SET company_name = $1, rating = $2, risk = $3, updated_at = $4 WHERE id = $5 RETURNING *`
	DeleteStockProduct             = `DELETE FROM stock_products WHERE id = $1 RETURNING *`

	GetLoanProductById = `SELECT id,name,max_amount,period_unit,min_credit_score,min_monthly_income,created_at,updated_at FROM loan_products WHERE id = $1`
	GetAllLoanProducts = `SELECT id,name,max_amount,period_unit,min_credit_score,min_monthly_income,created_at,updated_at FROM loan_products`
	UpdateLoanProductById = `UPDATE loan_products SET name = $1, max_amount = $2, period_unit = $3, min_credit_score = $4, min_monthly_income = $5, created_at = $6,updated_at = $7 WHERE id = $8`
	CreateLoanProduct = `INSERT INTO loan_products (name, max_amount, period_unit, min_credit_score, min_monthly_income, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id,name, max_amount, period_unit, min_credit_score, min_monthly_income, created_at, updated_at;`
	DeleteLoanProduct = `DELETE id,name,max_amount,period_unit,min_credit_score,min_monthly_income,created_at,updated_at FROM loan_products WHERE id = $1`
)


