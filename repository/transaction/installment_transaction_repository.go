package repository

import (
	"database/sql"
	"errors"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	rawquery "medioker-bank/utils/raw_query"
)

type InstallmentTransactionRepository interface {
	Create(payload model.InstallmentTransaction) (model.InstallmentTransaction, error)
	FindById(id string) (model.InstallmentTransaction, error)
	FindAll() ([]model.InstallmentTransaction, error)
	FindMany(payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error)
	FindByUserId(userId string, payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error)
	FindByUserIdAndTrxId(userId, trxId string) (model.InstallmentTransaction, error)
	UpdateById(id string) (string, error)
	DeleteById(id string) error
}

type installmentTransactionRepository struct {
	db *sql.DB
}

func (i *installmentTransactionRepository) Create(payload model.InstallmentTransaction) (model.InstallmentTransaction, error) {
	var trx model.InstallmentTransaction
	tx, err := i.db.Begin()
	if err != nil {
		return model.InstallmentTransaction{}, err
	}
	err = tx.QueryRow(rawquery.CreateInstallment, payload.UserId, payload.Status).Scan(
		&trx.Id, &trx.TrxDate, &trx.UserId, &trx.Status, &trx.CreatedAt, &trx.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return model.InstallmentTransaction{}, errors.New("trxdetail:" + err.Error())
	}
	payloadTrxd := payload.TrxDetail
	err = tx.QueryRow(rawquery.CreateInstallmentDetail, payloadTrxd.Loan.Id, payloadTrxd.InstallmentAmount, payloadTrxd.PaymentMethod, trx.Id).Scan(
		&trx.TrxDetail.Id, &trx.TrxDetail.Loan.Id, &trx.TrxDetail.InstallmentAmount, &trx.TrxDetail.PaymentMethod, &trx.TrxDetail.TrxId, &trx.TrxDetail.CreatedAt, &trx.TrxDetail.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return model.InstallmentTransaction{}, errors.New("trxdetail:" + err.Error())
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return model.InstallmentTransaction{}, err
	}
	return trx, nil
}

func (i *installmentTransactionRepository) FindById(id string) (model.InstallmentTransaction, error) {
	var trx model.InstallmentTransaction
	err := i.db.QueryRow(rawquery.FindInstallmentById, id).Scan(
		&trx.Id, &trx.TrxDate, &trx.UserId, &trx.Status, &trx.CreatedAt, &trx.UpdatedAt,
	)
	if err != nil {
		return model.InstallmentTransaction{}, err
	}
	rows, err := i.db.Query(rawquery.FindInstallmentDetailById, trx.Id)
	if err != nil {
		return model.InstallmentTransaction{}, err
	}
	for rows.Next() {
		var trxd model.InstallmentTransactionDetail
		err := rows.Scan(
			&trxd.Id, &trxd.Loan.Id, &trxd.InstallmentAmount, &trxd.PaymentMethod, &trxd.TrxId, &trxd.CreatedAt, &trxd.UpdatedAt,
		)
		if err != nil {
			return model.InstallmentTransaction{}, err
		}
		trxd.Loan.UserId = trx.UserId
		trx.TrxDetail = trxd
	}
	return trx, nil
}

func (i *installmentTransactionRepository) FindAll() ([]model.InstallmentTransaction, error) {
	var trxs []model.InstallmentTransaction
	var rows *sql.Rows
	var err error
	rows, err = i.db.Query(rawquery.FindInstallmentsAll)
	if err != nil {
		return []model.InstallmentTransaction{}, err
	}
	for rows.Next() {
		var trx model.InstallmentTransaction
		err := rows.Scan(
			&trx.Id, &trx.TrxDate, &trx.UserId, &trx.Status, &trx.CreatedAt, &trx.UpdatedAt,
		)
		if err != nil {
			return []model.InstallmentTransaction{}, err
		}
		rows, err := i.db.Query(rawquery.FindInstallmentDetailById, trx.Id)
		if err != nil {
			return []model.InstallmentTransaction{}, err
		}
		for rows.Next() {
			var trxd model.InstallmentTransactionDetail
			err := rows.Scan(
				&trxd.Id, &trxd.Loan.Id, &trxd.InstallmentAmount, &trxd.PaymentMethod, &trxd.TrxId, &trxd.CreatedAt, &trxd.UpdatedAt,
			)
			if err != nil {
				return []model.InstallmentTransaction{}, err
			}
			trxd.Loan.UserId = trx.UserId
			trx.TrxDetail = trxd
		}
		trxs = append(trxs, trx)
	}
	return trxs, nil
}

func (i *installmentTransactionRepository) FindMany(payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error) {
	var trxs []model.InstallmentTransaction
	var rows *sql.Rows
	var err error
	rows, err = i.db.Query(rawquery.FindInstallmentsBySearch, payload.TrxDate)
	if err != nil {
		return []model.InstallmentTransaction{}, err
	}
	for rows.Next() {
		var trx model.InstallmentTransaction
		err := rows.Scan(
			&trx.Id, &trx.TrxDate, &trx.UserId, &trx.Status, &trx.CreatedAt, &trx.UpdatedAt,
		)
		if err != nil {
			return []model.InstallmentTransaction{}, err
		}
		rows, err := i.db.Query(rawquery.FindInstallmentDetailById, trx.Id)
		if err != nil {
			return []model.InstallmentTransaction{}, err
		}
		for rows.Next() {
			var trxd model.InstallmentTransactionDetail
			err := rows.Scan(
				&trxd.Id, &trxd.Loan.Id, &trxd.InstallmentAmount, &trxd.PaymentMethod, &trxd.TrxId, &trxd.CreatedAt, &trxd.UpdatedAt,
			)
			if err != nil {
				return []model.InstallmentTransaction{}, err
			}
			trxd.Loan.UserId = trx.UserId
			trx.TrxDetail = trxd
		}
		trxs = append(trxs, trx)
	}
	return trxs, nil
}

func (i *installmentTransactionRepository) FindByUserId(userId string, payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error) {
	var trxs []model.InstallmentTransaction
	var rows *sql.Rows
	var err error
	var stmt *sql.Stmt
	var args []any
	args = append(args, userId)
	if payload.TrxDate == "" {
		stmt, err = i.db.Prepare(rawquery.FindInstallmentsByUserId)
		if err != nil {
			return []model.InstallmentTransaction{}, err
		}
	} else {
		stmt, err = i.db.Prepare(rawquery.FindInstallmentsByUserIdAndSearch)
		if err != nil {
			return []model.InstallmentTransaction{}, err
		}
		args = append(args, payload.TrxDate)
	}
	rows, err = stmt.Query(args...)
	if err != nil {
		return []model.InstallmentTransaction{}, err
	}
	for rows.Next() {
		var trx model.InstallmentTransaction
		err := rows.Scan(
			&trx.Id, &trx.TrxDate, &trx.UserId, &trx.Status, &trx.CreatedAt, &trx.UpdatedAt,
		)
		if err != nil {
			return []model.InstallmentTransaction{}, err
		}
		rows, err := i.db.Query(rawquery.FindInstallmentDetailById, trx.Id)
		if err != nil {
			return []model.InstallmentTransaction{}, err
		}
		for rows.Next() {
			var trxd model.InstallmentTransactionDetail
			err := rows.Scan(
				&trxd.Id, &trxd.Loan.Id, &trxd.InstallmentAmount, &trxd.PaymentMethod, &trxd.TrxId, &trxd.CreatedAt, &trxd.UpdatedAt,
			)
			if err != nil {
				return []model.InstallmentTransaction{}, err
			}
			trxd.Loan.UserId = trx.UserId
			trx.TrxDetail = trxd
		}
		trxs = append(trxs, trx)
	}
	return trxs, nil
}

func (i *installmentTransactionRepository) FindByUserIdAndTrxId(userId, trxId string) (model.InstallmentTransaction, error) {
	var trx model.InstallmentTransaction
	err := i.db.QueryRow(rawquery.FindInstallmentByUserIdAndTrxId, userId, trxId).Scan(
		&trx.Id, &trx.TrxDate, &trx.UserId, &trx.Status, &trx.CreatedAt, &trx.UpdatedAt,
	)
	if err != nil {
		return model.InstallmentTransaction{}, err
	}
	rows, err := i.db.Query(rawquery.FindInstallmentDetailById, trx.Id)
	if err != nil {
		return model.InstallmentTransaction{}, err
	}
	for rows.Next() {
		var trxd model.InstallmentTransactionDetail
		err := rows.Scan(
			&trxd.Id, &trxd.Loan.Id, &trxd.InstallmentAmount, &trxd.PaymentMethod, &trxd.TrxId, &trxd.CreatedAt, &trxd.UpdatedAt,
		)
		if err != nil {
			return model.InstallmentTransaction{}, err
		}
		trxd.Loan.UserId = trx.UserId
		trx.TrxDetail = trxd
	}
	return trx, nil
}

func (i *installmentTransactionRepository) UpdateById(id string) (string, error) {
	_, err := i.db.Exec(rawquery.UpdateInstallmentById, "success", id)
	if err != nil {
		return "", err
	}
	var loanId string
	err = i.db.QueryRow(rawquery.SelectLoanIdOnTrxd, id).Scan(&loanId)
	if err != nil {
		return "", err
	}
	return loanId, nil
}

func (i *installmentTransactionRepository) DeleteById(id string) error {
	_, err := i.db.Exec(rawquery.DeleteInstallmentById, id)
	if err != nil {
		return err
	}
	return nil
}

func NewInstallmentTransactionRepository(db *sql.DB) InstallmentTransactionRepository {
	return &installmentTransactionRepository{db: db}
}
