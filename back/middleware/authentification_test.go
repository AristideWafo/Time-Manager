package middleware

import (
	"net/http/httptest"
	"os"
	"testing"

	"TimeManager/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestAuthMiddleware_NoToken(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/protected", nil)

	handler := AuthMiddleware()
	handler(c)

	if !c.IsAborted() {
		t.Fatalf("expected middleware to abort when no token provided")
	}
	if w.Code != 400 {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
}

func TestAuthMiddleware_AllowedPath(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/swagger/index.html", nil)

	handler := AuthMiddleware()
	handler(c)

	if c.IsAborted() {
		t.Fatalf("did not expect middleware to abort for allowed path")
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("GET", "/api/protected", nil)
	req.Header.Set("api_token", "bad-token")
	c.Request = req

	handler := AuthMiddleware()
	handler(c)

	if !c.IsAborted() {
		t.Fatalf("expected middleware to abort for invalid token")
	}
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	os.Setenv("SECRET_KEY", "test-secret")

	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	team, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439022")
	tok, err := service.CreateToken(id, "MANAGER", team)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("GET", "/api/protected", nil)
	req.Header.Set("api_token", tok)
	c.Request = req

	handler := AuthMiddleware()
	handler(c)

	if c.IsAborted() {
		t.Fatalf("did not expect middleware to abort for valid token")
	}

	if _, ok := c.Get("claims"); !ok {
		t.Fatalf("expected claims to be set in context")
	}
}
