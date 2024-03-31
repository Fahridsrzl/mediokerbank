CREATE DATABASE mediokerbank_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE admins (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	username VARCHAR(100),
	email VARCHAR(100),
	password VARCHAR(100),
	role VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	username VARCHAR(100) UNIQUE,
	email VARCHAR(100) UNIQUE,
	password VARCHAR(100),
	role VARCHAR(100),
	credit_score INT,
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
	phone_number VARCHAR(100),
	urgent_phone_number VARCHAR(100),
	photo VARCHAR(100),
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

CREATE TABLE wallets (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	balance INT,
	loan_active INT,
	stock_active INT,
	user_id UUID,
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
	status INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE loan_products (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	name VARCHAR(100),
	max_amount INT,
	period_unit VARCHAR(100),
	min_credit_score INT,
	min_monthly_income INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE loan_transactions (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	trx_date VARCHAR(100),
	user_id UUID,
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
	trx_id UUID,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES loan_products(id) ON DELETE CASCADE,
    FOREIGN KEY (trx_id) REFERENCES loan_transactions(id) ON DELETE CASCADE
);

CREATE TABLE stock_products (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	company_name VARCHAR(100),
	rating INT,
	risk VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE stock_transactions (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	trx_date VARCHAR(100),
	user_id UUID,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE stock_transaction_details (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	product_id UUID,
	amount INT,
	purpose VARCHAR(100),
	trx_id UUID,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES stock_products(id) ON DELETE CASCADE,
    FOREIGN KEY (trx_id) REFERENCES stock_transactions(id) ON DELETE CASCADE
);

CREATE TABLE loans (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	wallet_id UUID,
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
	FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE CASCADE,
	FOREIGN KEY (product_id) REFERENCES loan_products(id) ON DELETE CASCADE
);

CREATE TABLE stocks (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	wallet_id UUID,
	product_id UUID,
	stock_price INT,
	status VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE CASCADE,
	FOREIGN KEY (product_id) REFERENCES stock_products(id) ON DELETE CASCADE
);

CREATE TABLE queue_loan_updates (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	trx_date DATE,
	user_id UUID,
	loan_id UUID UNIQUE,
	installment_amount INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (loan_id) REFERENCES loans(id) ON DELETE CASCADE
);