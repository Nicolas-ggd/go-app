package app

import (
	"github.com/gin-gonic/gin"
	"websocket/cmd/api/websocket"
)

func (a *Application) Routes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/ws", websocket.ServeWs(a.Websocket))

	return router
}
