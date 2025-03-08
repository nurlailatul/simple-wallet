package middleware

import (
	"net/http"

	"simple-wallet/config"

	"github.com/gin-gonic/gin"
)

type SwaggerAuthMiddleware struct {
	apiDocKey string
}

func NewSwaggerAuthMiddleware(swagger *config.SwaggerConfiguration) *SwaggerAuthMiddleware {
	return &SwaggerAuthMiddleware{swagger.ApiKey}
}

func (m *SwaggerAuthMiddleware) VerifyHeaderKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("x-api-key")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewDefaultResponse("Missing API Key"))
			return
		}

		if authHeader != m.apiDocKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewDefaultResponse("Invalid API Key"))
			return
		}

		c.Next()
	}
}
