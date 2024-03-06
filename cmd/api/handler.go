package api

import (
	"log/slog"
	"websocket/cmd/api/websocket"
)

type Application struct {
	Websocket *websocket.Websocket
	Logger    *slog.Logger
}

type Handler[T interface{}] struct {
	repository T
}
