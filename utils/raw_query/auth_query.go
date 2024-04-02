package rawquery

const (
	RegisterQueue                  = `INSERT INTO queue_register_users (username, email, password, v_code) VALUES ($1, $2, $3, $4)`
	RegisterUser                   = `INSERT INTO users (username, email, password, role, status, credit_score, balance) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`
	FindByVCode                    = `SELECT * FROM queue_register_users WHERE v_code = $1`
	DeleteQueueByVCode             = `DELETE FROM queue_register_users WHERE v_code = $1`
	FindByEmailUser                = `SELECT * FROM users WHERE email = $1`
	FindByUsernameUser             = `SELECT * FROM users WHERE username = $1`
	FindByEmailAndPasswordAdmin    = `SELECT * FROM admins WHERE email = $1 AND password = $2`
	FindByUsernameAndPasswordAdmin = `SELECT * FROM admins WHERE username = $1 AND password = $2`
)
