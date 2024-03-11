package mocks

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository simulates the behavior of the Repository interface for testing
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) QueryRow(query string, args ...interface{}) *sql.Row {
	params := m.Called(query, args)
	return params.Get(0).(*sql.Row)
}
