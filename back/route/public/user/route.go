package users

import (
	"TimeManager/middleware"
	"TimeManager/model"
	"TimeManager/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegisterUserRoutes(router *gin.Engine) {
	user := router.Group("/api/user").Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		user.GET("", FetchUser)
		user.POST("/update", UpdateUser)
	}
}

// FetchUser godoc
//
// @Summary     Fetch a single user
// @Description Fetch a single user and returns its info if successful
// @Tags        User
// @Produce     json
// @Success     200 {object} users.UserOutput
// @Failure     404 {string} string "User not found"
// @Failure     500 {string} string "Internal server error"
// @Router      /api/user [get]
// @Security ApiKeyAuth
func FetchUser(context *gin.Context) {

	_id, err := service.DecryptIDFromContextClaim(context)

	if err != nil {
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

	if user.Team != bson.NilObjectID {
		output.Team = user.Team.Hex()
	}

	if err = model.ValidateModel(&output); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": output})
}

// UpdateUser godoc
//
// @Summary     Update a single existing user
// @Description Update a single existing user from a specific user Input. Will return corresponding user if no modification is done
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       user body UpdateUserInput true "User fields to update"
// @Success     200 {object} users.UserOutput
// @Failure     400 {string} string "Invalid input"
// @Failure     500 {string} string "Internal server error"
// @Router      /api/user/update [post]
// @Security ApiKeyAuth
func UpdateUser(context *gin.Context) {

	_id, err := service.DecryptIDFromContextClaim(context)

	if err != nil {
		return
	}

	var input UpdateUserInput

	if err := context.BindJSON(&input); err != nil {
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fetched_user := &model.User{}

	if input.FirstName != "" || input.LastName != "" || input.Password != "" {

		update := bson.D{}

		if input.FirstName != "" {
			update = append(update, bson.D{{Key: "FirstName", Value: input.FirstName}}...)
		}

		if input.LastName != "" {
			update = append(update, bson.D{{Key: "LastName", Value: input.LastName}}...)
		}

		if input.Password != "" {
			update = append(update, bson.D{{Key: "Password", Value: input.Password}}...)
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

	if fetched_user.Team != bson.NilObjectID {
		output.Team = fetched_user.Team.Hex()
	}

	if err = model.ValidateModel(&output); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": output})

}
