package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	apiKey string
}

func NewAuthMiddleware(apiKey string) *AuthMiddleware {
	return &AuthMiddleware{apiKey}
}

func (m *AuthMiddleware) VerifyHeaderKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("x-api-key")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewDefaultResponse("Missing API Key"))
			return
		}

		if authHeader != m.apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewDefaultResponse("Invalid API Key"))
			return
		}

		c.Next()
	}
}
