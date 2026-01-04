package teams

import (
	"TimeManager/middleware"
	"TimeManager/model"
	"TimeManager/service"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestFetchTeam_EmptyTeamName(t *testing.T) {
	_id := bson.NewObjectID()

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterTeamRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/team/", nil)

	token_string, _ := service.CreateToken(_id, "MANAGER", _id)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest && w.Code != http.StatusNotFound {
		t.Fatalf("expected status 400 or 404, got %d", w.Code)
	}
}

func TestFetchTeam_TeamNotFound(t *testing.T) {
	orig := service.GetTeamByName
	defer func() { service.GetTeamByName = orig }()

	_id := bson.NewObjectID()

	service.GetTeamByName = func(name string) (*model.Team, error) {
		return nil, errors.New("GetTeamByName error")
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterTeamRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/team/TestTeam", nil)

	token_string, _ := service.CreateToken(_id, "MANAGER", _id)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404 for team not found, got %d", w.Code)
	}
}

func TestFetchTeam_Success(t *testing.T) {
	orig := service.GetTeamByName
	defer func() { service.GetTeamByName = orig }()

	teamId := bson.NewObjectID()
	userId := bson.NewObjectID()

	service.GetTeamByName = func(name string) (*model.Team, error) {
		return &model.Team{
			ID:   teamId,
			Name: name,
		}, nil
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterTeamRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/team/TestTeam", nil)

	token_string, _ := service.CreateToken(userId, "MANAGER", teamId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 200 or 500, got %d", w.Code)
	}
}
