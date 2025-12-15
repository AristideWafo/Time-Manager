package admin

import (
	"TimeManager/model"
	"TimeManager/repository"
	users "TimeManager/route/public/user"
	"TimeManager/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetUserById godoc
// @Summary     Fetch a single user by id
// @Description Fetch a single user and returns its info if successful
// @Tags        Admin
// @Produce     json
// @Param       id   path     string            true  "User id"
// @Success     200  {object} users.UserOutput
// @Failure     400  {string} string "Invalid input"
// @Failure     404  {string} string "User not found"
// @Failure     500  {string} string "Internal server error"
// @Router      /api/admin/{id} [get]
// @Security ApiKeyAuth
func GetUserById(context *gin.Context) {

	id, err := bson.ObjectIDFromHex(context.Param("id"))

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := service.GetUserByID(id)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var output users.UserOutput

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

}

// PostUser godoc
//
// @Summary     Create a single new user
// @Description Create a single new user from a specific user Input. Will return an error if user already exists.
// @Tags        Admin
// @Accept      json
// @Produce     json
// @Param       user  body      CreateUserInput  true  "User to create"
// @Success     200   {object}  users.UserOutput
// @Failure     400   {string}  string  "Invalid input"
// @Failure     409   {string}  string  "Conflict : user already exists"
// @Failure     500   {string}  string  "Internal server error"
// @Router      /api/admin/user/create [post]
// @Security ApiKeyAuth
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

	var output users.UserOutput

	if err = copier.Copy(&output, created_user); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output.ID = created_user.ID.Hex()

	if created_user.Team != bson.NilObjectID {
		output.Team = created_user.Team.Hex()
	}

	if err = model.ValidateModel(&output); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": output})
}

// UpdateUserTeam godoc
//
// @Summary     Update a user's team
// @Description Update the team of a user identified by id
// @Tags        Admin
// @Accept      json
// @Produce     json
// @Param       id    path     string              true  "User id"
// @Param       team  body     UpdateUserTeamInput true  "Team to associate with the user"
// @Success     200   {object} users.UserOutput
// @Failure     400   {string} string "Invalid input"
// @Failure     404   {string} string "User not found"
// @Failure     500   {string} string "Internal server error"
// @Router      /api/admin/update/team/user/{id} [put]
// @Security ApiKeyAuth
func UpdateUserTeam(context *gin.Context) {

	id, err := bson.ObjectIDFromHex(context.Param("id"))

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input UpdateUserTeamInput

	if err := context.BindJSON(&input); err != nil {
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.D{{Key: "Team", Value: input.TeamId}}

	user, err := service.UpdateUserByID(id, update)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var output users.UserOutput

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

// GetAllUsers godoc
//
// @Summary     Get all users
// @Description Returns a list of all users
// @Tags        Admin
// @Produce     json
// @Success     200  {array}   users.UserOutput
// @Failure     500  {string}  string  "Internal server error"
// @Router      /api/admin/all/user [get]
// @Security ApiKeyAuth
func GetAllUsers(context *gin.Context) {

	var all_users []*model.User
	err := repository.GetManyUsers(&all_users, bson.D{})

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var all_users_output []users.UserOutput

	for index, user := range all_users {

		all_users_output = append(all_users_output, users.UserOutput{})

		if err = copier.Copy(&all_users_output[index], user); err != nil {
			err = context.AbortWithError(http.StatusInternalServerError, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		all_users_output[index].ID = user.ID.Hex()

		if user.Team != bson.NilObjectID {
			all_users_output[index].Team = user.Team.Hex()
		}

		if err = model.ValidateModel(&all_users_output[index]); err != nil {
			err = context.AbortWithError(http.StatusInternalServerError, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"users": all_users_output})

}
