package repository

import (
	"database/sql"
	"medioker-bank/model"
	"medioker-bank/utils/common"
)
type LoanProductRepository interface {
	Get(id string) (model.LoanProduct, error)
	GetAll() ([]model.LoanProduct, error)
	Create(payload model.LoanProduct) (model.LoanProduct, error)
	Update(payload model.LoanProduct) error
	Delete(id string) error
}

type loanProductRepository struct {
	db *sql.DB
}

func (l *loanProductRepository) Get(id string) (model.LoanProduct, error) {
	var loanProduct model.LoanProduct
	err := l.db.QueryRow(common.GetLoanProductById, id).
		Scan(
			&loanProduct.Id,
			&loanProduct.Name,
			&loanProduct.MaxAmount,
			&loanProduct.PeriodUnit,
			&loanProduct.MinCreditScore,
			&loanProduct.MinMonthlyIncome,
			&loanProduct.CreatedAt,
			&loanProduct.UpdatedAt,
		)

	if err != nil {
		return model.LoanProduct{}, err
	}

	return loanProduct, nil
}

func (l *loanProductRepository) GetAll() ([]model.LoanProduct, error) {
	rows, err := l.db.Query(common.GetAllLoanProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var loanProducts []model.LoanProduct

	for rows.Next() {
		var loanProduct model.LoanProduct
		err := rows.Scan(
			&loanProduct.Id,
			&loanProduct.Name,
			&loanProduct.MaxAmount,
			&loanProduct.PeriodUnit,
			&loanProduct.MinCreditScore,
			&loanProduct.MinMonthlyIncome,
			&loanProduct.CreatedAt,
			&loanProduct.UpdatedAt,
		)
		if err != nil {
			return []model.LoanProduct{}, err
		}
		loanProducts = append(loanProducts, loanProduct)
	}

	return loanProducts, nil
}

func (l *loanProductRepository) Create(payload model.LoanProduct) (model.LoanProduct, error) {
	var createdProduct model.LoanProduct
	err := l.db.QueryRow(common.CreateLoanProduct,
		payload.Name, payload.MaxAmount, payload.PeriodUnit, payload.MinCreditScore, payload.MinMonthlyIncome).
		Scan(
			&createdProduct.Id,
			&createdProduct.Name,
			&createdProduct.MaxAmount,
			&createdProduct.PeriodUnit,
			&createdProduct.MinCreditScore,
			&createdProduct.MinMonthlyIncome,
			&createdProduct.CreatedAt,
			&createdProduct.UpdatedAt,
		)
	if err != nil {
		return model.LoanProduct{}, err
	}
	return createdProduct, nil
}

func (l *loanProductRepository) Update(payload model.LoanProduct) error {
	_, err := l.db.Exec(common.UpdateLoanProductById,
		payload.Id, payload.Name, payload.MaxAmount, payload.PeriodUnit, payload.MinCreditScore, payload.MinMonthlyIncome)
	if err != nil {
		return err
	}
	return nil
}

func (l *loanProductRepository) Delete(id string) error {
	_, err := l.db.Exec(common.DeleteLoanProduct, id)
	if err != nil {
		return err
	}
	return nil
}

func NewLoanProductRepository(db *sql.DB) LoanProductRepository {
	return &loanProductRepository{db:db}
}
