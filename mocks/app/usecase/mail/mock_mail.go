package mail

import (
	dto "notification/src/app/dtos/email"
	usecase "notification/src/app/usecase/mail"

	"github.com/stretchr/testify/mock"
)

type MockMailUC struct {
	mock.Mock
}

func NewMockMailUC() *MockMailUC {
	return &MockMailUC{}
}

var _ usecase.MailUsecaseInterface = &MockMailUC{}

func (m *MockMailUC) MailNotif(data *dto.EmailReqDTO) error {
	args := m.Called(data)
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}

func (m *MockMailUC) MailSent(data *dto.EmailReqDTO) error {
	args := m.Called(data)
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
