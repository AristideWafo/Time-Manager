package middleware

import (
	"net/http/httptest"
	"testing"

	"TimeManager/service"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

func TestBasePrivileges_MissingClaims(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/x", nil)

	BasePrivilegesMiddleWare(c, []string{"ADMIN"})

	if !c.IsAborted() {
		t.Fatalf("expected abort when claims missing")
	}
}

func TestBasePrivileges_WrongClaimsType(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/x", nil)
	c.Set("claims", "not-claims")

	BasePrivilegesMiddleWare(c, []string{"ADMIN"})

	if !c.IsAborted() {
		t.Fatalf("expected abort when claims have wrong type")
	}
}

func TestBasePrivileges_InsufficientPrivilege(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/x", nil)

	c.Set("claims", service.TokenClaims{RegisteredClaims: jwt.RegisteredClaims{Audience: jwt.ClaimStrings{"EMPLOYEE"}, Subject: "sub"}, Team: "t1"})

	BasePrivilegesMiddleWare(c, []string{"MANAGER"})

	if !c.IsAborted() {
		t.Fatalf("expected abort when privilege insufficient")
	}
}

func TestBasePrivileges_Allowed(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/x", nil)

	c.Set("claims", service.TokenClaims{RegisteredClaims: jwt.RegisteredClaims{Audience: jwt.ClaimStrings{"MANAGER"}, Subject: "sub"}, Team: "t1"})

	BasePrivilegesMiddleWare(c, []string{"MANAGER", "ADMIN"})

	if c.IsAborted() {
		t.Fatalf("did not expect abort when privilege sufficient")
	}
}
