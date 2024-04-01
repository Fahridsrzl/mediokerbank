package rawquery

const (
	CreateTopup = `INSERT INTO topup_transactions (trx_date, user_id, amount, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, trx_date, user_id, amount, status, created_at, updated_at`
	GetTopUpByTopupId = `SELECT id, trx_date, user_id, amount, status, created_at, updated_at
	FROM topup_transactions
	WHERE id = $1`
	GetTopupByUserId = `SELECT id, trx_date, user_id, amount, status, created_at, updated_at
	FROM topup_transactions
	WHERE user_id = $1`
	GetAllTopUp = `SELECT id, trx_date, user_id, amount, status, created_at, updated_at
	FROM topup_transactions`
)
