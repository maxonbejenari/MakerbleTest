package test

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/maxonbejenari/testWebApp/middleware"
	"github.com/maxonbejenari/testWebApp/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

// helper to generate a valid token for testing
func getTestToken(role string) string {
	token, _ := utils.GenerateJWT("testuser", role)
	return token
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.AuthMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		c.String(200, "ok")
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.AuthMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		c.String(200, "ok")
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalidation")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleWare_ValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.AuthMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		role, _ := c.Get("role")
		c.String(200, role.(string))
	})

	token := getTestToken("doctor")
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer"+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "doctor", w.Body.String())
}
