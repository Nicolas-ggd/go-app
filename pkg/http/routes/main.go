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

	router.Use(middleware.CORSOptions())

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
