package member

import (
	"github.com/stretchr/testify/mock"
)

// MockRepository mocks Repository for testing purposes
type MockRepository struct {
	mock.Mock
}

// Add mock
func (m MockRepository) Add(mem Member) error {
	args := m.Called(mem)
	return args.Error(0)
}
