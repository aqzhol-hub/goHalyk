package middlewares

import (
	"home/custom"
	"home/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsAuthorized(cs custom.Custom) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Check for header exists
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}

		// Check for token
		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}

		// Token validation
		err, claims := token.ValidateToken(clientToken)
		if err != nil {
			c.JSON(401, err)
			c.Abort()
			return
		}
		c.Set("username", claims.Username)

		// Search from Redis
		rds := cs.GetRedisClient()
		userID, errRedis := rds.GetAuth(clientToken)
		if errRedis != nil || userID == 0 {
			c.JSON(401, "Not authenticated")
			c.Abort()
			return
		}

		// Success
		c.Set("userID", userID)
		c.Next()
	}
}
