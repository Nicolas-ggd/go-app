package api

import (
	"log/slog"
	"websocket/cmd/api/websocket"
	"websocket/internal/models"
)

type Application struct {
	Websocket  *websocket.Websocket
	Logger     *slog.Logger
	Repository models.Repository
}
