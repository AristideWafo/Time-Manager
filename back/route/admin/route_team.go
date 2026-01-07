package admin

import (
	"TimeManager/model"
	"TimeManager/repository"
	teams "TimeManager/route/public/team"
	users "TimeManager/route/public/user"
	"TimeManager/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GetAllTeams godoc
//
// @Summary     Get all teams
// @Description Returns a list of all teams
// @Tags        Admin
// @Produce     json
// @Success     200  {array}   teams.TeamOutput
// @Failure     500  {string}  string  "Internal server error"
// @Router      /api/admin/all/team [get]
// @Security ApiKeyAuth
func GetAllTeams(context *gin.Context) {

	var all_teams []*model.Team
	err := repository.GetManyTeams(&all_teams, bson.D{})

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var all_teams_output []teams.TeamOutput

	for index, team := range all_teams {

		all_teams_output = append(all_teams_output, teams.TeamOutput{})

		if err = copier.Copy(&all_teams_output[index], team); err != nil {
			err = context.AbortWithError(http.StatusInternalServerError, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		all_teams_output[index].ID = team.ID.Hex()

		if err = model.ValidateModel(&all_teams_output[index]); err != nil {
			err = context.AbortWithError(http.StatusInternalServerError, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"teams": all_teams_output})

}

// GetAllUsersFromTeam godoc
//
// @Summary     Return all users from a team
// @Description Return all users from the specified team. Will return an error if team doesn't exist. Will return an empty array if no user is in the team
// @Tags        Admin
// @Produce     json
// @Param       name  path      string  true  "Team name"
// @Success     200   {array}   users.UserOutput
// @Failure     400   {string}  string  "Invalid input"
// @Failure     404   {string}  string  "Team not found"
// @Failure     500   {string}  string  "Internal server error"
// @Router      /api/admin/team/users/{name} [get]
// @Security ApiKeyAuth
func GetAllUsersFromTeam(context *gin.Context) {

	name := context.Param("name")

	if name == "" {
		err := errors.New("team name must be specified")
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	team, err := service.GetTeamByName(name)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = context.AbortWithError(http.StatusNotFound, err)
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
	}
	_id := team.ID
	user_list, err := service.GetTeamUsers(_id)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var output []users.UserOutput

	if err = copier.Copy(&output, user_list); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i := range output {
		output[i].ID = user_list[i].ID.Hex()
		output[i].Team = user_list[i].Team.Hex()

		if err := model.ValidateModel(&output[i]); err != nil {
			err = context.AbortWithError(http.StatusInternalServerError, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"users": output})
}

// PostTeam godoc
//
// @Summary     Create a single new team
// @Description Create a single new team from a specific team Input. Will return an error if team with same name already exists.
// @Tags        Admin
// @Accept      json
// @Produce     json
// @Param       team  body      CreateTeamInput  true  "Team to create"
// @Success     200   {object}  teams.TeamOutput
// @Failure     400   {string}  string  "Invalid input"
// @Failure     409   {string}  string  "Conflict : team already exists"
// @Failure     500   {string}  string  "Internal server error"
// @Router      /api/admin/team/create [post]
// @Security ApiKeyAuth
func PostTeam(context *gin.Context) {
	var team CreateTeamInput

	if err := context.BindJSON(&team); err != nil {
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created_team, err := service.CreateTeam(team.Name)

	if err != nil {
		err = context.AbortWithError(http.StatusConflict, err)
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	var output teams.TeamOutput

	if err = copier.Copy(&output, created_team); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output.ID = created_team.ID.Hex()

	if err = model.ValidateModel(&output); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"team": output})

}

// UpdateTeam godoc
//
// @Summary     Update a single existing team
// @Description Update a single existing team from a specific team Input. Will return corresponding team if no modification is done
// @Tags        Admin
// @Accept      json
// @Produce     json
// @Param       team  body      UpdateTeamInput  true  "Current team name and New team name"
// @Success     200   {object}  teams.TeamOutput
// @Failure     400   {string}  string  "Invalid input"
// @Failure     404   {string}  string  "Team not found"
// @Failure     409   {string}  string  "Conflict : team already exists"
// @Failure     500   {string}  string  "Internal server error"
// @Router      /api/admin/team/update [put]
// @Security ApiKeyAuth
func UpdateTeam(context *gin.Context) {

	var input UpdateTeamInput

	if err := context.BindJSON(&input); err != nil {
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fetched_team, err := service.GetTeamByName(input.CurrentName)

	if _, err := service.GetTeamByName(input.NewName); err == nil {
		err = errors.New("new name already exists")
		err = context.AbortWithError(http.StatusConflict, err)
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		err = context.AbortWithError(http.StatusNotFound, err)
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if input.NewName != "" {
		update := bson.D{}
		update = append(update, bson.D{{Key: "Name", Value: input.NewName}}...)

		fetched_team, err = service.UpdateTeam(input.CurrentName, update)
	}

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var output teams.TeamOutput

	if err = copier.Copy(&output, fetched_team); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output.ID = fetched_team.ID.Hex()

	if err = model.ValidateModel(&output); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"team": output})

}

// DeleteTeamById godoc
// @Summary     Delete a single team by id
// @Description Delete a single team
// @Tags        Admin
// @Produce     json
// @Param       id   path     string            true  "Team id"
// @Success 204 {string}  "OK"
// @Failure     400  {string} string "Invalid input"
// @Failure     500  {string} string "Internal server error"
// @Router      /api/admin/team/delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteTeamById(context *gin.Context) {

	id, err := bson.ObjectIDFromHex(context.Param("id"))

	if err != nil {
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.DeleteTeamById(id)

	if err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}
