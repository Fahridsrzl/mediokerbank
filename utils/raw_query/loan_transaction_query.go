package rawquery

const (
	CreateLoanTransaction = `INSERT INTO loan_transactions ( trx_date, user_id, status)
	VALUES ($1, $2, $3)
	RETURNING *;`
	CreateLoanTransactionDetail = `INSERT INTO loan_transaction_details (product_id, amount, purpose, interest, installment_period, installment_unit, installment_amount, trx_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING *;`
	// GetLoanTransactionById = `SELECT lt.id,lt.trx_date`

	GetLoanTransactionById = `
	SELECT 
		lt.id AS transaction_id, 
		lt.trx_date, 
		u.id AS user_id, 
		u.username AS username, 
		u.email AS user_email,
		lt.status, 
		lt.created_at AS transaction_created_at, 
		lt.updated_at AS transaction_updated_at,
		ltd.id AS detail_id, 
		lp.name AS product_name, 
		lp.max_amount, 
		lp.min_installment_period, 
		lp.max_installment_period, 
		lp.installment_period_unit, 
		lp.admin_fee, 
		lp.min_credit_score, 
		lp.min_monthly_income, 
		lp.created_at, 
		lp.updated_at, 
		ltd.amount, 
		ltd.purpose, 
		ltd.interest, 
		ltd.installment_period, 
		ltd.installment_unit, 
		ltd.installment_amount, 
		ltd.created_at AS detail_created_at,
		ltd.updated_at AS detail_updated_at
    FROM loan_transactions lt
    JOIN users u ON u.id = lt.user_id
    JOIN loan_transaction_details ltd ON ltd.trx_id = lt.id
    JOIN loan_products lp ON lp.id = ltd.product_id
	WHERE lt.id = $1;
`
	GetLoanTransactionByUserId = `
	SELECT
		lt.id AS transaction_id, 
		lt.trx_date, 
		u.id AS user_id, 
		u.username AS username, 
		u.email AS user_email,
		lt.status, 
		lt.created_at AS transaction_created_at, 
		lt.updated_at AS transaction_updated_at,
		ltd.id AS detail_id, 
		lp.name AS product_name, 
		lp.max_amount, 
		lp.min_installment_period, 
		lp.max_installment_period, 
		lp.installment_period_unit, 
		lp.admin_fee, 
		lp.min_credit_score, 
		lp.min_monthly_income, 
		lp.created_at, 
		lp.updated_at, 
		ltd.amount, 
		ltd.purpose, 
		ltd.interest, 
		ltd.installment_period, 
		ltd.installment_unit, 
		ltd.installment_amount, 
		ltd.created_at AS detail_created_at,
		ltd.updated_at AS detail_updated_at
	FROM loan_transactions lt
	JOIN users u ON u.id = lt.user_id
	JOIN loan_transaction_details ltd ON ltd.trx_id = lt.id
	JOIN loan_products lp ON lp.id = ltd.product_id
	WHERE u.id = $1;
	`
	GetAllLoanTransaction = `
	SELECT
		lt.id AS transaction_id, 
		lt.trx_date, 
		u.id AS user_id, 
		u.username AS username, 
		u.email AS user_email,
		lt.status, 
		lt.created_at AS transaction_created_at, 
		lt.updated_at AS transaction_updated_at,
		ltd.id AS detail_id, 
		lp.name AS product_name, 
		lp.max_amount, 
		lp.min_installment_period, 
		lp.max_installment_period, 
		lp.installment_period_unit, 
		lp.admin_fee, 
		lp.min_credit_score, 
		lp.min_monthly_income, 
		lp.created_at, 
		lp.updated_at, 
		ltd.amount, 
		ltd.purpose, 
		ltd.interest, 
		ltd.installment_period, 
		ltd.installment_unit, 
		ltd.installment_amount, 
		ltd.created_at AS detail_created_at,
		ltd.updated_at AS detail_updated_at
	FROM loan_transactions lt
	JOIN users u ON u.id = lt.user_id
	JOIN loan_transaction_details ltd ON ltd.trx_id = lt.id
	JOIN loan_products lp ON lp.id = ltd.product_id;
	`
)

// 	SelectAllLoanTransaction = `SELECT lt.id AS transaction_id, lt.trx_date, u.id AS user_id, u.name AS user_name, u.email AS user_email,
// 	ltd.id AS detail_id, lp.name AS product_name, ltd.amount, ltd.purpose, ltd.interest, ltd.installment_period,
// 	ltd.installment_unit, ltd.installment_amount, lt.status, lt.created_at AS transaction_created_at,
// 	ltd.created_at AS detail_created_at
// FROM loan_transactions lt
// JOIN users u ON u.id = lt.user_id
// JOIN loan_transaction_details ltd ON ltd.trx_id = lt.id
// JOIN loan_products lp ON lp.id = ltd.product_id
// ORDER BY lt.trx_date DESC;`

// )
