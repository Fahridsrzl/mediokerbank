package rawquery

const (
	CreateInstallment                 = `INSERT INTO installment_transactions (user_id, status) VALUES ($1, $2) RETURNING *`
	FindInstallmentById               = `SELECT * FROM installment_transactions WHERE id = $1`
	FindInstallmentsAll               = `SELECT * FROM installment_transactions`
	FindInstallmentsBySearch          = `SELECT * FROM installment_transactions WHERE trx_date = $1`
	FindInstallmentsByUserId          = `SELECT * FROM installment_transactions WHERE user_id = $1`
	FindInstallmentsByUserIdAndSearch = `SELECT * FROM installment_transactions WHERE user_id = $1 AND trx_date =$2`
	FindInstallmentByUserIdAndTrxId   = `SELECT * FROM installment_transactions WHERE user_id = $1 AND id = $1`
	UpdateInstallmentById             = `UPDATE installment_transactions SET status = $1 WHERE id = $1`
	DeleteInstallmentById             = `DELETE FROM installment_transactions WHERE id = $1`

	CreateInstallmentDetail   = `INSERT INTO installment_transactions (loan_id, installment_amount, payment_method, trx_id) VALUES ($1, $2, $3, $4) RETURNING *`
	FindInstallmentDetailById = `SELECT * FROM installment_transaction_details WHERE trx_id = $1`
)
