package app

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"websocket/cmd/api/websocket"
)

func (a *Application) Routes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")

	authRoutes := v1.Group("/auth")
	{
		authRoutes.POST("/signup", a.Handler.InsertUserHandler)
		authRoutes.POST("/signin", a.Handler.UserAuthenticationHandler)
		authRoutes.POST("/logout", a.Handler.UserLogout)
	}

	v1.GET("/ws", websocket.ServeWs(a.Websocket))

	url := ginSwagger.URL("http://localhost:7000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	return router
}
