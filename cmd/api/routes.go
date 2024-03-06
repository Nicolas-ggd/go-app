package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"websocket/cmd/api/websocket"
	"websocket/internal/models"
)

func Routes() *gin.Engine {
	router := gin.Default()

	app := &Application{}

	router.Use(app.CORSOptions())

	v1 := router.Group("/v1")

	authRoutes := v1.Group("/auth")
	{
		authRoutes.POST("/signup", app.InsertUserHandler(&Handler[models.UserRepository]{repository: models.UserRepository{}}))
		authRoutes.POST("/signin", app.UserAuthenticationHandler(&Handler[models.TokenRepository]{repository: models.TokenRepository{}}))
		authRoutes.POST("/logout", app.UserLogout(&Handler[models.TokenRepository]{repository: models.TokenRepository{}}))
	}

	v1.GET("/ws", websocket.ServeWs(app.Websocket))

	url := ginSwagger.URL("http://localhost:7000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	return router
}
