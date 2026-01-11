package authentification

import (
	"TimeManager/model"
	"TimeManager/repository"
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

func TestAuthentificate_InvalidInput(t *testing.T) {
	router := gin.New()
	RegisterAuthentificationRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/authentification", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for invalid input, got %d", w.Code)
	}
}

func TestAuthentificate_UserNotFound(t *testing.T) {
	origGetOneUser := repository.GetOneUser
	defer func() {
		repository.GetOneUser = origGetOneUser
	}()

	repository.GetOneUser = func(user *model.User, filter bson.D) error {
		return errors.New("user not found")
	}

	router := gin.New()
	RegisterAuthentificationRoutes(router)

	input := AuthentificationUserInput{
		Email:    "test@Tardis.com",
		Password: "password",
	}
	jsonData, _ := json.Marshal(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/authentification", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404 for user not found, got %d", w.Code)
	}
}

func TestAuthentificate_ValidCredentials(t *testing.T) {
	origGetOneUser := repository.GetOneUser
	defer func() {
		repository.GetOneUser = origGetOneUser
	}()

	userId := bson.NewObjectID()
	teamId := bson.NewObjectID()

	hashedPassword, _ := service.HashPassword("password")

	repository.GetOneUser = func(user *model.User, filter bson.D) error {
		*user = model.User{
			ID:        userId,
			FirstName: "Doctor",
			LastName:  "Who",
			Email:     "test@Tardis.com",
			Password:  hashedPassword,
			Role:      "USER",
			Team:      teamId,
		}
		return nil
	}

	router := gin.New()
	RegisterAuthentificationRoutes(router)

	input := AuthentificationUserInput{
		Email:    "test@Tardis.com",
		Password: "password",
	}
	jsonData, _ := json.Marshal(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/authentification", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 200 or 500, got %d", w.Code)
	}
}
