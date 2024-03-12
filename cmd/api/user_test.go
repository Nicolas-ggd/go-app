package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"websocket/internal/models"
)

func TestUserSignUp(t *testing.T) {

	TestSetup(t)
	app := &Application{}

	testCases := []struct {
		Name           string
		RequestPayload models.UserForm
		ExpectedStatus int
	}{
		{
			Name: "Failed",
			RequestPayload: models.UserForm{
				Name:     "",
				Email:    "",
				Password: "",
			},
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
		{
			Name: "Success",
			RequestPayload: models.UserForm{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "password",
			},
			ExpectedStatus: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			router := gin.Default()

			router.POST("/v1/auth/signup", app.InsertUserHandler())

			reqBody, err := json.Marshal(tc.RequestPayload)
			assert.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/v1/auth/signup", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)

			request.Header.Add("Content-Type", "application/json")

			router.ServeHTTP(rr, request)
			fmt.Println(rr.Body, "<--- rr body")
			assert.Equal(t, tc.ExpectedStatus, rr.Code)
		})
	}
}

func TestUserAuthentication(t *testing.T) {
	TestSetup(t)

	app := &Application{}

	testCase := []struct {
		Name           string
		RequestPayload models.AuthUser
		ExpectedStatus int
	}{
		{
			Name: "Failed",
			RequestPayload: models.AuthUser{
				Email:    "",
				Password: "",
			},
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
		{
			Name: "Success",
			RequestPayload: models.AuthUser{
				Email:    "test@gmail.com",
				Password: "random",
			},
			ExpectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			router := gin.Default()

			router.POST("/v1/auth/signin", app.UserAuthenticationHandler())

			reqBody, err := json.Marshal(tc.RequestPayload)
			assert.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/v1/auth/signin", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)

			request.Header.Add("Content-Type", "application/json")

			router.ServeHTTP(rr, request)
			assert.Equal(t, tc.ExpectedStatus, rr.Code)
		})
	}
}
