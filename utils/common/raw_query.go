package common

const (
	GetLoanProductById = `SELECT id,name,max_amount,period_unit,min_credit_score,min_monthly_income,created_at,updated_at FROM loan_products WHERE id = $1`
	GetAllLoanProducts = `SELECT id,name,max_amount,period_unit,min_credit_score,min_monthly_income,created_at,updated_at FROM loan_products`
	UpdateLoanProductById = `UPDATE loan_products 
	SET name = $2, max_amount = $3, period_unit = $4, min_credit_score = $5, min_monthly_income = $6, updated_at = CURRENT_TIMESTAMP 
	WHERE id = $1`
	CreateLoanProduct = `INSERT INTO loan_products (name, max_amount, period_unit, min_credit_score, min_monthly_income, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
	RETURNING id`
	DeleteLoanProduct = `DELETE id,name,max_amount,period_unit,min_credit_score,min_monthly_income,created_at,updated_at FROM loan_products WHERE id = $1`
)


