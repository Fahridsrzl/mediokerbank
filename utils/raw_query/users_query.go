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

	GetUserById = `SELECT
    u.id,
    u.username,
    u.email,
    u.password,
    u.role,
    u.status,
    u.credit_score,
    u.balance,
    u.created_at,
    u.updated_at,
    p.id,
    p.first_name,
    p.last_name,
    p.citizenship,
    p.national_id,
    p.birth_place,
    p.birth_date,
    p.gender,
    p.marital_status,
    p.occupation,
    p.monthly_income,
    p.phone_number,
    p.urgent_phone_number,
    p.photo,
    p.id_card,
    p.salary_slip,
    p.user_id,
    p.created_at,
    p.updated_at,
    a.id,
    a.address_line,
    a.city,
    a.province,
    a.postal_code,
    a.country,
    a.profile_id,
    a.created_at,
    a.updated_at
FROM
    users u
JOIN
    profiles p ON u.id = p.user_id
JOIN
    addresses a ON p.id = a.profile_id
WHERE
    u.id = $1`

	DeleteUser = `DELETE FROM users WHERE id = $1`

	GetAllUser = `SELECT id, username, email, role, status, credit_score, balance, created_at, updated_at
	FROM users LIMIT $1 OFFSET $2`

	UpdateBalance = `UPDATE users SET balance = balance - $1 WHERE id = $2 RETURNING balance`
)
