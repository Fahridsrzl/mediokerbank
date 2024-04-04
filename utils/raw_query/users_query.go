package rawquery

const (
	UpdateUser = `UPDATE users SET status = $1 WHERE id = $2`

	CreateProfile = `INSERT INTO profiles (
		first_name, last_name, citizenship, national_id, birth_place, birth_date, gender,
		marital_status, occupation, monthly_income, phone_number, urgent_phone_number,
		photo, id_card, salary_slip, user_id, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) RETURNING id, first_name, last_name, citizenship, national_id, birth_place, birth_date, gender, marital_status, occupation, monthly_income, phone_number, urgent_phone_number, photo, id_card, salary_slip, user_id, created_at, updated_at`

	CreateAddress = `INSERT INTO addresses (
		address_line, city, province, postal_code, country, profile_id, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, address_line, city, province, postal_code, country, profile_id, created_at, updated_at`

	GetUserByStatus = `SELECT id,username,email,role, status, credit_score, balance,created_at,updated_at FROM users WHERE status = $1`

	GetUserById = `SELECT * FROM users WHERE id = $1`

	GetProfile = `SELECT * FROM profiles WHERE user_id = $1`

	GetAddress = `SELECT * FROM addresses WHERE profile_id = $1`

	DeleteUser = `DELETE FROM users WHERE id = $1`

	GetAllUser = `SELECT id, username, email, role, status, credit_score, balance, created_at, updated_at
	FROM users LIMIT $1 OFFSET $2`

	UpdateBalance = `UPDATE users SET balance = balance - $1 WHERE id = $2 RETURNING balance`
)
