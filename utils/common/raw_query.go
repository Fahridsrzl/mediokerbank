package common

const (
	GetUserByUsername = `SELECT id,email,username,password,role FROM users WHERE username = $1 OR email = $1`
	GetUserById       = `SELECT id,name,email,username,role,created_at,updated_at FROM users WHERE id = $1`
	CreateUser        = `INSERT INTO users (name,email,username,password,role,updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id,name,email,username,role,created_at,updated_at`

	GetProductById = `SELECT id,name,price,type,created_at,updated_at FROM products WHERE id = $1`

	GetCustomerById = `SELECT id,name,phone_number,address,created_at,updated_at FROM customers WHERE id = $1`

	GetBillById           = `SELECT b.id,b.bill_date, c.id,c.name,c.phone_number,c.address,c.created_at,c.updated_at,u.id,u.name,u.email,u.username,u.role,u.created_at,u.updated_at,b.created_at,b.updated_at FROM bills b JOIN customers c ON c.id = b.customer_id JOIN users u ON u.id = b.user_id WHERE `
	GetBillDetailByBillId = `SELECT	bd.id,p.id,p.name,p.price,p.type,p.created_at,p.updated_at,bd.qty,bd.price,bd.created_at,bd.updated_at FROM bill_details bd JOIN bills b ON b.id = bd.bill_id JOIN products p ON p.id = bd.product_id WHERE b.id = $1`
	CreateBill            = `INSERT INTO bills (bill_date,customer_id,user_id,updated_at) VALUES ($1,$2,$3,$4) RETURNING id,bill_date,created_at,updated_at`
	CreateBillDetail      = `INSERT INTO bill_details (bill_id,product_id,qty,price,updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING id,qty,price,created_at,updated_at`
)

// INSERT INTO products (name, price, type) VALUES ('cuci sprei',10000,'pcs');
// INSERT INTO customers (name, phone_number, address) VALUES ('Tika Yesi', '098786365', 'jakarta selatan');
