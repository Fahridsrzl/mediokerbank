package rawquery

const (
	CreateLoan       = `INSERT INTO loans (user_id, product_id, amount, interest, installment_amount, installment_period, installment_unit, period_left, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *`
	FindLoanByUserId = `SELECT * FROM loans WHERE user_id = $1`
	UpdateLoanPeriod = `UPDATE loans SET period_left = period_left - 1 WHERE id = $1`
	DeleteLoan       = `DELETE FROM loans WHERE id = $1`
)
