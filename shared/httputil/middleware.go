package httputil

import (
	"ddd/shared/errors"
	"ddd/shared/logger"

	"github.com/gin-gonic/gin"
)

func OrgIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		orgID := c.GetHeader("X-Organization-ID")
		if orgID == "" {
			NewErrorResponse(c, errors.NewBadRequest("Organization ID is required", nil))
			c.Abort()
			return
		}

		c.Set("orgID", orgID)
		c.Next()
	}
}

// func LoggerMiddleware(logger *logger.Logger) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		logger := logger.Info(c, "gin middleware", map[string]interface{}{})

// 		// map[string]interface{}{
// 		// 	"path":   c.Request.URL.Path,
// 		// 	"method": c.Request.Method,
// 		// })

// 		c.Set("logger", logger)
// 		c.Next()
// 	}
// }

func RecoveryMiddleware(logger logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(c, errors.NewInternal("err", nil), "panic recovered:", map[string]interface{}{})
				NewErrorResponse(c, errors.NewInternal("Internal server error", nil))
				c.Abort()
			}
		}()
		c.Next()
	}
}
