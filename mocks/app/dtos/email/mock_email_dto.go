package email

import (
	dto "notification/src/app/dtos/email"

	"github.com/stretchr/testify/mock"
)

type MockEmailDTO struct {
	mock.Mock
}

func NewMockEmailDTO() *MockEmailDTO {
	return &MockEmailDTO{}
}

var _ dto.EmailDTOInterface = &MockEmailDTO{}

func (m *MockEmailDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
