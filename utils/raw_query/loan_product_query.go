package rawquery

const (
	GetLoanProductById    = `SELECT * FROM loan_products WHERE id = $1`
	GetAllLoanProducts    = `SELECT * FROM loan_products LIMIT $1 OFFSET $2`
	UpdateLoanProductById = `UPDATE loan_products SET name = $1, max_amount = $2, min_installment_period = $3, max_installment_period = $4, installment_period_unit = $5, admin_fee = $6, min_credit_score = $7, min_monthly_income = $8 WHERE id = $9`
	CreateLoanProduct     = `INSERT INTO loan_products (name, max_amount, min_installment_period, max_installment_period, installment_period_unit, admin_fee, min_credit_score, min_monthly_income) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;`
	DeleteLoanProduct = `DELETE FROM loan_products WHERE id = $1`
)
