package http_middleware

import (
	"log/slog"

	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewAuthMiddleware(deps *http_deps.Dependencies, logs *slog.Logger) *AuthMiddleware {
	return &AuthMiddleware{deps: deps, logs: logs}
}

func (m *AuthMiddleware) MiddlewareFunc() func(c *gin.Context) {
	return func(c *gin.Context) {	
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			m.logs.Error("Authorization header is required")
			m.logs.Error("Request missing Authorization header", "path", c.Request.URL.Path, "method", c.Request.Method)
			m.logs.Error("Request headers", "headers", c.GetHeader("Authorization"))			
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is required"})
			return
		}
		isValid, err := m.deps.App.Services.AuthService.ValidateAccessToken(tokenString)
		if err != nil || !isValid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}
		c.Next()
	}	
}