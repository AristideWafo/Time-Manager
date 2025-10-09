package user

import (
	"TimeManager/repository"
	"TimeManager/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func RegisterUserRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("", FetchUser)
		user.POST("/create", PostUser)
	}
}

func FetchUser(context *gin.Context) {

	claims, exists := context.Get("claims")

	if !exists {
		err := errors.New("INTERNAL ISSUE WITH TOKEN")
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	asserted_claims, ok := claims.(service.TokenClaims)

	if !ok {
		err := errors.New("INTERNAL ISSUE WITH TOKEN CONTENT")
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	user, success, err := repository.GetUser(asserted_claims.Subject)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if !success {
		err := errors.New("USER NOT FOUND")
		context.AbortWithError(http.StatusNotFound, err)
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	var output UserOutput

	err = copier.Copy(&output, &user)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"user": output})
}

func PostUser(context *gin.Context) {
	var user UserInput

	if err := context.BindJSON(&user); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created_user, success, err := repository.CreateUser(user.Username, user.Password)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if !success {
		context.AbortWithError(http.StatusConflict, err)
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	}

	var output UserOutput

	err = copier.Copy(&output, &created_user)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"user": output})
}
