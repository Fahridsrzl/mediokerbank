package repository

import (
	"database/sql"
	"medioker-bank/model"
	rawquery "medioker-bank/utils/raw_query"
)

type LoanRepository interface {
	Create(payload model.Loan) (model.Loan, error)
	FindByUserId(userId string) ([]model.Loan, error)
	UpdatePeriod() error
	Delete(id string) error
}

type loanRepository struct {
	db *sql.DB
}

func (l *loanRepository) Create(payload model.Loan) (model.Loan, error) {
	var loan model.Loan
	err := l.db.QueryRow(rawquery.CreateLoan, payload).Scan(
		&loan.Id, &loan.UserId, &loan.LoanProduct.Id, &loan.Amount, &loan.Interest, &loan.InstallmentAmount, &loan.InstallmentPeriod, &loan.InstallmentUnit, &loan.PeriodLeft, &loan.Status, &loan.CreatedAt, &loan.UpdatedAt,
	)
	if err != nil {
		return model.Loan{}, err
	}
	return loan, nil
}

func (l *loanRepository) FindByUserId(userId string) ([]model.Loan, error) {
	var loans []model.Loan
	rows, err := l.db.Query(rawquery.FindLoanByUserId, userId)
	if err != nil {
		return []model.Loan{}, err
	}
	for rows.Next() {
		var loan model.Loan
		err := rows.Scan(
			&loan.Id, &loan.UserId, &loan.LoanProduct.Id, &loan.Amount, &loan.Interest, &loan.InstallmentAmount, &loan.InstallmentPeriod, &loan.InstallmentUnit, &loan.PeriodLeft, &loan.Status, &loan.CreatedAt, &loan.UpdatedAt,
		)
		if err != nil {
			return []model.Loan{}, err
		}
		loans = append(loans, loan)
	}
	return loans, nil
}

func (l *loanRepository) UpdatePeriod() error {
	_, err := l.db.Exec(rawquery.UpdateLoanPeriod)
	if err != nil {
		return err
	}
	return nil
}

func (l *loanRepository) Delete(id string) error {
	_, err := l.db.Exec(rawquery.DeleteLoan, id)
	if err != nil {
		return err
	}
	return nil
}

func NewLoanRepository(db *sql.DB) LoanRepository {
	return &loanRepository{db: db}
}
