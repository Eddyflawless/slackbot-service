package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractToken(r *http.Request) string {

	bearerToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func Authenticate() gin.HandlerFunc {

	return func(c *gin.Context) {

		accessToken := extractToken(c.Request)

		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		}

		if accessToken != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// logic here..
		c.Next()
	}
}
