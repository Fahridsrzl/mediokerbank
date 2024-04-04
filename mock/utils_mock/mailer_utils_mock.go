package utilsmock

import (
	"medioker-bank/model/dto"

	"github.com/stretchr/testify/mock"
)

type MailerUtilsMock struct {
	mock.Mock
}

func (m *MailerUtilsMock) SendEmail(payload dto.Mail) error {
	args := m.Called(payload)
	return args.Error(0)
}
