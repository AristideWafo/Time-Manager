package admin

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

func TestPostTeam_InvalidInput(t *testing.T) {
	_id := bson.NewObjectID()

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterAdminRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/admin/team/create", nil)
	req.Header.Set("Content-Type", "application/json")

	token_string, _ := service.CreateToken(_id, "ADMIN", _id)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for invalid input, got %d", w.Code)
	}
}

func TestUpdateTeam_InvalidInput(t *testing.T) {
	_id := bson.NewObjectID()

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterAdminRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/admin/team/update", nil)
	req.Header.Set("Content-Type", "application/json")

	token_string, _ := service.CreateToken(_id, "ADMIN", _id)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for invalid input, got %d", w.Code)
	}
}

func TestGetAllUsersFromTeam_EmptyTeamName(t *testing.T) {
	_id := bson.NewObjectID()

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterAdminRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/admin/team/users/", nil)

	token_string, _ := service.CreateToken(_id, "ADMIN", _id)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest && w.Code != http.StatusNotFound {
		t.Fatalf("expected status 400 or 404, got %d", w.Code)
	}
}

func TestGetAllTeams_DatabaseError(t *testing.T) {

	orig := repository.GetManyTeams
	defer func() { repository.GetManyTeams = orig }()

	_id := bson.NewObjectID()
	repository.GetManyTeams = func(teams *[]*model.Team, filter bson.D) error {
		return errors.New("GetManyTeams error")
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterAdminRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/admin/all/team", nil)

	token_string, _ := service.CreateToken(_id, "ADMIN", _id)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", w.Code)
	}
}

func TestGetAllUsersFromTeam_ValidTeamName(t *testing.T) {

	orig := service.GetTeamByName
	defer func() { service.GetTeamByName = orig }()

	teamId := bson.NewObjectID()
	adminId := bson.NewObjectID()

	service.GetTeamByName = func(name string) (*model.Team, error) {
		return &model.Team{
			ID:   teamId,
			Name: name,
		}, nil
	}

	orig_2 := service.GetTeamUsers
	defer func() { service.GetTeamUsers = orig_2 }()

	service.GetTeamUsers = func(_id bson.ObjectID) ([]*model.User, error) {
		users := []*model.User{{
			ID:        _id,
			Password:  "River Song",
			FirstName: "Doctor",
			LastName:  "Who",
			Email:     "DoctoWho@Tardis.com",
			Role:      "ADMIN",
		}}

		return users, nil
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterAdminRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/admin/team/users/TestTeam", nil)

	token_string, _ := service.CreateToken(adminId, "ADMIN", adminId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 200 or 500, got %d", w.Code)
	}
}

func TestGetAllUsersFromTeam_WithMockedGetManyUsers(t *testing.T) {
	origGetTeamByName := service.GetTeamByName
	origGetManyUsers := repository.GetManyUsers
	defer func() {
		service.GetTeamByName = origGetTeamByName
		repository.GetManyUsers = origGetManyUsers
	}()

	teamId := bson.NewObjectID()
	userId := bson.NewObjectID()
	adminId := bson.NewObjectID()

	service.GetTeamByName = func(name string) (*model.Team, error) {
		return &model.Team{
			ID:   teamId,
			Name: name,
		}, nil
	}

	repository.GetManyUsers = func(users *[]*model.User, filter bson.D) error {
		*users = []*model.User{
			{
				ID:        userId,
				FirstName: "Doctor",
				LastName:  "Who",
				Email:     "Doctor@Tardis.com",
				Password:  "hashed",
				Role:      "USER",
				Team:      teamId,
			},
		}
		return nil
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterAdminRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/admin/team/users/TestTeam", nil)

	token_string, _ := service.CreateToken(adminId, "ADMIN", adminId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 200 or 500, got %d", w.Code)
	}
}
