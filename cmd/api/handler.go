package api

import (
	"websocket/cmd/api/websocket"
	"websocket/internal/models"
)

type Handler struct {
	UserService  models.UserModelInterface
	TokenService models.TokenModelInterface
	ChatService  models.ChatModelInterface
	Websocket    *websocket.Websocket
}

func NewHandler() *Handler {
	return &Handler{
		UserService:  &models.User{},
		TokenService: &models.Token{},
		ChatService:  &models.Chat{},
	}
}
