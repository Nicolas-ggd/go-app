package models

type IBank struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Balance uint64 `json:"balance"`
	Fees    uint64 `json:"fees"`
	UserID  uint64 `json:"user_id"`
}
