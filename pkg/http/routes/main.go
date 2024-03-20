package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"websocket/internal/db"
	"websocket/internal/models"
	"websocket/pkg/http/middleware"
	http "websocket/pkg/http/rest"
)

func ServeApp() {
	router := gin.Default()

	router.Use(middleware.CORSOptions())

	h := http.Handler{
		Repository: models.Repository{
			DB: db.DB,
		},
	}

	api := router.Group("api")
	{
		AuthRoutes(api.Group("account"), &h)
	}

	err := router.Run(":7000")
	if err != nil {
		log.Fatal(err)
	}
}
