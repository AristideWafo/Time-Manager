package team

import (
	"TimeManager/middleware"
	"TimeManager/model"
	"TimeManager/route/user"
	"TimeManager/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterTeamRoutes(router *gin.Engine) {
	team := router.Group("/api/team").Use(middleware.ManagerMiddleware())
	{
		team.GET("/:name", FetchTeam)
		team.POST("/create", PostTeam)
		team.PUT("/update", UpdateTeam)
		team.GET("/:name/users", FetchAllUserFromTeam)
	}
}

// FetchTeam godoc
//
//	@Summary		Fetch a single team
//	@Description	Fetch a single team and returns its info if successful
//	@Tags			Team
//	@Produce		json
//	@Success		200		{object}	TeamOutput
//	@Failure		400		"Invalid Input"
//	@Failure		404		"Team not found"
//	@Failure		500		"Internal server error"
//	@Router			/api/team/:name [get]
func FetchTeam(context *gin.Context) {
	name := context.Param("name")

	if name == "" {
		err := errors.New("team name must be specified")
		err = context.AbortWithError(http.StatusBadRequest, err)
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	team, err := service.GetTeamByName(name)

	if err != nil {
		err = context.AbortWithError(http.StatusNotFound, err)
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var output TeamOutput

	if err = copier.Copy(&output, team); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output.ID = team.ID.Hex()

	if err = model.ValidateModel(&output); err != nil {
		err = context.AbortWithError(http.StatusInternalServerError, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"team": output})

}

// PostTeam godoc
//
//	@Summary		Create a single new team
//	@Description	Create a single new team from a specific team Input. Will return an error if team with same name already exists.
//	@Tags			Team
//	@Accept			json
//	@Produce		json
//	@Param			team	body		CreateTeamInput	true	"Team to create"
//	@Success		200		{object}	TeamOutput
//	@Failure		400		"Invalid input"
//	@Failure		409		"Conflict : team already exists"
//	@Failure		500		"Internal server error"
//	@Router			/api/team/create [post]
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

	var output TeamOutput

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
//	@Summary		Update a single existing team
//	@Description	Update a single existing team from a specific team Input. Will return corresponding team if no modification is done
//	@Tags			Team
//	@Accept			json
//	@Produce		json
//	@Param			team	body		UpdateTeamInput	true	"Current team name and New team name"
//	@Success		200		{object}	TeamOutput
//	@Failure		400		"Invalid input"
//	@Failure		404		"Team not found"
//	@Failure		409		"Conflict : team already exists"
//	@Failure		500		"Internal server error"
//	@Router			/api/team/update [put]
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

	var output TeamOutput

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

// FetchAllUserFromTeam godoc
//
//		@Summary		Return all users from a team
//		@Description	Return all users from the specified team. Will return an error if team doesn't exist. Will return an empty array if not user is in the team
//		@Tags			Team
//		@Produce		json
//		@Success		200		{array}		user.UsersInput
//		@Failure		400		"Invalid input"
//	 	@Failure		404		"Team Not Found"
//		@Failure		500		"Internal server error"
//		@Router			/api/team/:name/users [get]
func FetchAllUserFromTeam(context *gin.Context) {
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

	var output []user.UserOutput

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

	context.JSON(http.StatusOK, gin.H{"team": output})
}
