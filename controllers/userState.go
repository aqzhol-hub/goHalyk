package controllers

import (
	"home/custom"
	"home/models"
	service "home/services"

	"github.com/gin-gonic/gin"
)

type UserStateController struct {
	cs      custom.Custom
	service *service.UserStateService
}

func NewUserStateController(custom custom.Custom) *UserStateController {
	return &UserStateController{
		cs:      custom,
		service: service.NewUserStateService(custom),
	}
}

func (controller *UserStateController) AddUserState(c *gin.Context) {
	var userState models.UserState

	// JSON validation
	jsonErr := c.ShouldBind(&userState)
	if jsonErr != nil {
		c.JSON(400, gin.H{"status": "Invalid JSON"})
		c.Abort()
		return
	}

	// Saving userState record
	rep := controller.cs.GetRepository()
	err := userState.CreateUserStateRecord(rep)
	if err != nil {
		c.JSON(500, gin.H{"status": "can not save user state"})
		c.Abort()
		return
	}

	// Success case
	c.JSON(200, userState)
}

func (controller *UserStateController) AllUserStates(c *gin.Context) {
	err, userStates := controller.service.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"status": "can not get user state"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"result": userStates})
}
