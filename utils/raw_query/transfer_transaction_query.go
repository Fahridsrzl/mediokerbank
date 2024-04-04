package rawquery

const (
	CreateTransfer = `INSERT INTO transfer_transactions (trx_date, sender_id, receiver_id, amount, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, trx_date, sender_id, receiver_id, amount, status, created_at, updated_at`
	GetTansferByTransferId = `SELECT id, trx_date, sender_id, receiver_id, amount, status, created_at, updated_at
	FROM transfer_transactions
	WHERE id = $1`
	GetTransferBySenderId = `SELECT id, trx_date, sender_id, receiver_id, amount, status, created_at, updated_at
	FROM transfer_transactions
	WHERE sender_id = $1`
	GetAllTransfer = `SELECT id, trx_date, sender_id, receiver_id, amount, status, created_at, updated_at
	FROM transfer_transactions LIMIT $1 OFFSET $2`
	UpdateSenderBalance   = `UPDATE users SET balance = balance - $1 WHERE id = $2`
	UpdateReceiverBalance = `UPDATE users SET balance = balance + $1 WHERE id = $2`
)
