package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"websocket/internal/db"
	"websocket/pkg/http/middleware"
	http "websocket/pkg/http/rest"
)

func ServeApp() {
	database := db.ConnectionDB()
	defer database.Close()

	router := gin.Default()

	router.Use(middleware.CORSOptions())

	h := http.Handler{}

	api := router.Group("api")
	{
		AuthRoutes(api.Group("account"), &h)
	}

	err := router.Run(":7000")
	if err != nil {
		log.Fatal(err)
	}
}
