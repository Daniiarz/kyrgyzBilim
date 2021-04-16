package utils

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/service"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		accessToken := strings.Split(header, " ")
		user, err := service.ValidateToken(accessToken[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)
	}
}
