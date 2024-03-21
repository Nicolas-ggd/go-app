package models

import (
	"context"
	"fmt"
	"log"
	"time"
)

type ICoin struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Balance   uint64    `json:"balance"`
	Fees      uint64    `json:"fees"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type ICoinCreate struct {
	Name string `json:"name"`
}

func (r *Repository) CreateBankAccount(account ICoinCreate, userId uint64) (*ICoin, error) {
	acc := &ICoin{
		Name:    account.Name,
		Balance: 0,
		Fees:    50,
		UserID:  userId,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stmt := `INSERT INTO user_bank_account (name, balance, fees, user_id)
	VALUES($1, $2, $3, $4)
	RETURNING id, name, balance, fees, user_id, created_at`

	_, err := r.DB.ExecContext(ctx, stmt, acc.Name, acc.Balance, acc.Fees, acc.UserID)
	if err != nil {
		log.Println("Can't create bank account")
		return nil, fmt.Errorf("can't create bank account with user ID: %v, got error: %v", userId, err)
	}

	return nil, nil
}
