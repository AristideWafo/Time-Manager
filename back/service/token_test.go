package service

import (
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestCreateAndValidateToken_Success(t *testing.T) {

	os.Setenv("SECRET_KEY", "test-secret")

	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	team, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439022")
	role := "MANAGER"

	tok, err := CreateToken(id, role, team)
	if err != nil {
		t.Fatalf("CreateToken error: %v", err)
	}

	claims, err := ValidateToken(tok)
	if err != nil {
		t.Fatalf("ValidateToken error: %v", err)
	}

	if claims.Team != team.Hex() {
		t.Fatalf("expected team %s, got %s", team.Hex(), claims.Team)
	}
	if claims.Subject != id.Hex() {
		t.Fatalf("expected subject %s, got %s", id.Hex(), claims.Subject)
	}
}

func TestValidateToken_Invalid(t *testing.T) {
	os.Setenv("SECRET_KEY", "test-secret")
	_, err := ValidateToken("not-a-token")
	if err == nil {
		t.Fatalf("expected error for invalid token")
	}
}

func TestValidateToken_Expired(t *testing.T) {
	os.Setenv("SECRET_KEY", "test-secret")

	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	team := "teamid"

	claims := TokenClaims{
		Team: team,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   id.Hex(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}

	_, err = ValidateToken(tokenString)
	if err == nil {
		t.Fatalf("expected error for expired token")
	}
}

func TestDecryptClaimsFromContext(t *testing.T) {
	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)

	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	team, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439022")

	c.Set("claims", TokenClaims{Team: team.Hex(), RegisteredClaims: jwt.RegisteredClaims{Subject: id.Hex()}})

	gotID, err := DecryptIDFromContextClaim(c)
	if err != nil {
		t.Fatalf("DecryptIDFromContextClaim error: %v", err)
	}
	if gotID != id {
		t.Fatalf("expected id %v, got %v", id, gotID)
	}

	gotTeam, err := DecryptTeamFromContextClaim(c)
	if err != nil {
		t.Fatalf("DecryptTeamFromContextClaim error: %v", err)
	}
	if gotTeam != team {
		t.Fatalf("expected team %v, got %v", team, gotTeam)
	}
}

func TestDecryptClaims_MissingOrWrong(t *testing.T) {
	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)

	_, err := DecryptIDFromContextClaim(c)
	if err == nil {
		t.Fatalf("expected error when claims missing")
	}

	c.Set("claims", "not-claims")
	_, err = DecryptTeamFromContextClaim(c)
	if err == nil {
		t.Fatalf("expected error when claims have wrong type")
	}
}
