CREATE TABLE loan_products (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	name VARCHAR(100),
	max_amount INT,
	period_unit VARCHAR(100),
	min_credit_score INT,
	min_monthly_income INT,
	created_at DATE,
	updated_at DATE
);

INSERT INTO loan_products (name, max_amount, period_unit, min_credit_score, min_monthly_income) VALUES
('classic', 100000, 'week', 5, 25000),
('silver', 1000000, 'month', 10, )