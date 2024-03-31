CREATE DATABASE mediokerbank_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE admins (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	username VARCHAR(100) UNIQUE,
	email VARCHAR(100) UNIQUE,
	password VARCHAR(100),
	role VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE queue_register_users (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	username VARCHAR(100) UNIQUE,
	email VARCHAR(100) UNIQUE,
	password VARCHAR(100),
	v_code INT UNIQUE
);

CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	username VARCHAR(100) UNIQUE,
	email VARCHAR(100) UNIQUE,
	password VARCHAR(100),
	role VARCHAR(100),
	status VARCHAR(100),
	credit_score INT,
	balance INT,
	loan_active INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profiles (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	first_name VARCHAR(100),
	last_name VARCHAR(100),
	citizenship VARCHAR(100),
	national_id VARCHAR(100),
	birth_place VARCHAR(100),
	birth_date VARCHAR(100),
	gender VARCHAR(100),
	marital_status VARCHAR(100),
	occupation VARCHAR(100),
	monthly_income INT,
	phone_number VARCHAR(100) UNIQUE,
	urgent_phone_number VARCHAR(100),
	photo VARCHAR(100),
	id_card VARCHAR(100),
	salary_slip VARCHAR(100),
    user_id UUID UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE addresses (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	address_line VARCHAR(100),
	city VARCHAR(100),
	province VARCHAR(100),
	postal_code VARCHAR(100),
	country VARCHAR(100),
    profile_id UUID UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE
);

CREATE TABLE topup_transactions (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	trx_date VARCHAR(100),
	user_id UUID,
	amount INT,
	status VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE transfer_transactions (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	trx_date VARCHAR(100),
	sender_id UUID,
	receiver_id UUID,
	amount INT,
	status VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE loan_products (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	name VARCHAR(100),
	max_amount INT,
    min_installment_period INT,
    max_installment_period INT,
	installment_period_unit VARCHAR(100),
	admin_fee INT,
	min_credit_score INT,
	min_monthly_income INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE loan_transactions (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	trx_date VARCHAR(100),
	user_id UUID,
	status VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE loan_transaction_details (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	product_id UUID,
	amount INT,
	purpose VARCHAR(100),
	interest INT,
	installment_period INT,
	installment_unit VARCHAR(100),
	installment_amount INT,
	trx_id UUID UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES loan_products(id) ON DELETE CASCADE,
    FOREIGN KEY (trx_id) REFERENCES loan_transactions(id) ON DELETE CASCADE
);

CREATE TABLE loans (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id UUID,
	product_id UUID,
	amount INT,
	interest INT,
	installment_amount INT,
	installment_period INT,
	installment_unit VARCHAR(100),
	period_left INT,
	status VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (product_id) REFERENCES loan_products(id) ON DELETE CASCADE
);

CREATE TABLE installment_transactions (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	trx_date DATE,
	user_id UUID,
	status VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE installment_transaction_details (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	loan_id UUID UNIQUE,
	installment_amount INT,
	payment_method VARCHAR(100),
	trx_id UUID UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (trx_id) REFERENCES installment_transactions(id) ON DELETE CASCADE,
	FOREIGN KEY (loan_id) REFERENCES loans(id) ON DELETE CASCADE
);