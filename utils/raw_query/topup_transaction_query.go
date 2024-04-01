package rawquery

const (
	CreateTopup = `INSERT INTO topup_transactions (trx_date, user_id, amount, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id, trx_date, user_id, amount, status, created_at, updated_at`
)
