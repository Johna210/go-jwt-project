package middleware

import (
	"fmt"
	"net/http"

	"os"

	helper "github.com/Johna210/go-jwt-project/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Auhtorization header provided"})
			c.Abort()
			return
		}

		fmt.Fprintln(os.Stdout, []any{"Here................"}...)
		claims, err := helper.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.UID)
		c.Set("user_type", claims.User_type)
		c.Next()
	}
}
