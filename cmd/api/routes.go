package api

import (
	"github.com/gin-gonic/gin"
	"websocket/cmd/api/websocket"
)

func (c *Config) routes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/ws", websocket.ServeWs(c.A.Websocket))

	return router
}
