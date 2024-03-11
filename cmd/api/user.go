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
func (app *Application) InsertUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userAuth models.UserForm

		err := c.ShouldBind(&userAuth)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request data, missing request fields!"})
			return
		}

		user, err := app.Repository.InsertUser(&userAuth)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": &user})
	}
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
func (app *Application) UserAuthenticationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userAuth models.AuthUser

		err := c.ShouldBind(&userAuth)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request data, missing request fields!"})
			return
		}

		user, err := app.Repository.GetByEmail(userAuth.Email)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		err = models.CompareHashAndPasswordBcrypt(user.Password, userAuth.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}

		token, err := app.Repository.CreateJWT(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
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

		err = app.Repository.InsertToken(&userToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"access_token": token})
	}
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
func (app *Application) UserLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := app.ParseJWTClaims(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		_, err = app.Repository.DeleteToken(userId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login successfully!"})
	}
}
