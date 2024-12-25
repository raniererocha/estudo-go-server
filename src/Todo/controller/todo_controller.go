package todo_controller

import (
	"net/http"
	Dto "server/src/Todo/dto"
	"server/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var newTodo Dto.CreateTodoDto

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	if err := utils.Validate(newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
}

func ListTodo(c *gin.Context) {

}
