package app

import (
	"websocket/cmd/api"
	"websocket/cmd/api/websocket"
)

type Application struct {
	Websocket *websocket.Websocket
	Handler   *api.Handler
}

func NewApplication(websocket *websocket.Websocket, handler *api.Handler) *Application {
	return &Application{
		Websocket: websocket,
		Handler:   handler,
	}
}
