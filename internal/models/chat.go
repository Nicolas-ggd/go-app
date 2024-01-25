package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"websocket/internal/db"
)

type ChatModelInterface interface {
	InsertMessage(c *Chat) error
}

type Chat struct {
	ID         uint64         `json:"id" gorm:"primaryKey"`
	SenderID   uint64         `json:"sender_id" binding:"required"`
	ReceiverID uint64         `json:"receiver_id" binding:"required"`
	Message    string         `json:"message" binding:"required"`
	CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

func (c *Chat) InsertMessage(ch *Chat) error {
	err := db.DB.Create(&ch).Error
	if err != nil {
		return fmt.Errorf("can't insert message in database")
	}

	return nil
}
