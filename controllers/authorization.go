package controllers

import (
	"home/models"
	"home/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func (controller *UserController) Login(c *gin.Context) {
	var user models.User

	// JSON validation
	jsonErr := c.ShouldBind(&user)
	if jsonErr != nil {
		c.JSON(400, gin.H{"status": "Invalid JSON"})
		c.Abort()
		return
	}

	// Username validation
	dbErr, dbUser := controller.service.FindByUsername(user.Username)
	if dbErr != nil {
		c.JSON(401, gin.H{"status": "Incorrect login or password"})
		c.Abort()
		return
	}

	// Password validation
	incorrectErr := dbUser.CheckPassword(user.Password)
	if incorrectErr != nil {
		c.JSON(401, gin.H{"status": "Incorrect login or password"})
		c.Abort()
		return
	}

	// Generate JWT token
	errToken, jwtToken := token.GenerateToken(user.Username)
	if errToken != nil {
		c.JSON(500, gin.H{"status": "Invalid token generation"})
		c.Abort()
		return
	}

	// Save token in Redis
	rds := controller.cs.GetRedisClient()
	rds.SaveAuth(jwtToken, uint(dbUser.ID))

	// Success Response
	c.JSON(200, gin.H{"token": jwtToken})
}

func (controller *UserController) Logout(c *gin.Context) {
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

	// Remove from redis
	rds := controller.cs.GetRedisClient()
	removeErr := rds.RemoveAuth(clientToken)
	if removeErr != nil {
		c.JSON(500, "Could not remove auth token")
		c.Abort()
		return
	}

	// Success logout
	c.JSON(200, "You are successfully logged out!")
}

func (controller *UserController) Signup(c *gin.Context) {
	var user models.User

	// Json validation
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{"status": "invalid JSON"})
		c.Abort()
		return
	}

	// Hashing password
	err = user.HashPassword()
	if err != nil {
		c.JSON(500, gin.H{"status": "invalid password"})
		c.Abort()
		return
	}

	// Saving user record
	rep := controller.cs.GetRepository()
	err = user.CreateUserRecord(rep)
	if err != nil {
		c.JSON(500, gin.H{"status": "can not save user"})
		c.Abort()
		return
	}

	// Success case
	c.JSON(200, user)
}
