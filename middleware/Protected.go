package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Protected() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		logged := session.Get("logged")

		if logged != true {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User not authenticated"})
			return
		}

		c.Next()
	}
}
