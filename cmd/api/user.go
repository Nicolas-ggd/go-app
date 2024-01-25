package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/internal/models"
)

func (h *Handler) InsertUserHandler(c *gin.Context) {
	var userSignUp models.UserSignUp

	err := c.ShouldBind(&userSignUp)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request data, missing request fields!"})
		return
	}

	user, err := h.UserService.InsertUser(&userSignUp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &user})
}
