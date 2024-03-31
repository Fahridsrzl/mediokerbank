package repository

import (
	"database/sql"
	"medioker-bank/model"
	rawquery "medioker-bank/utils/raw_query"
	"time"
)

type LoanProductRepository interface {
	GetById(id string) (model.LoanProduct, error)
	GetAll() ([]model.LoanProduct, error)
	Create(payload model.LoanProduct) (model.LoanProduct, error)
	Update(id string, payload model.LoanProduct) (model.LoanProduct, error)
	Delete(id string) (model.LoanProduct, error)
}

type loanProductRepository struct {
	db *sql.DB
}

func (l *loanProductRepository) GetById(id string) (model.LoanProduct, error) {
	var loanProduct model.LoanProduct
	err := l.db.QueryRow(rawquery.GetLoanProductById, id).
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
	var loanProducts []model.LoanProduct
	rows, err := l.db.Query(rawquery.GetAllLoanProducts)
	if err != nil {
		return []model.LoanProduct{}, err
	}

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
	var createdLoanProduct model.LoanProduct
	err := l.db.QueryRow(rawquery.CreateLoanProduct,
		payload.Name, payload.MaxAmount, payload.PeriodUnit, payload.MinCreditScore, payload.MinMonthlyIncome, time.Now(), time.Now()).
		Scan(
			&createdLoanProduct.Id,
			&createdLoanProduct.Name,
			&createdLoanProduct.MaxAmount,
			&createdLoanProduct.PeriodUnit,
			&createdLoanProduct.MinCreditScore,
			&createdLoanProduct.MinMonthlyIncome,
			&createdLoanProduct.CreatedAt,
			&createdLoanProduct.UpdatedAt,
		)
	if err != nil {
		return model.LoanProduct{}, err
	}
	return createdLoanProduct, nil
}

func (l *loanProductRepository) Update(id string, payload model.LoanProduct) (model.LoanProduct, error) {
	var loanProduct model.LoanProduct
	err := l.db.QueryRow(rawquery.UpdateLoanProductById,
		payload.Name, payload.MaxAmount, payload.PeriodUnit, payload.MinCreditScore, payload.MinMonthlyIncome, time.Now() , time.Now(), id).Scan(&loanProduct.Id, &loanProduct.Name, &loanProduct.MaxAmount, &loanProduct.PeriodUnit, &loanProduct.MinCreditScore, &loanProduct.MinMonthlyIncome, &loanProduct.CreatedAt, &loanProduct.UpdatedAt)
	if err != nil {
		return model.LoanProduct{}, err
	}
	return loanProduct, nil
}

func (l *loanProductRepository) Delete(id string) (model.LoanProduct, error) {
	var loanProduct model.LoanProduct
	err := l.db.QueryRow(rawquery.DeleteLoanProduct, id).Scan(&loanProduct.Id, &loanProduct.Name, &loanProduct.MaxAmount, &loanProduct.PeriodUnit, &loanProduct.MinCreditScore, &loanProduct.MinMonthlyIncome, &loanProduct.CreatedAt, &loanProduct.UpdatedAt)
	if err != nil {
		return model.LoanProduct{}, err
	}
	return loanProduct, nil
}

func NewLoanProductRepository(db *sql.DB) LoanProductRepository {
	return &loanProductRepository{db: db}
}
