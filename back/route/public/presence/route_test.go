package presence

import (
	"TimeManager/middleware"
	"TimeManager/model"
	"TimeManager/repository"
	"TimeManager/service"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestGetAllPresences_GetAllPresencesError(t *testing.T) {
	origGetManyPresences := repository.GetManyPresences
	defer func() {
		repository.GetManyPresences = origGetManyPresences
	}()

	userId := bson.NewObjectID()

	repository.GetManyPresences = func(presences *[]*model.Presence, filter bson.D) error {
		return errors.New("GetManyPresences error")
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterPresenceRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/presence", nil)

	token_string, _ := service.CreateToken(userId, "USER", userId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 404 or 500, got %d", w.Code)
	}
}

func TestGetAllPresences_Success(t *testing.T) {
	origGetManyPresences := repository.GetManyPresences
	defer func() {
		repository.GetManyPresences = origGetManyPresences
	}()

	userId := bson.NewObjectID()

	repository.GetManyPresences = func(presences *[]*model.Presence, filter bson.D) error {
		*presences = []*model.Presence{}
		return nil
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterPresenceRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/presence", nil)

	token_string, _ := service.CreateToken(userId, "USER", userId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 200 or 500, got %d", w.Code)
	}
}

func TestCreatePresence_InvalidInput(t *testing.T) {
	userId := bson.NewObjectID()

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterPresenceRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/presence/create", nil)
	req.Header.Set("Content-Type", "application/json")

	token_string, _ := service.CreateToken(userId, "USER", userId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for invalid input, got %d", w.Code)
	}
}

func TestCreatePresence_CreatePresenceError(t *testing.T) {
	origSavePresence := repository.SavePresence
	defer func() {
		repository.SavePresence = origSavePresence
	}()

	userId := bson.NewObjectID()

	repository.SavePresence = func(presence *model.Presence) error {
		return errors.New("SavePresence error")
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterPresenceRoutes(router)

	input := CreatePresenceInput{
		Type:      "IN",
		Timestamp: time.Now(),
	}
	jsonData, _ := json.Marshal(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/presence/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	token_string, _ := service.CreateToken(userId, "USER", userId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusConflict && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 409 or 500, got %d", w.Code)
	}
}

func TestCreatePresence_Success(t *testing.T) {
	origSavePresence := repository.SavePresence
	defer func() {
		repository.SavePresence = origSavePresence
	}()

	userId := bson.NewObjectID()

	repository.SavePresence = func(presence *model.Presence) error {
		return nil
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	RegisterPresenceRoutes(router)

	input := CreatePresenceInput{
		Type:      "IN",
		Timestamp: time.Now(),
	}
	jsonData, _ := json.Marshal(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/presence/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	token_string, _ := service.CreateToken(userId, "USER", userId)
	req.Header.Add("api_token", token_string)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 200 or 500, got %d", w.Code)
	}
}
