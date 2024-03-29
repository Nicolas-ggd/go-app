package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/internal/models"
	"websocket/pkg/http/helpers"
)

func (h *Handler) RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userAuth models.UserForm

		err := c.ShouldBind(&userAuth)
		if err != nil {
			h.ValidateError(err, c)
			return
		}

		user, err := h.Repository.InsertUser(&userAuth)
		if err != nil {
			h.GenerateResponse(http.StatusInternalServerError, err.Error(), err, c)
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": &user})
	}
}

func (h *Handler) LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userAuth models.AuthUser

		err := c.ShouldBind(&userAuth)
		if err != nil {
			h.ValidateError(err, c)
			return
		}

		user, err := h.Repository.GetByEmail(userAuth.Email)
		if err != nil {
			h.GenerateResponse(http.StatusNotFound, err.Error(), err, c)
			return
		}

		err = models.CompareHashAndPasswordBcrypt(user.Password, userAuth.Password)
		if err != nil {
			h.GenerateResponse(http.StatusUnauthorized, "Invalid email or password", err, c)
			return
		}

		token, err := h.Repository.CreateJWT(user.ID)
		if err != nil {
			h.GenerateResponse(http.StatusInternalServerError, err.Error(), err, c)
			return
		}

		userToken := models.Token{
			UserID: user.ID,
			Hash:   []byte(token),
			Type:   models.Auth,
		}

		var usrToken []byte

		for _, token := range user.Token {
			usrToken = token.Hash
		}

		if usrToken != nil {
			c.JSON(http.StatusOK, gin.H{"access_token": token})
			return
		}

		err = h.Repository.InsertToken(&userToken)
		if err != nil {
			h.GenerateResponse(http.StatusInternalServerError, err.Error(), err, c)
			return
		}

		c.JSON(http.StatusOK, gin.H{"access_token": token})
	}
}

func (h *Handler) LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userObject := helpers.GetTokenClaim(c)
		if userObject == nil {
			return
		}

		_, err := h.Repository.DeleteToken(userObject.UserId)
		if err != nil {
			h.GenerateResponse(http.StatusNotFound, err.Error(), err, c)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Logout successfully!"})
	}
}

func (h *Handler) DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		userObject := helpers.GetTokenClaim(c)
		if userObject == nil {
			return
		}

		err := h.Repository.DeleteAccount(userObject.UserId)
		if err != nil {
			h.GenerateResponse(http.StatusNotFound, err.Error(), err, c)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully!"})
	}
}

func (h *Handler) RecoverAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		userObject := helpers.GetTokenClaim(c)
		if userObject == nil {
			return
		}

		err := h.Repository.RecoverAccount(userObject.UserId)
		if err != nil {
			h.GenerateResponse(http.StatusNotFound, err.Error(), err, c)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Your account recovered!"})
	}
}

func (h *Handler) GetAccountInformation() gin.HandlerFunc {
	return func(c *gin.Context) {
		userObject := helpers.GetTokenClaim(c)
		if userObject == nil {
			return
		}

		user, err := h.Repository.GetUserProfile(userObject.UserId)
		if err != nil {
			h.GenerateResponse(http.StatusNotFound, err.Error(), err, c)
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": &user})
	}
}

func (h *Handler) UpdateProfileHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u models.User

		userObject := helpers.GetTokenClaim(c)
		if userObject == nil {
			return
		}

		err := c.ShouldBind(&u)
		if err != nil {
			h.ValidateError(err, c)
			return
		}

		user, err := h.Repository.UpdateProfile(userObject.UserId, &u)
		if err != nil {
			h.GenerateResponse(http.StatusInternalServerError, err.Error(), err, c)
			return
		}

		c.JSON(http.StatusOK, &user)
	}
}
