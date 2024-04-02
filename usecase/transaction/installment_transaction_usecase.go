package usecase

import (
	"errors"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	master "medioker-bank/repository/master"
	transaction "medioker-bank/repository/transaction"
	usecase "medioker-bank/usecase/master"
	"medioker-bank/utils/common"
)

type InstallmentTransactionUseCase interface {
	CreateTrx(payload dto.InstallmentTransactionRequestDto) (dto.InstallmentTransactionResponseDto, error)
	FindTrxById(id string) (model.InstallmentTransaction, error)
	FindTrxMany(payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error)
	FindTrxByUserId(userId string, payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error)
	FindTrxByUserIdAndTrxId(userId, trxId string) (model.InstallmentTransaction, error)
	UpdateTrxById(id string) error
}

type installmentTransactionUseCase struct {
	repo            transaction.InstallmentTransactionRepository
	loan            master.LoanRepository
	userUc          usecase.UserUseCase
	productUc       usecase.LoanProductUseCase
	midtransService common.MidtransService
}

func (i *installmentTransactionUseCase) CreateTrx(payload dto.InstallmentTransactionRequestDto) (dto.InstallmentTransactionResponseDto, error) {
	validPaymentMethod := []string{"medioker balance", "payment gateway"}
	validate := false
	for _, paymentMethod := range validPaymentMethod {
		if payload.PaymentMethod == paymentMethod {
			validate = true
		}
	}
	if !validate {
		return dto.InstallmentTransactionResponseDto{}, errors.New("allowed paymentMethod: 'medioker balance', 'payment gateway'")
	}
	loans, err := i.loan.FindByUserId(payload.UserId)
	if err != nil {
		return dto.InstallmentTransactionResponseDto{}, err
	}
	if len(loans) == 0 {
		return dto.InstallmentTransactionResponseDto{}, errors.New("loan id not found")
	}
	var loan model.Loan
	for _, item := range loans {
		if item.Id != payload.LoanId {
			return dto.InstallmentTransactionResponseDto{}, errors.New("loanId not found")
		} else {
			loan = item
		}
	}
	loanProduct, err := i.productUc.FindLoanProductById(loan.LoanProduct.Id)
	if err != nil {
		return dto.InstallmentTransactionResponseDto{}, err
	}
	loan.LoanProduct = loanProduct
	trxd := model.InstallmentTransactionDetail{
		Loan:              loan,
		InstallmentAmount: loan.InstallmentAmount,
		PaymentMethod:     payload.PaymentMethod,
	}
	trxReq := model.InstallmentTransaction{
		UserId:    payload.UserId,
		TrxDetail: trxd,
	}
	user, _, err := i.userUc.GetUserByID(trxReq.UserId)
	if err != nil {
		return dto.InstallmentTransactionResponseDto{}, err
	}
	var trxRes model.InstallmentTransaction
	var response dto.InstallmentTransactionResponseDto
	switch payload.PaymentMethod {
	case "medioker balance":
		trxReq.Status = "success"
		if user.Balance < trxReq.TrxDetail.Loan.InstallmentAmount {
			return dto.InstallmentTransactionResponseDto{}, errors.New("too low balance")
		}
		trxRes, err = i.repo.Create(trxReq)
		if err != nil {
			return dto.InstallmentTransactionResponseDto{}, err
		}
		err = i.loan.UpdatePeriod(trxReq.TrxDetail.Loan.Id)
		if err != nil {
			delErr := i.repo.DeleteById(trxRes.Id)
			if delErr != nil {
				return dto.InstallmentTransactionResponseDto{}, errors.New("[1]" + err.Error() + " [2]" + delErr.Error())
			}
			return dto.InstallmentTransactionResponseDto{}, err
		}
		_, err := i.userUc.UpdateUserBalance(user.ID, trxRes.TrxDetail.InstallmentAmount)
		if err != nil {
			return dto.InstallmentTransactionResponseDto{}, err
		}
		response.PaymentLink = "-"
		response.Message = "transaction success, your loan updated"
		response.Transaction = trxRes
	case "payment gateway":
		trxReq.Status = "pending"
		trxRes, err = i.repo.Create(trxReq)
		newMidTransPayment := dto.MidtransPayment{
			OrderId:     trxRes.Id,
			Amount:      trxRes.TrxDetail.InstallmentAmount,
			FirstName:   user.Profile.FirstName,
			LastName:    user.Profile.LastName,
			Email:       user.Email,
			PhoneNumber: user.Profile.PhoneNumber,
		}
		if err != nil {
			return dto.InstallmentTransactionResponseDto{}, errors.New("repo create: " + err.Error())
		}
		paymentLink, err := i.midtransService.CreatePayment(newMidTransPayment)
		if err != nil {
			delErr := i.repo.DeleteById(trxReq.Id)
			if delErr != nil {
				return dto.InstallmentTransactionResponseDto{}, errors.New("[1]" + err.Error() + " [2]" + delErr.Error())
			}
			return dto.InstallmentTransactionResponseDto{}, err
		}
		response.PaymentLink = paymentLink
		response.Message = "transaction success, check the link below to finish your payment"
	}
	trxRes.TrxDetail.Loan = loan
	response.Transaction = trxRes

	newLoans, err := i.loan.FindByUserId(payload.UserId)
	if err != nil {
		return dto.InstallmentTransactionResponseDto{}, err
	}
	var newLoan model.Loan
	for _, item := range newLoans {
		if item.Id != payload.LoanId {
			return dto.InstallmentTransactionResponseDto{}, errors.New("loanId not found")
		} else {
			newLoan = item
		}
	}
	if newLoan.PeriodLeft == 0 {
		err := i.loan.Delete(newLoan.Id)
		if err != nil {
			return dto.InstallmentTransactionResponseDto{}, err
		}
	}

	return response, nil
}

func (i *installmentTransactionUseCase) FindTrxById(id string) (model.InstallmentTransaction, error) {
	trx, err := i.repo.FindById(id)
	if err != nil {
		return model.InstallmentTransaction{}, errors.New("findTrx: " + err.Error())
	}
	loans, err := i.loan.FindByUserId(trx.UserId)
	if err != nil {
		return model.InstallmentTransaction{}, errors.New("find loans: " + err.Error())
	}
	var loan model.Loan
	for _, item := range loans {
		if item.Id == trx.TrxDetail.Loan.Id {
			product, err := i.productUc.FindLoanProductById(item.LoanProduct.Id)
			if err != nil {
				return model.InstallmentTransaction{}, errors.New("find loan product :" + err.Error())
			}
			item.LoanProduct = product
			loan = item
		}
	}
	trx.TrxDetail.Loan = loan
	return trx, err
}

func (i *installmentTransactionUseCase) FindTrxMany(payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error) {
	var trxs []model.InstallmentTransaction
	var err error
	if payload.TrxDate == "" {
		trxs, err = i.repo.FindAll()
		if err != nil {
			return []model.InstallmentTransaction{}, err
		}
	} else {
		trxs, err = i.repo.FindMany(payload)
		if err != nil {
			return []model.InstallmentTransaction{}, err
		}
	}
	var transactions []model.InstallmentTransaction
	for _, trx := range trxs {
		loans, err := i.loan.FindByUserId(trx.UserId)
		if err != nil {
			return []model.InstallmentTransaction{}, errors.New("find loans: " + err.Error())
		}
		var loan model.Loan
		for _, item := range loans {
			if item.Id == trx.TrxDetail.Loan.Id {
				product, err := i.productUc.FindLoanProductById(item.LoanProduct.Id)
				if err != nil {
					return []model.InstallmentTransaction{}, errors.New("find loan product :" + err.Error())
				}
				item.LoanProduct = product
				loan = item
			}
		}
		trx.TrxDetail.Loan = loan
		transactions = append(transactions, trx)
	}
	return transactions, nil
}

func (i *installmentTransactionUseCase) FindTrxByUserId(userId string, payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error) {
	trxs, err := i.repo.FindByUserId(userId, payload)
	if err != nil {
		return []model.InstallmentTransaction{}, err
	}
	var transactions []model.InstallmentTransaction
	for _, trx := range trxs {
		loans, err := i.loan.FindByUserId(trx.UserId)
		if err != nil {
			return []model.InstallmentTransaction{}, errors.New("find loans: " + err.Error())
		}
		var loan model.Loan
		for _, item := range loans {
			if item.Id == trx.TrxDetail.Loan.Id {
				product, err := i.productUc.FindLoanProductById(item.LoanProduct.Id)
				if err != nil {
					return []model.InstallmentTransaction{}, errors.New("find loan product :" + err.Error())
				}
				item.LoanProduct = product
				loan = item
			}
		}
		trx.TrxDetail.Loan = loan
		transactions = append(transactions, trx)
	}
	return transactions, nil
}

func (i *installmentTransactionUseCase) FindTrxByUserIdAndTrxId(userId, trxId string) (model.InstallmentTransaction, error) {
	trx, err := i.repo.FindByUserIdAndTrxId(userId, trxId)
	if err != nil {
		return model.InstallmentTransaction{}, err
	}
	loans, err := i.loan.FindByUserId(trx.UserId)
	if err != nil {
		return model.InstallmentTransaction{}, errors.New("find loans: " + err.Error())
	}
	var loan model.Loan
	for _, item := range loans {
		if item.Id == trx.TrxDetail.Loan.Id {
			product, err := i.productUc.FindLoanProductById(item.LoanProduct.Id)
			if err != nil {
				return model.InstallmentTransaction{}, errors.New("find loan product :" + err.Error())
			}
			item.LoanProduct = product
			loan = item
		}
	}
	trx.TrxDetail.Loan = loan
	return trx, nil
}

func (i *installmentTransactionUseCase) UpdateTrxById(id string) error {
	loanId, err := i.repo.UpdateById(id)
	if err != nil {
		return err
	}
	err = i.loan.UpdatePeriod(loanId)
	if err != nil {
		return err
	}
	return nil
}

func NewInstallmentTransactionUseCase(repo transaction.InstallmentTransactionRepository, loan master.LoanRepository, userUc usecase.UserUseCase, productUc usecase.LoanProductUseCase, midtransService common.MidtransService) InstallmentTransactionUseCase {
	return &installmentTransactionUseCase{repo: repo, loan: loan, userUc: userUc, productUc: productUc, midtransService: midtransService}
}
