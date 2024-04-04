package usecase

import (
	"fmt"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	rMaster "medioker-bank/repository/master"
	rTransaction "medioker-bank/repository/transaction"
	uMaster "medioker-bank/usecase/master"
)

type LoanTransactionUseCase interface {
	FindAllLoanTransaction(page, limit int) ([]model.LoanTransaction, error)
	FIndLoanTransactionByUserIdAndTrxId(userId, trxId string) ([]model.LoanTransaction, error)
	FindById(id string) (model.LoanTransaction, error)
	FindByUserId(userId string) (model.LoanTransaction, error)
	RegisterNewTransaction(payload dto.LoanTransactionRequestDto) (model.LoanTransaction, error)
}

type loanTransactionUseCase struct {
	repo      rTransaction.LoanTransactionRepository
	userUC    uMaster.UserUseCase
	productUC uMaster.LoanProductUseCase
	loanRepo  rMaster.LoanRepository
}

func (l *loanTransactionUseCase) FIndLoanTransactionByUserIdAndTrxId(userId, trxId string) ([]model.LoanTransaction, error) {
	var loanTransaction []model.LoanTransaction
	var err error
	loanTransaction, err = l.repo.GetByUserIdAndTrxId(userId, trxId)
	if err != nil {
		return []model.LoanTransaction{}, err
	}
	return loanTransaction, nil
}

func (l *loanTransactionUseCase) FindAllLoanTransaction(page, limit int) ([]model.LoanTransaction, error) {
	var loanTransaction []model.LoanTransaction
	var err error
	loanTransaction, err = l.repo.GetAll(page, limit)
	if err != nil {
		return []model.LoanTransaction{}, err
	}
	return loanTransaction, nil
}

func (l *loanTransactionUseCase) FindByUserId(userId string) (model.LoanTransaction, error) {
	trx, err := l.repo.GetByUserID(userId)
	if err != nil {
		return model.LoanTransaction{}, fmt.Errorf("user with ID %s not found", userId)
	}
	return trx, nil
}

func (l *loanTransactionUseCase) FindById(id string) (model.LoanTransaction, error) {
	trx, err := l.repo.GetByID(id)
	if err != nil {
		return model.LoanTransaction{}, fmt.Errorf("transaction with ID %s not found", id)
	}
	return trx, nil
}

func (l *loanTransactionUseCase) RegisterNewTransaction(payload dto.LoanTransactionRequestDto) (model.LoanTransaction, error) {
	fmt.Println("bills log: ", payload)
	user, _, err := l.userUC.GetUserByID(payload.UserId)
	if err != nil {
		return model.LoanTransaction{}, err
	}
	var loanProduct model.LoanProduct
	var loanTransactionDetails model.LoanTransactionDetail
	var loanTransactionDetail []model.LoanTransactionDetail
	for _, vTrx := range payload.LoanTransactionDetail {
		product, err := l.productUC.FindLoanProductById(vTrx.ProductId)
		if err != nil {
			return model.LoanTransaction{}, err
		}
		loanProduct = product
		if vTrx.Amount >= product.MaxAmount {
			return model.LoanTransaction{}, fmt.Errorf("amount exceeds maximum allowed amount")
		}
		if vTrx.InstallmentPeriod >= product.MaxInstallmentPeriod {
			return model.LoanTransaction{}, fmt.Errorf("period unit exceeds maximum allowed period unit")
		}
		interest := vTrx.InstallmentPeriod
		loanTransactionDetails = model.LoanTransactionDetail{LoanProduct: product, Amount: vTrx.Amount, Purpose: vTrx.Purpose, Interest: interest, InstallmentPeriod: vTrx.InstallmentPeriod, InstallmentUnit: "month", InstallmentAmount: vTrx.Amount}
		loanTransactionDetail = append(loanTransactionDetail, model.LoanTransactionDetail{LoanProduct: product, Amount: vTrx.Amount, Purpose: vTrx.Purpose, Interest: interest, InstallmentPeriod: vTrx.InstallmentPeriod, InstallmentUnit: "month", InstallmentAmount: vTrx.Amount})
		fmt.Println(loanTransactionDetail)
	}
	newTransactionPayload := model.LoanTransaction{
		User:                    user,
		LoanTransactionDetaills: loanTransactionDetail,
	}

	trx, err := l.repo.Create(newTransactionPayload)
	if err != nil {
		return model.LoanTransaction{}, err
	}
	loanPayload := model.Loan{
		UserId:            trx.User.ID,
		LoanProduct:       loanProduct,
		Amount:            loanTransactionDetails.Amount,
		Interest:          loanTransactionDetails.Interest,
		InstallmentAmount: loanTransactionDetails.InstallmentAmount,
		InstallmentPeriod: loanTransactionDetails.InstallmentPeriod,
		InstallmentUnit:   loanTransactionDetails.InstallmentUnit,
		PeriodLeft:        loanTransactionDetails.InstallmentPeriod,
		Status:            "Succes",
	}
	_, err = l.loanRepo.Create(loanPayload)
	if err != nil {
		return model.LoanTransaction{}, err
	}

	return trx, nil
}

func NewLoanTransactionUseCase(
	repo rTransaction.LoanTransactionRepository,
	userUC uMaster.UserUseCase,
	productUC uMaster.LoanProductUseCase,
	loanRepo rMaster.LoanRepository,
) LoanTransactionUseCase {
	return &loanTransactionUseCase{
		repo:      repo,
		userUC:    userUC,
		productUC: productUC,
		loanRepo:  loanRepo,
	}
}
