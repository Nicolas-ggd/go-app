package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/internal/models"
)

func (h *Handler) InsertMessageHandler(c *gin.Context) {
	var chat models.Chat

	err := c.ShouldBind(&chat)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request data, missing request fields"})
		return
	}

	err = h.ChatService.InsertMessage(&chat)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &chat})
}
