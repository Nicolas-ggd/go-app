package routes

import (
	"github.com/gin-gonic/gin"
	"websocket/pkg/http/middleware"
	http "websocket/pkg/http/rest"
)

func AuthRoutes(e *gin.RouterGroup, auth *http.Handler) {
	e.POST("/register", auth.RegisterHandler())
	e.POST("/login", auth.LoginHandler())
	e.Use(middleware.ValidateJWTToken())
	e.POST("/logout", auth.LogoutHandler())
	e.GET("/profile", auth.GetAccountInformation())
	e.PATCH("/profile/update", auth.UpdateProfileHandler())
	e.POST("/delete", auth.DeleteAccount())
	e.POST("/recover", auth.RecoverAccount())
}
