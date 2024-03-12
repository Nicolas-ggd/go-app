package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
	"websocket/internal/db"
)

func TestSetup(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := db.TestDatabaseConnection()
	assert.NoError(t, err)
}
