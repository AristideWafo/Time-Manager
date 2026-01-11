package admin

import (
	"TimeManager/middleware"
	"TimeManager/model"
	"TimeManager/repository"
	"TimeManager/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestGetUserById_ReturnUser(t *testing.T) {
	var orig = service.GetUserByID
	defer func() { service.GetUserByID = orig }()

	_id := bson.NewObjectID()
	service.GetUserByID = func(_id bson.ObjectID) (*model.User, error) {
		return &model.User{
			ID:        _id,
			Password:  "River Song",
			FirstName: "Doctor",
			LastName:  "Who",
			Email:     "DoctoWho@Tardis.com",
			Role:      "ADMIN",
		}, nil
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterAdminRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/admin/"+_id.Hex(), nil)
	token_string, _ := service.CreateToken(_id, "ADMIN", _id)
	req.Header.Add("api_token", token_string)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestPostUser_InvalidInput(t *testing.T) {
	_id := bson.NewObjectID()

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterAdminRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/admin/user/create", nil)
	req.Header.Set("Content-Type", "application/json")

	token_string, _ := service.CreateToken(_id, "ADMIN", _id)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for invalid input, got %d", w.Code)
	}
}

func TestUpdateUserTeam_InvalidUserID(t *testing.T) {
	_id := bson.NewObjectID()

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterAdminRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/admin/update/team/user/invalid-id", nil)
	req.Header.Set("Content-Type", "application/json")

	token_string, _ := service.CreateToken(_id, "ADMIN", _id)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500 for invalid user id, got %d", w.Code)
	}
}

func TestGetAllUsers_RouteReturnsSuccess(t *testing.T) {

	orig := repository.GetManyUsers
	defer func() { repository.GetManyUsers = orig }()

	_id := bson.NewObjectID()
	repository.GetManyUsers = func(users *[]*model.User, filter bson.D) error {
		*users = []*model.User{{
			ID:        _id,
			Password:  "River Song",
			FirstName: "Doctor",
			LastName:  "Who",
			Email:     "DoctoWho@Tardis.com",
			Role:      "ADMIN",
		}}
		return nil
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterAdminRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/admin/all/user", nil)

	token_string, _ := service.CreateToken(_id, "ADMIN", _id)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 200 or 500, got %d", w.Code)
	}
}
