package autentification

import (
	"errors"
	"net/http"
	"testApi/service"

	"github.com/gin-gonic/gin"
)

func RegisterAuthentificationRoutes(router *gin.Engine) {
	auth := router.Group("/authentification")
	{
		auth.POST("", authentificate)
	}
}

func authentificate(context *gin.Context) {

	var user AutentificationUserInput

	if err := context.BindJSON(&user); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fetch_user, success, err := service.Authentificate(user.Username, user.Password)

	if !success {
		err := errors.New("USER NOT FOUND")
		context.AbortWithError(http.StatusNotFound, error(err))
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := service.CreateToken(fetch_user.Username, fetch_user.Role, fetch_user.Team)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})

}
