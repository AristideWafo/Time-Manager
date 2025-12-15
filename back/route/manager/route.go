package manager

import (
	"TimeManager/middleware"
	"TimeManager/model"
	"TimeManager/route/public/presence"
	users "TimeManager/route/public/user"
	"TimeManager/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegisterAdminRoutes(router *gin.Engine) {
	manager := router.Group("/api/manager").Use(middleware.ManagerMiddleware())
	{

		manager.GET("/team/users/:name", GetAllUsersFromTeam)
		manager.GET("/team/user/presence/:id", GetUserPresences)
	}
}

// GetAllUsersFromTeam godoc
//
// @Summary     Return all users from the manager's team
// @Description Return all users belonging to the team the manager manages (team is derived from the manager's token/claim)
// @Tags        Manager
// @Produce     json
// @Success     200  {array}   users.UserOutput
// @Failure     400  {string}  string  "Invalid input"
// @Failure     404  {string}  string  "Not found"
// @Failure     500  {string}  string  "Internal server error"
// @Router      /api/manager/team/users/{name} [get]
// @Security ApiKeyAuth
func GetAllUsersFromTeam(context *gin.Context) {

	team_id, err := service.DecryptTeamFromContextClaim(context)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	all_users, err := service.GetTeamUsers(team_id)

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
		all_users_output[index].Team = user.Team.Hex()

		if err = model.ValidateModel(&all_users_output[index]); err != nil {
			err = context.AbortWithError(http.StatusInternalServerError, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"users": all_users_output})

}

// GetUserPresences godoc
//
// @Summary     Get presences for a user
// @Description Returns all presence records for the specified user if the user belongs to the manager's team
// @Tags        Manager
// @Produce     json
// @Param       id   path     string  true  "User id"
// @Success     200  {array}  presence.PresenceOutput
// @Failure     400  {string} string  "Invalid input / not member of team"
// @Failure     404  {string} string  "User not found"
// @Failure     500  {string} string  "Internal server error"
// @Router      /api/manager/team/user/presence/{id} [get]
// @Security ApiKeyAuth
func GetUserPresences(context *gin.Context) {

	_id, err := bson.ObjectIDFromHex(context.Param("id"))

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	team_id, err := service.DecryptTeamFromContextClaim(context)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	is_user_in_team, err := service.IsUserInTeam(_id, team_id)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !is_user_in_team {
		err := errors.New("User is not a member of your team")
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	presences, err := service.GetAllPresencesByUserID(_id)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var presences_output []presence.PresenceOutput

	for index, pres := range *presences {

		presences_output = append(presences_output, presence.PresenceOutput{})

		if err = copier.Copy(&presences_output[index], pres); err != nil {
			err = context.AbortWithError(http.StatusInternalServerError, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err = model.ValidateModel(&presences_output[index]); err != nil {
			err = context.AbortWithError(http.StatusInternalServerError, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"presences": presences_output})

}
