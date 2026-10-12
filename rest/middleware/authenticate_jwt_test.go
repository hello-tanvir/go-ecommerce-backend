package middleware

import (
	"ecommerce/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticateJwtReturnsUnauthorizedForMalformedHeader(t *testing.T) {
	m := NewMeddleWares(&config.Config{JwtSecretkey: "test-secret"})

	handler := m.AuthenticateJwt(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("middleware should stop before calling the next handler")
	}))

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	req.Header.Set("Authorization", "Bearer")
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	if res.Code != http.StatusUnauthorized {
		t.Fatalf("expected %d but got %d", http.StatusUnauthorized, res.Code)
	}
}
