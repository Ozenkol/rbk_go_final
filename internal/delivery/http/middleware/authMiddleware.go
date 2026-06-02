package http_middleware

import (
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	deps *http_deps.Dependencies
}

func NewAuthMiddleware(deps *http_deps.Dependencies) *AuthMiddleware {
	return &AuthMiddleware{deps: deps}
}

func (m *AuthMiddleware) MiddlewareFunc() func(c *gin.Context) {
	return func(c *gin.Context) {	
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
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