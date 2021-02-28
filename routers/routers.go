package routers

import (
	"home/controllers"
	"home/custom"
	"home/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouters(api *gin.Engine, cs custom.Custom) error {
	api.Use(gin.Recovery())

	user := controllers.NewUserController(cs)
	state := controllers.NewUserStateController(cs)

	api.POST(controllers.Login, user.Login)   // /api/public/login
	api.POST(controllers.SignUp, user.Signup) // /api/public/signup

	protected := api.Group(controllers.Protected)
	protected.Use(middlewares.IsAuthorized(cs))

	protected.POST(controllers.Logout, user.Logout)               // /api/protected/logout
	protected.POST(controllers.AddUserState, state.AddUserState)  // /api/protected/addstate
	protected.POST(controllers.UpdateMyState, user.UpdateMyState) // /api/protected/updateState
	protected.GET(controllers.GetStatus, func(c *gin.Context) {   // /api/protected/mystatus
		c.JSON(200, gin.H{"status": "authenticated"})
	})
	protected.GET(controllers.AllUserState, state.AllUserStates) // /api/protected/allstates

	return nil
}

/*
{
    "Username" : "aqzhol@gmail.com",
    "Password" : "123"
}
*/
