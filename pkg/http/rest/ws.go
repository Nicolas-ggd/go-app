package http

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"websocket/pkg/http/helpers"
	"websocket/pkg/http/ws"
)

// ServeWs handles websocket request from the peer
func (h *Handler) ServeWs() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientKey := c.Query("key")
		if clientKey == "" {
			return
		}

		userObject := helpers.GetTokenClaim(c)
		if userObject == nil {
			return
		}

		conn, err := ws.ConnectionUpgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		// initialize websocket client
		client := &ws.Client{
			Ws:       h.Websocket,
			Conn:     conn,
			ClientId: strconv.FormatUint(userObject.UserId, 10),
			Send:     make(chan []byte, 256),
		}

		// register client
		client.Ws.Register <- client

		// Allow collection of memory referenced by the caller by doing all work in
		// another goroutines.
		go client.WritePump()
		go client.ReadPump()
	}
}
