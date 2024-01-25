package api

import "websocket/internal/models"

type Handler struct {
	UserService  models.UserModelInterface
	TokenService models.TokenModelInterface
}

func NewHandler(userService models.UserModelInterface, TokenService models.TokenModelInterface) *Handler {
	return &Handler{
		UserService:  userService,
		TokenService: TokenService,
	}
}
