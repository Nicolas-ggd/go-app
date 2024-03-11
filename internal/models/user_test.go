package models

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"websocket/internal/db"
	"websocket/internal/models/mocks"
)

func TestInsertUser(t *testing.T) {

	mockDB := &mocks.MockUserRepository{}

	// Create a test user form
	testUser := &UserForm{
		Name:     "test user",
		Email:    "test@gmail.com",
		Password: "password",
	}

	// Set up expectations for the mock's QueryRow method
	mockRow := sql.Row{}
	mockDB.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRow)

	// Create repository with mock database
	repo := &Repository{DB: db.DB}

	// Call the InsertUser method of the repository
	_, err := repo.InsertUser(testUser)

	// Check if there are any errors
	assert.NoError(t, err)

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)

}
