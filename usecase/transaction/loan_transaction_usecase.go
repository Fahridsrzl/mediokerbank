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
	FindAllLoanTransaction(page, limit int) ([]dto.LoanTransactionResponseDto, error)
	FIndLoanTransactionByUserIdAndTrxId(userId, trxId string) (dto.LoanTransactionResponseDto, error)
	FindById(id string) (dto.LoanTransactionResponseDto, error)
	FindByUserId(userId string) ([]dto.LoanTransactionResponseDto, error)
	RegisterNewTransaction(payload dto.LoanTransactionRequestDto) (model.LoanTransaction, error)
}

type loanTransactionUseCase struct {
	repo      rTransaction.LoanTransactionRepository
	userUC    uMaster.UserUseCase
	productUC uMaster.LoanProductUseCase
	loanRepo  rMaster.LoanRepository
}

func (l *loanTransactionUseCase) FIndLoanTransactionByUserIdAndTrxId(userId, trxId string) (dto.LoanTransactionResponseDto, error) {
	loanTransaction, err := l.repo.GetByUserIdAndTrxId(userId, trxId)
	if err != nil {
		return dto.LoanTransactionResponseDto{}, err
	}
	return loanTransaction, nil
}

func (l *loanTransactionUseCase) FindAllLoanTransaction(page, limit int) ([]dto.LoanTransactionResponseDto, error) {
	loanTransaction, err := l.repo.GetAll(page, limit)
	if err != nil {
		return []dto.LoanTransactionResponseDto{}, err
	}
	return loanTransaction, nil
}

func (l *loanTransactionUseCase) FindByUserId(userId string) ([]dto.LoanTransactionResponseDto, error) {
	trx, err := l.repo.GetByUserID(userId)
	if err != nil {
		return []dto.LoanTransactionResponseDto{}, fmt.Errorf("user with ID %s not found", userId)
	}
	return trx, nil
}

func (l *loanTransactionUseCase) FindById(id string) (dto.LoanTransactionResponseDto, error) {
	trx, err := l.repo.GetByID(id)
	if err != nil {
		return dto.LoanTransactionResponseDto{}, fmt.Errorf(err.Error())
	}
	return trx, nil
}

func (l *loanTransactionUseCase) RegisterNewTransaction(payload dto.LoanTransactionRequestDto) (model.LoanTransaction, error) {
	user, _, err := l.userUC.GetUserByID(payload.UserId)
	if err != nil {
		return model.LoanTransaction{}, err
	}
	if user.Status != "verified" {
		return model.LoanTransaction{}, fmt.Errorf("user is not verified")
	}
	var loanTransactionDetails model.LoanTransactionDetail
	var loanTransactionDetail []model.LoanTransactionDetail
	for _, vTrx := range payload.LoanTransactionDetail {
		product, err := l.productUC.FindLoanProductById(vTrx.ProductId)
		if err != nil {
			return model.LoanTransaction{}, err
		}

		if vTrx.Amount < 100000 {
			return model.LoanTransaction{}, fmt.Errorf("min amount 100000")
		}
		if vTrx.Amount > product.MaxAmount {
			return model.LoanTransaction{}, fmt.Errorf("amount exceeds maximum allowed amount")
		}
		if vTrx.InstallmentPeriod > product.MaxInstallmentPeriod {
			return model.LoanTransaction{}, fmt.Errorf("installment period exceeds maximum installment period")
		}
		if vTrx.InstallmentPeriod < product.MinInstallmentPeriod {
			return model.LoanTransaction{}, fmt.Errorf("installment period is less than the minimum limit")
		}
		if user.CreditScore < product.MinCreditScore {
			return model.LoanTransaction{}, fmt.Errorf("your credit score is too low")
		}
		if user.Profile.MonthlyIncome < product.MinMonthlyIncome {
			return model.LoanTransaction{}, fmt.Errorf("your monthly income is too low")
		}
		interest := vTrx.InstallmentPeriod
		loanTransactionDetails = model.LoanTransactionDetail{LoanProduct: product, Amount: vTrx.Amount, InstallmentUnit: "month", InstallmentAmount: vTrx.Amount + vTrx.Amount*interest/100 + vTrx.Amount*product.AdminFee/100, Purpose: vTrx.Purpose, Interest: interest, InstallmentPeriod: vTrx.InstallmentPeriod}
		loanTransactionDetail = append(loanTransactionDetail, loanTransactionDetails)
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

	for _, v := range loanTransactionDetail {
		loanPayload := model.Loan{
			UserId:            trx.User.ID,
			LoanProduct:       v.LoanProduct,
			Amount:            v.Amount,
			Interest:          v.Interest,
			InstallmentAmount: v.InstallmentAmount,
			InstallmentPeriod: v.InstallmentPeriod,
			InstallmentUnit:   v.InstallmentUnit,
			PeriodLeft:        v.InstallmentPeriod,
			Status:            "active",
		}
		_, err = l.loanRepo.Create(loanPayload)
		if err != nil {
			return model.LoanTransaction{}, err
		}
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
