package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"websocket/cmd/api/websocket"
	"websocket/internal/db"
	"websocket/internal/models"
)

func Routes() *gin.Engine {
	router := gin.Default()

	app := &Application{
		Repository: models.Repository{
			DB: db.DB,
		},
	}

	router.Use(app.CORSOptions())

	v1 := router.Group("/v1")

	authRoutes := v1.Group("/auth")
	{
		authRoutes.POST("/signup", app.InsertUserHandler())
		authRoutes.POST("/signin", app.UserAuthenticationHandler())
		authRoutes.Use(app.validateJWTToken())
		authRoutes.POST("/logout", app.UserLogout())
	}

	v1.Use(app.validateJWTToken())

	accountRoutes := v1.Group("/account")
	{
		accountRoutes.GET("/profile", app.GetAccount())
		accountRoutes.PATCH("/profile/update", app.UpdateProfileHandler())
		accountRoutes.POST("/delete", app.DeleteUserAccount())
		accountRoutes.POST("/recover", app.RecoverUserAccount())
	}

	v1.GET("/ws", websocket.ServeWs(app.Websocket))

	url := ginSwagger.URL("http://localhost:7000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	return router
}
