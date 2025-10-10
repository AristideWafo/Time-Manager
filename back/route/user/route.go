package user

import (
	"TimeManager/model"
	"TimeManager/repository"
	"TimeManager/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	_id, err := primitive.ObjectIDFromHex(asserted_claims.Subject)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	user, success, err := repository.GetUser(_id)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if !success {
		err = errors.New("USER NOT FOUND")
		context.AbortWithError(http.StatusNotFound, err)
		context.JSON(http.StatusNotFound, gin.H{"error": asserted_claims.Subject})
	}

	var output UserOutput

	if err = copier.Copy(&output, &user); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err = model.ValidateModel(&output); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"user": output})
}

func PostUser(context *gin.Context) {
	var user CreateUserInput

	if err := context.BindJSON(&user); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created_user, success, err := repository.CreateUser(user.Email, user.Password)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if !success {
		context.AbortWithError(http.StatusConflict, err)
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	}

	var output UserOutput

	if err = copier.Copy(&output, &created_user); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err = model.ValidateModel(&output); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"user": output})
}
