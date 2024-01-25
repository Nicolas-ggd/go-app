package api

import (
	"fmt"
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

	isExist, err := h.UserService.CheckUserIsExist(chat.ReceiverID)
	fmt.Println(chat.ReceiverID, isExist)
	if err != nil || !isExist {
		c.JSON(http.StatusNotFound, gin.H{"error": "User doesn't exist"})
		return
	}

	err = h.ChatService.InsertMessage(&chat)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &chat})
}
