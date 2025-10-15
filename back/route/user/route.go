package user

import (
	"TimeManager/middleware"
	"TimeManager/model"
	"TimeManager/repository"
	"TimeManager/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegisterUserRoutes(router *gin.Engine) {
	user := router.Group("/api/user").Use(middleware.AdminMiddleware())
	{
		user.GET("", FetchUser)
		user.POST("/create", PostUser)
		user.POST("/update", UpdateUser)
	}
}

func FetchUser(context *gin.Context) {

	claims, exists := context.Get("claims")

	if !exists {
		err := errors.New("INTERNAL ISSUE WITH TOKEN")
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	asserted_claims, ok := claims.(service.TokenClaims)

	if !ok {
		err := errors.New("INTERNAL ISSUE WITH TOKEN CONTENT")
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_id, err := bson.ObjectIDFromHex(asserted_claims.Subject)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := service.GetUserByID(_id)

	if err != nil {
		err = context.AbortWithError(http.StatusNotFound, err)
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var output UserOutput

	if err = copier.Copy(&output, user); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output.ID = user.ID.Hex()

	if user.Team != repository.NULL_ID {
		output.Team = user.Team.Hex()
	}

	if err = model.ValidateModel(&output); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": output})
}

func PostUser(context *gin.Context) {
	var user CreateUserInput

	if err := context.BindJSON(&user); err != nil {
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created_user, err := service.CreateUser(user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.Team)

	if err != nil {
		err = context.AbortWithError(http.StatusConflict, err)
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	var output UserOutput

	if err = copier.Copy(&output, created_user); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output.ID = created_user.ID.Hex()

	if created_user.Team != repository.NULL_ID {
		output.Team = created_user.Team.Hex()
	}

	if err = model.ValidateModel(&output); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": output})
}

func UpdateUser(context *gin.Context) {

	claims, exists := context.Get("claims")

	if !exists {
		err := errors.New("INTERNAL ISSUE WITH TOKEN")
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	asserted_claims, ok := claims.(service.TokenClaims)

	if !ok {
		err := errors.New("INTERNAL ISSUE WITH TOKEN CONTENT")
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_id, err := bson.ObjectIDFromHex(asserted_claims.Subject)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user UpdateUserInput

	if err := context.BindJSON(&user); err != nil {
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fetched_user := &model.User{}

	if user.FirstName != "" || user.LastName == "" || user.Password == "" {

		update := bson.D{}

		if user.FirstName != "" {
			update = append(update, bson.D{{Key: "FirstName", Value: user.FirstName}}...)
		}

		if user.LastName != "" {
			update = append(update, bson.D{{Key: "LastName", Value: user.LastName}}...)
		}

		if user.Password != "" {
			update = append(update, bson.D{{Key: "Password", Value: user.Password}}...)
		}

		fetched_user, err = service.UpdateUserByID(_id, update)

	} else {
		fetched_user, err = service.GetUserByID(_id)
	}

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var output UserOutput

	if err = copier.Copy(&output, fetched_user); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output.ID = fetched_user.ID.Hex()

	if fetched_user.Team != repository.NULL_ID {
		output.Team = fetched_user.Team.Hex()
	}

	if err = model.ValidateModel(&output); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": output})

}
