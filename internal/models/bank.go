package models

type IBankAccount interface {
	GetBalance(userId uint64) uint64
	Deposit(amount uint64)
	Withdraw(amount uint64) error
}
