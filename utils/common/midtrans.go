package common

import (
	"medioker-bank/config"
	"medioker-bank/model/dto"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService interface {
	CreatePayment(payload dto.MidtransPayment) (string, error)
}

type midtransService struct {
	cfg config.MidtransConfig
}

func (m *midtransService) CreatePayment(payload dto.MidtransPayment) (string, error) {
	s := snap.Client{}
	s.New(m.cfg.MidtransServerKey, midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payload.TrxId,
			GrossAmt: int64(payload.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: payload.FirstName,
			LName: payload.LastName,
			Email: payload.Email,
			Phone: payload.PhoneNumber,
		},
	}

	snapResp, err := s.CreateTransaction(req)
	if err != nil {
		return "", nil
	}
	return snapResp.RedirectURL, nil
}

func NewMidtransService(cfg config.MidtransConfig) MidtransService {
	return &midtransService{cfg: cfg}
}
