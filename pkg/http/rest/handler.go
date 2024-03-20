package http

import (
	"log/slog"
	"websocket/internal/models"
	"websocket/pkg/http/ws"
)

type Handler struct {
	Websocket  *ws.Websocket
	Logger     *slog.Logger
	Repository models.Repository
}
