package controllers

import (
	"home/custom"
	service "home/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	cs      custom.Custom
	service *service.UserService
}

func NewUserController(custom custom.Custom) *UserController {
	return &UserController{
		cs:      custom,
		service: service.NewUserService(custom),
	}
}

func (controller *UserController) UpdateMyState(c *gin.Context) {
	userIDInterf, _ := c.Get("userID")
	userID, _ := userIDInterf.(uint)
	// username, _ := c.Get("username")

	// Get User model
	errUser, user := controller.service.FindByID(userID)
	if errUser != nil {
		c.JSON(500, gin.H{"status": "Invalid userID"})
		c.Abort()
		return
	}

	// Get State IDs
	stateIDStr := c.PostForm("stateID")
	stateID, err := strconv.Atoi(stateIDStr)
	if err != nil {
		c.JSON(400, gin.H{"status": "Invalid stateID"})
		c.Abort()
		return
	}

	// Update state
	rep := controller.cs.GetRepository()
	err = user.UpdateState(rep, uint(stateID))
	if err != nil {
		c.JSON(500, gin.H{"status": "Can not update state"})
		c.Abort()
		return
	}

	// Success
	c.JSON(200, gin.H{"status": "Successfully updated"})
}
