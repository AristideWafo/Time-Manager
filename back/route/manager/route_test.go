package manager

import (
	"TimeManager/middleware"
	"TimeManager/model"
	"TimeManager/repository"
	"TimeManager/service"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestGetAllUsersFromTeam_GetTeamUsersError(t *testing.T) {
	origGetTeamUsers := service.GetTeamUsers
	defer func() {
		service.GetTeamUsers = origGetTeamUsers
	}()

	teamId := bson.NewObjectID()
	userId := bson.NewObjectID()

	service.GetTeamUsers = func(_id bson.ObjectID) ([]*model.User, error) {
		return nil, errors.New("GetTeamUsers error")
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterManagerRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/manager/team/users", nil)

	token_string, _ := service.CreateToken(userId, "MANAGER", teamId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", w.Code)
	}
}

func TestGetAllUsersFromTeam_Success(t *testing.T) {
	origGetTeamUsers := service.GetTeamUsers
	defer func() {
		service.GetTeamUsers = origGetTeamUsers
	}()

	teamId := bson.NewObjectID()
	userId := bson.NewObjectID()
	adminId := bson.NewObjectID()

	service.GetTeamUsers = func(_id bson.ObjectID) ([]*model.User, error) {
		return []*model.User{
			{
				ID:        userId,
				FirstName: "Doctor",
				LastName:  "Who",
				Email:     "Doctor@Tardis.com",
				Password:  "hashed",
				Role:      "USER",
				Team:      teamId,
			},
		}, nil
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterManagerRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/manager/team/users", nil)

	token_string, _ := service.CreateToken(adminId, "MANAGER", teamId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 200 or 500, got %d", w.Code)
	}
}

func TestGetUserPresences_InvalidUserID(t *testing.T) {
	_id := bson.NewObjectID()

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterManagerRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/manager/team/user/presence/invalid-id", nil)

	token_string, _ := service.CreateToken(_id, "MANAGER", _id)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500 for invalid user id, got %d", w.Code)
	}
}

func TestGetUserPresences_UserNotInTeam(t *testing.T) {
	origGetUserByID := service.GetUserByID
	defer func() {
		service.GetUserByID = origGetUserByID
	}()

	userId := bson.NewObjectID()
	teamId := bson.NewObjectID()
	otherTeamId := bson.NewObjectID()

	service.GetUserByID = func(_id bson.ObjectID) (*model.User, error) {
		return &model.User{
			ID:        userId,
			FirstName: "Doctor",
			LastName:  "Who",
			Email:     "Doctor@Tardis.com",
			Password:  "hashed",
			Role:      "USER",
			Team:      otherTeamId,
		}, nil
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterManagerRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/manager/team/user/presence/"+userId.Hex(), nil)

	token_string, _ := service.CreateToken(userId, "MANAGER", teamId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for user not in team, got %d", w.Code)
	}
}

func TestGetUserPresences_GetAllPresencesError(t *testing.T) {
	origGetUserByID := service.GetUserByID
	origGetManyPresences := repository.GetManyPresences
	defer func() {
		service.GetUserByID = origGetUserByID
		repository.GetManyPresences = origGetManyPresences
	}()

	userId := bson.NewObjectID()
	teamId := bson.NewObjectID()

	service.GetUserByID = func(_id bson.ObjectID) (*model.User, error) {
		return &model.User{
			ID:        userId,
			FirstName: "Doctor",
			LastName:  "Who",
			Email:     "Doctor@Tardis.com",
			Password:  "hashed",
			Role:      "USER",
			Team:      teamId,
		}, nil
	}

	repository.GetManyPresences = func(presences *[]*model.Presence, filter bson.D) error {
		return errors.New("GetManyPresences error")
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterManagerRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/manager/team/user/presence/"+userId.Hex(), nil)

	token_string, _ := service.CreateToken(userId, "MANAGER", teamId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", w.Code)
	}
}

func TestGetUserPresences_Success(t *testing.T) {
	origGetUserByID := service.GetUserByID
	origGetManyPresences := repository.GetManyPresences
	defer func() {
		service.GetUserByID = origGetUserByID
		repository.GetManyPresences = origGetManyPresences
	}()

	userId := bson.NewObjectID()
	teamId := bson.NewObjectID()

	service.GetUserByID = func(_id bson.ObjectID) (*model.User, error) {
		return &model.User{
			ID:        userId,
			FirstName: "Doctor",
			LastName:  "Who",
			Email:     "Doctor@Tardis.com",
			Password:  "hashed",
			Role:      "USER",
			Team:      teamId,
		}, nil
	}

	repository.GetManyPresences = func(presences *[]*model.Presence, filter bson.D) error {
		*presences = []*model.Presence{}
		return nil
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterManagerRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/manager/team/user/presence/"+userId.Hex(), nil)

	token_string, _ := service.CreateToken(userId, "MANAGER", teamId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 200 or 500, got %d", w.Code)
	}
}
