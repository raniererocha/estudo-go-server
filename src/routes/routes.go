package routes

import (
	todo_controller "server/src/Todo/controller"
	user_controller "server/src/User/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	// USERS
	userGroup := app.Group("/users")
	{
		userGroup.POST("/", user_controller.CreateUser)
		userGroup.GET("/", user_controller.ListUsers)
	}

	// TODOS
	todoGroup := app.Group("/todos")
	{
		todoGroup.POST("/", todo_controller.CreateTodo)
		todoGroup.GET("/", todo_controller.ListTodo)
	}

}
