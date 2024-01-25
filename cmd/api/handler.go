package api

import "websocket/internal/models"

type Handler struct {
	UserService models.UserModelInterface
}

func NewHandler(userService models.UserModelInterface) *Handler {
	return &Handler{
		UserService: userService,
	}
}
