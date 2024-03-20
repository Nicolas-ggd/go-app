package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"websocket/internal/db"
	"websocket/internal/models"
	"websocket/pkg/http/middleware"
	http "websocket/pkg/http/rest"
	"websocket/pkg/http/ws"
)

func ServeApp() {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(middleware.CORSOptions(), gin.Logger(), gin.Recovery())

	h := http.Handler{
		Repository: models.Repository{
			DB: db.DB,
		},
		Websocket: ws.NewWebsocket(),
	}

	api := router.Group("api")
	{
		AuthRoutes(api.Group("account"), &h)
	}

	api.GET("/ws", h.ServeWs())

	err := router.Run(":7000")
	if err != nil {
		log.Fatal(err)
	}
}
