package user

import (
	"net/http"
	"testApi/model"
	"testApi/repository"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("", fetch_user)
		user.POST("/post", post_user)
	}
}

func fetch_user(context *gin.Context) {
	var username string = "Blue"
	user, _, _ := repository.FetchUser(username)

	context.JSON(http.StatusOK, gin.H{"user": user})
}

func post_user(context *gin.Context) {
	var user model.User

	if err := context.BindJSON(&user); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.CreateUser(user)

	context.JSON(http.StatusOK, gin.H{"user": user})
}
