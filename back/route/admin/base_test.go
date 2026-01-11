package admin

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterAdminRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	RegisterAdminRoutes(router)

	routes := router.Routes()
	if len(routes) == 0 {
		t.Fatalf("expected routes to be registered, got 0")
	}

	expectedPaths := []string{
		"/api/admin/:id",
		"/api/admin/user/create",
		"/api/admin/update/team/user/:id",
		"/api/admin/all/user",
		"/api/admin/all/team",
		"/api/admin/team/users/:name",
		"/api/admin/team/create",
		"/api/admin/team/update",
	}

	found := 0
	for _, expectedPath := range expectedPaths {
		for _, route := range routes {
			if route.Path == expectedPath {
				found++
				break
			}
		}
	}

	if found != len(expectedPaths) {
		t.Fatalf("expected %d routes to be registered, found %d", len(expectedPaths), found)
	}
}

func TestRegisterAdminRoutes_HasMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	RegisterAdminRoutes(router)

	routes := router.Routes()
	hasAdminGroup := false
	for _, route := range routes {
		if route.Path == "/api/admin/:id" {
			hasAdminGroup = true
			break
		}
	}

	if !hasAdminGroup {
		t.Fatalf("expected admin routes to be registered")
	}
}

func TestRegisterAdminRoutes_VerifyPaths(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	RegisterAdminRoutes(router)

	routes := router.Routes()

	pathCount := 0
	for _, route := range routes {
		if route.Path == "/api/admin/:id" ||
			route.Path == "/api/admin/user/create" ||
			route.Path == "/api/admin/update/team/user/:id" ||
			route.Path == "/api/admin/all/user" ||
			route.Path == "/api/admin/all/team" ||
			route.Path == "/api/admin/team/users/:name" ||
			route.Path == "/api/admin/team/create" ||
			route.Path == "/api/admin/team/update" {
			pathCount++
		}
	}

	if pathCount < 8 {
		t.Fatalf("expected at least 8 admin paths to be registered, found %d", pathCount)
	}
}

func TestAdminRoutes_HTTPMethods(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	RegisterAdminRoutes(router)

	routes := router.Routes()
	methodMap := make(map[string]string)
	for _, route := range routes {
		methodMap[route.Path] = route.Method
	}

	expectedMethods := map[string]string{
		"/api/admin/:id":                  "GET",
		"/api/admin/user/create":          "POST",
		"/api/admin/update/team/user/:id": "PUT",
		"/api/admin/all/user":             "GET",
		"/api/admin/all/team":             "GET",
		"/api/admin/team/users/:name":     "GET",
		"/api/admin/team/create":          "POST",
		"/api/admin/team/update":          "PUT",
	}

	for path, expectedMethod := range expectedMethods {
		if method, exists := methodMap[path]; exists {
			if method != expectedMethod {
				t.Fatalf("expected %s to have method %s, got %s", path, expectedMethod, method)
			}
		}
	}
}
