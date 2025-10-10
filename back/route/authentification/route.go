package authentification

import (
	"TimeManager/model"
	"TimeManager/repository"
	"TimeManager/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func RegisterAuthentificationRoutes(router *gin.Engine) {
	auth := router.Group("/authentification")
	{
		auth.POST("", authentificate)
	}
}

func authentificate(context *gin.Context) {

	var user AuthentificationUserInput

	if err := context.BindJSON(&user); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fetched_user, success, err := repository.Login(user.Email, user.Password)

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

	token, err := service.CreateToken(fetched_user.ID, fetched_user.Role, fetched_user.Team)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var output AuthentificationUserOutput

	if err = copier.Copy(&output, &fetched_user); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	output.Token = token

	if err = model.ValidateModel(&output); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, output)

}
