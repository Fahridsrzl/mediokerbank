package common

const (
	CreateStockProduct             = `INSERT INTO stock_products (company_name, rating, risk, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *`
	FindStockProductById           = `SELECT * FROM stock_products WHERE id = $1`
	FindAllStockProducts           = `SELECT * FROM stock_products`
	FindStockProductsByQueryRating = `SELECT * FROM stock_products WHERE rating = $1`
	FindStockProductsByQueryRisk   = `SELECT * FROM stock_products WHERE risk = $1`
	FindStockProductsByQueryBoth   = `SELECT * FROM stock_products WHERE rating = $1 AND risk = $2`
	UpdateStockProduct             = `UPDATE stock_products SET company_name = $1, rating = $2, risk = $3, updated_at = $4 WHERE id = $5 RETURNING *`
	DeleteStockProduct             = `DELETE FROM stock_products WHERE id = $1 RETURNING *`

	CreateUser                          = `INSERT INTO users (username,email,password,role,credit_score,created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id,username,email,password,role,credit_score,created_at,updated_at`
	GetUserById                         = `SELECT id,username,email,role,credit_score,created_at,updated_at FROM users WHERE id = $1`
	GetUserByUsernameOrEmailAndPassword = `SELECT id, username, email, role, credit_score, created_at, updated_at FROM users WHERE (username = $1 OR email = $2) AND password = $3`
	GetAllUsers                         = `SELECT id, username, email, role, credit_score, created_at, updated_at FROM users`
	UpdateUser                          = `UPDATE users SET username=$2, email=$3, password=$4, role=$5, credit_score=$6, updated_at=$7 WHERE id=$1`
	DeleteUser                          = `DELETE FROM users WHERE id=$1`
)
