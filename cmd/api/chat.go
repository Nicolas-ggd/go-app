package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/internal/models"
)

// @Tags   Insert Message
// @Summary Insert Message
// @Description Insert Message
// @Accept  json
// @Produce  json
// @Param   message     path    string     true        "Message"
// @Param   receiver_id     path    string     true        "ReceiverID"
// @Param   sender_id     path    string     true        "SenderID"
// @Success 200 {object} models.Chat	"ok"
// @Failure 401 {object} models.ErrorResponse "Error"
// @Failure 404 {object} models.ErrorResponse "Not Found"
// @Failure 422 {object} models.ErrorResponse "Error"
// @Router /chat/message [post]
func (h *Handler) InsertMessageHandler(c *gin.Context) {
	var chat models.Chat

	err := c.ShouldBind(&chat)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request data, missing request fields"})
		return
	}

	isExist, err := h.UserService.CheckUserIsExist(chat.ReceiverID)
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
