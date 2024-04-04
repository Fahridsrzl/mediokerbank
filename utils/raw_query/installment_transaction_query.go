package rawquery

const (
	CreateInstallment                 = `INSERT INTO installment_transactions (user_id, status) VALUES ($1, $2) RETURNING *`
	FindInstallmentById               = `SELECT * FROM installment_transactions WHERE id = $1`
	FindInstallmentsAll               = `SELECT * FROM installment_transactions LIMIT $2 OFFSET $3`
	FindInstallmentsBySearch          = `SELECT * FROM installment_transactions WHERE trx_date = $1 LIMIT $2 OFFSET $3`
	FindInstallmentsByUserId          = `SELECT * FROM installment_transactions WHERE user_id = $1`
	FindInstallmentsByUserIdAndSearch = `SELECT * FROM installment_transactions WHERE user_id = $1 AND trx_date =$2`
	FindInstallmentByUserIdAndTrxId   = `SELECT * FROM installment_transactions WHERE user_id = $1 AND id = $2`
	UpdateInstallmentById             = `UPDATE installment_transactions SET status = $1 WHERE id = $2`
	DeleteInstallmentById             = `DELETE FROM installment_transactions WHERE id = $1`

	CreateInstallmentDetail   = `INSERT INTO installment_transaction_details (loan_id, installment_amount, payment_method, trx_id) VALUES ($1, $2, $3, $4) RETURNING *`
	FindInstallmentDetailById = `SELECT * FROM installment_transaction_details WHERE trx_id = $1`
	FindStatusByLoanId        = `SELECT t.status, t.id, t.created_at FROM installment_transaction_details AS td INNER JOIN installment_transactions AS t ON t.id = td.trx_id WHERE td.loan_id = $1 AND t.status = $2`
	SelectLoanIdOnTrxd        = `SELECT loan_id from installment_transaction_details WHERE trx_id = $1`
)
