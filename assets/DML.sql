INSERT INTO admins (username, email, password, role) VALUES
('admin', 'admin@medioker.com', 'admin123', 'admin');

INSERT INTO loan_products (name, max_amount, min_installment_period, max_installment_period, installment_period_unit, admin_fee, min_credit_score, min_monthly_income) VALUES
('poor', 300000, 3, 10, 'month', 5, 10, 103000),
('classic', 1000000, 5, 10, 'month', 5, 20, 210000),
('silver', 10000000, 12, 24, 'month', 5, 50, 940000),
('gold', 100000000, 12, 36, 'month', 5, 80, 9400000),
('platinum', 1000000000, 24, 60, 'month', 5, 100, 52000000);
