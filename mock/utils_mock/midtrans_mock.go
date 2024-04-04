package utilsmock

import (
	"medioker-bank/model/dto"

	"github.com/stretchr/testify/mock"
)

type MidtrasnUtilsMock struct {
	mock.Mock
}

func (m *MidtrasnUtilsMock) CreatePayment(payload dto.MidtransPayment) (string, error) {
	args := m.Called(payload)
	return args.String(0), args.Error(1)
}
