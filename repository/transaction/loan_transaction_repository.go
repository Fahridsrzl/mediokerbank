package repository

import (
	"database/sql"
	"fmt"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	rawquery "medioker-bank/utils/raw_query"
)

type LoanTransactionRepository interface {
	GetAll(page, limit int) ([]dto.LoanTransactionResponseDto, error)
	GetByID(id string) (dto.LoanTransactionResponseDto, error)
	GetByUserID(userId string) ([]dto.LoanTransactionResponseDto, error)
	GetByUserIdAndTrxId(userId, trxId string) (dto.LoanTransactionResponseDto, error)
	Create(payload model.LoanTransaction) (model.LoanTransaction, error)
}

type loanTransaction struct {
	db *sql.DB
}

func (l *loanTransaction) GetByUserIdAndTrxId(userId, trxId string) (dto.LoanTransactionResponseDto, error) {
	var loanTransactions []model.LoanTransaction

	query := rawquery.GetLoanTransactionByUserIdAndTrxId
	row := l.db.QueryRow(query, userId, trxId)

	var user model.User
	var loanTransactionDetail model.LoanTransactionDetail
	var loanProduct model.LoanProduct
	var loanTransaction model.LoanTransaction

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
		return dto.LoanTransactionResponseDto{}, err
	}

	loanTransaction.User = user
	loanTransactionDetail.LoanProduct = loanProduct
	loanTransaction.LoanTransactionDetaills = append(loanTransaction.LoanTransactionDetaills, loanTransactionDetail)
	fmt.Println(loanTransaction)
	loanTransactions = append(loanTransactions, loanTransaction)

	result := dto.LoanTransactionResponseDto{
		Id:                      loanTransaction.Id,
		TrxDate:                 loanTransaction.TrxDate,
		UserId:                  loanTransaction.User.ID,
		LoanTransactionDetaills: loanTransaction.LoanTransactionDetaills,
		Status:                  loanTransaction.Status,
		CreatedAt:               loanTransaction.CreatedAt,
		UpdatedAt:               loanTransaction.UpdatedAt,
	}
	return result, nil
}

func (l *loanTransaction) GetAll(page, limit int) ([]dto.LoanTransactionResponseDto, error) {
	var results []dto.LoanTransactionResponseDto

	offset := (page - 1) * limit
	query := rawquery.GetAllLoanTransaction
	rows, err := l.db.Query(query, limit, offset)
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

		result := dto.LoanTransactionResponseDto{
			Id:                      loanTransaction.Id,
			TrxDate:                 loanTransaction.TrxDate,
			UserId:                  loanTransaction.User.ID,
			LoanTransactionDetaills: loanTransaction.LoanTransactionDetaills,
			Status:                  loanTransaction.Status,
			CreatedAt:               loanTransaction.CreatedAt,
			UpdatedAt:               loanTransaction.UpdatedAt,
		}
		results = append(results, result)
	}
	return results, nil
}

func (l *loanTransaction) GetByUserID(userId string) ([]dto.LoanTransactionResponseDto, error) {
	var results []dto.LoanTransactionResponseDto

	query := rawquery.GetLoanTransactionByUserId
	rows, err := l.db.Query(query, userId)
	if err != nil {
		return []dto.LoanTransactionResponseDto{}, err
	}

	for rows.Next() {
		var loanTransaction model.LoanTransaction
		var user model.User
		var loanTransactionDetail model.LoanTransactionDetail
		var loanProduct model.LoanProduct

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
			&loanProduct.Id,
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
			return []dto.LoanTransactionResponseDto{}, err
		}

		loanTransaction.User = user
		loanTransactionDetail.LoanProduct = loanProduct
		loanTransactionDetail.TrxId = loanTransaction.Id
		loanTransaction.LoanTransactionDetaills = append(loanTransaction.LoanTransactionDetaills, loanTransactionDetail)

		result := dto.LoanTransactionResponseDto{
			Id:                      loanTransaction.Id,
			TrxDate:                 loanTransaction.TrxDate,
			UserId:                  loanTransaction.User.ID,
			LoanTransactionDetaills: loanTransaction.LoanTransactionDetaills,
			Status:                  loanTransaction.Status,
			CreatedAt:               loanTransaction.CreatedAt,
			UpdatedAt:               loanTransaction.UpdatedAt,
		}
		results = append(results, result)
	}

	return results, nil
}

func (l *loanTransaction) GetByID(id string) (dto.LoanTransactionResponseDto, error) {
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
		&loanProduct.Id,
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
		return dto.LoanTransactionResponseDto{}, err
	}

	loanTransaction.User = user
	loanTransactionDetail.LoanProduct = loanProduct
	loanTransactionDetail.TrxId = loanTransaction.Id
	loanTransaction.LoanTransactionDetaills = append(loanTransaction.LoanTransactionDetaills, loanTransactionDetail)

	result := dto.LoanTransactionResponseDto{
		Id:                      loanTransaction.Id,
		TrxDate:                 loanTransaction.TrxDate,
		UserId:                  loanTransaction.User.ID,
		LoanTransactionDetaills: loanTransaction.LoanTransactionDetaills,
		Status:                  loanTransaction.Status,
		CreatedAt:               loanTransaction.CreatedAt,
		UpdatedAt:               loanTransaction.UpdatedAt,
	}

	return result, nil
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
