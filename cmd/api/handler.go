package api

import "websocket/internal/models"

type Handler struct {
	UserService  models.UserModelInterface
	TokenService models.TokenModelInterface
	ChatService  models.ChatModelInterface
}

func NewHandler() *Handler {
	return &Handler{
		UserService:  &models.User{},
		TokenService: &models.Token{},
		ChatService:  &models.Chat{},
	}
}
