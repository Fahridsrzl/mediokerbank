package rawquery

const (
	CreateUser                          = `INSERT INTO users (username,email,password,role,credit_score,created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id,username,email,password,role,credit_score,created_at,updated_at`
	GetUserById                         = `SELECT id,username,email,role,credit_score,created_at,updated_at FROM users WHERE id = $1`
	GetUserByUsernameOrEmailAndPassword = `SELECT id, username, email, role, credit_score, created_at, updated_at FROM users WHERE (username = $1 OR email = $2) AND password = $3`
	GetAllUsers                         = `SELECT id, username, email, role, credit_score, created_at, updated_at FROM users`
	UpdateUser                          = `UPDATE users SET username=$2, email=$3, password=$4, role=$5, credit_score=$6, updated_at=$7 WHERE id=$1`
	DeleteUser                          = `DELETE FROM users WHERE id=$1`
)


