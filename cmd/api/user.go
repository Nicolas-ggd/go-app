package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/internal/models"
)

// @Tags   User Registration
// @Summary Sign up user generating jwt token
// @Description register user
// @Accept  json
// @Produce  json
// @Param   email     path    string     true        "Email"
// @Param   password     path    string     true        "Password"
// @Success 200 {object} models.User	"ok"
// @Failure 401 {object} models.ErrorResponse "Error"
// @Failure 404 {object} models.ErrorResponse "Not Found"
// @Failure 422 {object} models.ErrorResponse "Error"
// @Router /auth/signup [post]
func (h *Handler) InsertUserHandler(c *gin.Context) {
	var userAuth models.AuthUser

	err := c.ShouldBind(&userAuth)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request data, missing request fields!"})
		return
	}

	user, err := h.UserService.InsertUser(&userAuth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &user})
}

// @Tags   User Authentication
// @Summary Sign In user generating jwt token
// @Description authenticate user
// @Accept  json
// @Produce  json
// @Param   email     path    string     true        "Email"
// @Param   password     path    string     true        "Password"
// @Success 200 {object} models.User	"ok"
// @Failure 401 {object} models.ErrorResponse "Error"
// @Failure 404 {object} models.ErrorResponse "Not Found"
// @Failure 422 {object} models.ErrorResponse "Error"
// @Router /auth/signin [post]
func (h *Handler) UserAuthenticationHandler(c *gin.Context) {
	var userAuth models.AuthUser

	err := c.ShouldBind(&userAuth)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request data, missing request fields!"})
		return
	}

	user, err := h.UserService.GetByEmail(userAuth.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	err = models.CompareHashAndPasswordBcrypt(user.Password, userAuth.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	token, err := h.TokenService.CreateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	userToken := &models.Token{
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

	err = h.TokenService.InsertToken(userToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't insert token, something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

// @Tags   User Sign out with token
// @Summary User Sign out
// @Description User Sign out
// @Accept  json
// @Produce  json
// @Success 200 {object} string	"ok"
// @Failure 401 {object} models.ErrorResponse "Error"
// @Failure 404 {object} models.ErrorResponse "Not Found"
// @Failure 422 {object} models.ErrorResponse "Error"
// @Router /auth/logout [post]
func (h *Handler) UserLogout(c *gin.Context) {
	userId, err := h.ParseJWTClaims(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err = h.TokenService.DeleteToken(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully!"})
}
