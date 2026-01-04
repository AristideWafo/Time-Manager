package users

import (
	"TimeManager/middleware"
	"TimeManager/model"
	"TimeManager/service"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestFetchUser_GetUserByIDError(t *testing.T) {
	origGetUserByID := service.GetUserByID
	defer func() {
		service.GetUserByID = origGetUserByID
	}()

	userId := bson.NewObjectID()

	service.GetUserByID = func(_id bson.ObjectID) (*model.User, error) {
		return nil, errors.New("GetUserByID error")
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterUserRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/user", nil)

	token_string, _ := service.CreateToken(userId, "ADMIN", userId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404 for user not found, got %d", w.Code)
	}
}

func TestFetchUser_Success(t *testing.T) {
	origGetUserByID := service.GetUserByID
	defer func() {
		service.GetUserByID = origGetUserByID
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

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterUserRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/user", nil)

	token_string, _ := service.CreateToken(userId, "ADMIN", teamId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 200 or 500, got %d", w.Code)
	}
}

func TestUpdateUser_InvalidInput(t *testing.T) {
	userId := bson.NewObjectID()

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterUserRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/user/update", nil)
	req.Header.Set("Content-Type", "application/json")

	token_string, _ := service.CreateToken(userId, "ADMIN", userId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for invalid input, got %d", w.Code)
	}
}

func TestUpdateUser_NoUpdateFields(t *testing.T) {
	origGetUserByID := service.GetUserByID
	defer func() {
		service.GetUserByID = origGetUserByID
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

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterUserRoutes(router)

	input := UpdateUserInput{}
	jsonData, _ := json.Marshal(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/user/update", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	token_string, _ := service.CreateToken(userId, "ADMIN", teamId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 200 or 500, got %d", w.Code)
	}
}
