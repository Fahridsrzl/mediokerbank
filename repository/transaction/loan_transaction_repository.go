package repository

import (
	"database/sql"
	"fmt"
	"medioker-bank/model"
	rawquery "medioker-bank/utils/raw_query"
)

type LoanTransactionRepository interface {
	GetAll() ([]model.LoanTransaction, error)
	GetByID(id string) (model.LoanTransaction, error)
	GetByUserID(userId string) (model.LoanTransaction, error)
	GetByUserIdAndTrxId(userId, trxId string) ([]model.LoanTransaction, error)
	Create(payload model.LoanTransaction) (model.LoanTransaction, error)
}

type loanTransaction struct {
	db *sql.DB
}

func (l *loanTransaction) GetByUserIdAndTrxId(userId, trxId string) ([]model.LoanTransaction, error) {
	var loanTransactions []model.LoanTransaction

	query := rawquery.GetLoanTransactionByUserIdAndTrxId
	rows, err := l.db.Query(query, userId, trxId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		var loanTransactionDetail model.LoanTransactionDetail
		var loanProduct model.LoanProduct
		var loanTransaction model.LoanTransaction

		err := rows.Scan(
			&loanTransaction.Id,
			&loanTransaction.TrxDate,
			&user.ID,
			&user.Username,
			&user.Email,
			&loanTransaction.Status,
			&loanTransaction.CreatedAt,
			&loanTransaction.UpdatedAt,
			&loanTransactionDetail.Id,
			&loanProduct.Name,
			&loanProduct.MaxAmount,
			&loanProduct.MinInstallmentPeriod,
			&loanProduct.MaxInstallmentPeriod,
			&loanProduct.InstallmentPeriodUnit,
			&loanProduct.AdminFee,
			&loanProduct.MinCreditScore,
			&loanProduct.MinMonthlyIncome,
			&loanProduct.CreatedAt,
			&loanProduct.UpdatedAt,
			&loanTransactionDetail.Amount,
			&loanTransactionDetail.Purpose,
			&loanTransactionDetail.Interest,
			&loanTransactionDetail.InstallmentPeriod,
			&loanTransactionDetail.InstallmentUnit,
			&loanTransactionDetail.InstallmentAmount,
			&loanTransactionDetail.CreatedAt,
			&loanTransactionDetail.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		loanTransaction.User = user
		loanTransactionDetail.LoanProduct = loanProduct
		loanTransaction.LoanTransactionDetaills = append(loanTransaction.LoanTransactionDetaills, loanTransactionDetail)
		fmt.Println(loanTransaction)
		loanTransactions = append(loanTransactions, loanTransaction)
	}
	return loanTransactions, nil
}

func (l *loanTransaction) GetAll() ([]model.LoanTransaction, error) {
	var loanTransactions []model.LoanTransaction

	query := rawquery.GetAllLoanTransaction
	rows, err := l.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		var loanTransactionDetail model.LoanTransactionDetail
		var loanProduct model.LoanProduct
		var loanTransaction model.LoanTransaction

		err := rows.Scan(
			&loanTransaction.Id,
			&loanTransaction.TrxDate,
			&user.ID,
			&user.Username,
			&user.Email,
			&loanTransaction.Status,
			&loanTransaction.CreatedAt,
			&loanTransaction.UpdatedAt,
			&loanTransactionDetail.Id,
			&loanProduct.Name,
			&loanProduct.MaxAmount,
			&loanProduct.MinInstallmentPeriod,
			&loanProduct.MaxInstallmentPeriod,
			&loanProduct.InstallmentPeriodUnit,
			&loanProduct.AdminFee,
			&loanProduct.MinCreditScore,
			&loanProduct.MinMonthlyIncome,
			&loanProduct.CreatedAt,
			&loanProduct.UpdatedAt,
			&loanTransactionDetail.Amount,
			&loanTransactionDetail.Purpose,
			&loanTransactionDetail.Interest,
			&loanTransactionDetail.InstallmentPeriod,
			&loanTransactionDetail.InstallmentUnit,
			&loanTransactionDetail.InstallmentAmount,
			&loanTransactionDetail.CreatedAt,
			&loanTransactionDetail.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		loanTransaction.User = user
		loanTransactionDetail.LoanProduct = loanProduct
		loanTransaction.LoanTransactionDetaills = append(loanTransaction.LoanTransactionDetaills, loanTransactionDetail)
		fmt.Println(loanTransaction)
		loanTransactions = append(loanTransactions, loanTransaction)
	}
	return loanTransactions, nil
}

func (l *loanTransaction) GetByUserID(userId string) (model.LoanTransaction, error) {
	var loanTransaction model.LoanTransaction

	query := rawquery.GetLoanTransactionByUserId
	row := l.db.QueryRow(query, userId)

	var user model.User
	var loanTransactionDetail model.LoanTransactionDetail
	var loanProduct model.LoanProduct

	err := row.Scan(
		&loanTransaction.Id,
		&loanTransaction.TrxDate,
		&user.ID,
		&user.Username,
		&user.Email,
		&loanTransaction.Status,
		&loanTransaction.CreatedAt,
		&loanTransaction.UpdatedAt,
		&loanTransactionDetail.Id,
		&loanProduct.Name,
		&loanProduct.MaxAmount,
		&loanProduct.MinInstallmentPeriod,
		&loanProduct.MaxInstallmentPeriod,
		&loanProduct.InstallmentPeriodUnit,
		&loanProduct.AdminFee,
		&loanProduct.MinCreditScore,
		&loanProduct.MinMonthlyIncome,
		&loanProduct.CreatedAt,
		&loanProduct.UpdatedAt,
		&loanTransactionDetail.Amount,
		&loanTransactionDetail.Purpose,
		&loanTransactionDetail.Interest,
		&loanTransactionDetail.InstallmentPeriod,
		&loanTransactionDetail.InstallmentUnit,
		&loanTransactionDetail.InstallmentAmount,
		&loanTransactionDetail.CreatedAt,
		&loanTransactionDetail.UpdatedAt,
	)
	if err != nil {
		return model.LoanTransaction{}, err
	}

	loanTransaction.User = user
	loanTransactionDetail.LoanProduct = loanProduct
	loanTransaction.LoanTransactionDetaills = append(loanTransaction.LoanTransactionDetaills, loanTransactionDetail)

	return loanTransaction, nil
}

func (l *loanTransaction) GetByID(id string) (model.LoanTransaction, error) {
	var loanTransaction model.LoanTransaction

	query := rawquery.GetLoanTransactionById
	row := l.db.QueryRow(query, id)

	var user model.User
	var loanTransactionDetail model.LoanTransactionDetail
	var loanProduct model.LoanProduct

	err := row.Scan(
		&loanTransaction.Id,
		&loanTransaction.TrxDate,
		&user.ID,
		&user.Username,
		&user.Email,
		&loanTransaction.Status,
		&loanTransaction.CreatedAt,
		&loanTransaction.UpdatedAt,
		&loanTransactionDetail.Id,
		&loanProduct.Name,
		&loanProduct.MaxAmount,
		&loanProduct.MinInstallmentPeriod,
		&loanProduct.MaxInstallmentPeriod,
		&loanProduct.InstallmentPeriodUnit,
		&loanProduct.AdminFee,
		&loanProduct.MinCreditScore,
		&loanProduct.MinMonthlyIncome,
		&loanProduct.CreatedAt,
		&loanProduct.UpdatedAt,
		&loanTransactionDetail.Amount,
		&loanTransactionDetail.Purpose,
		&loanTransactionDetail.Interest,
		&loanTransactionDetail.InstallmentPeriod,
		&loanTransactionDetail.InstallmentUnit,
		&loanTransactionDetail.InstallmentAmount,
		&loanTransactionDetail.CreatedAt,
		&loanTransactionDetail.UpdatedAt,
	)
	if err != nil {
		return model.LoanTransaction{}, err
	}

	loanTransaction.User = user
	loanTransactionDetail.LoanProduct = loanProduct
	loanTransaction.LoanTransactionDetaills = append(loanTransaction.LoanTransactionDetaills, loanTransactionDetail)

	return loanTransaction, nil
}

func (l *loanTransaction) Create(payload model.LoanTransaction) (model.LoanTransaction, error) {
	tx, err := l.db.Begin()
	if err != nil {
		return model.LoanTransaction{}, err
	}

	var loanTransaction model.LoanTransaction
	err = tx.QueryRow(rawquery.CreateLoanTransaction, payload.User.ID, "active").Scan(
		&loanTransaction.Id,
		&loanTransaction.TrxDate,
		&loanTransaction.User.ID,
		&loanTransaction.Status,
		&loanTransaction.CreatedAt,
		&loanTransaction.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return model.LoanTransaction{}, err
	}

	var loanTransactionDetails []model.LoanTransactionDetail
	fmt.Println(payload.LoanTransactionDetaills)
	for _, v := range payload.LoanTransactionDetaills {
		var loanTransactionDetail model.LoanTransactionDetail
		loanProduct := model.LoanProduct{Id: v.LoanProduct.Id}
		err = tx.QueryRow(rawquery.CreateLoanTransactionDetail, loanProduct.Id, v.Amount, v.Purpose, v.Interest, v.InstallmentPeriod, v.InstallmentUnit, v.InstallmentAmount, loanTransaction.Id).Scan(
			&loanTransactionDetail.Id,
			&loanTransactionDetail.LoanProduct.Id,
			&loanTransactionDetail.Amount,
			&loanTransactionDetail.Purpose,
			&loanTransactionDetail.Interest,
			&loanTransactionDetail.InstallmentPeriod,
			&loanTransactionDetail.InstallmentUnit,
			&loanTransactionDetail.InstallmentAmount,
			&loanTransactionDetail.TrxId,
			&loanTransactionDetail.CreatedAt,
			&loanTransactionDetail.UpdatedAt,
		)
		if err != nil {
			tx.Rollback()
			return model.LoanTransaction{}, err
		}
		loanTransactionDetail.LoanProduct = v.LoanProduct
		loanTransactionDetails = append(loanTransactionDetails, loanTransactionDetail)
	}
	loanTransaction.User = payload.User
	loanTransaction.LoanTransactionDetaills = loanTransactionDetails
	if err := tx.Commit(); err != nil {
		return model.LoanTransaction{}, err
	}
	return loanTransaction, nil
}

func NewLoanTransactionRepository(db *sql.DB) LoanTransactionRepository {
	return &loanTransaction{db: db}
}
