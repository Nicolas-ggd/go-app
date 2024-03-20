package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"strings"
	"websocket/internal/models"
	"websocket/pkg/http/ws"
)

type Handler struct {
	Websocket  *ws.Websocket
	Logger     *slog.Logger
	Repository models.Repository
}

func (h *Handler) GenerateResponse(status int, msg string, error any, c *gin.Context) {
	c.JSON(status, gin.H{"message": msg, "error": error})
}

func (h *Handler) ValidateError(err error, c *gin.Context) {
	var valid validator.ValidationErrors

	var errs = make(map[string]ErrorResponse)

	if errors.As(err, &valid) {
		for _, val := range valid {
			errs[strings.ToLower(val.Field())] = ErrorResponse{
				Field: val.Field(),
				Error: val.Error(),
				Tag:   val.Tag(),
				Param: val.Param(),
			}
		}
	}

	h.GenerateResponse(http.StatusUnprocessableEntity, "Validation Error", errs, c)
}
