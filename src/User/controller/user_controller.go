package user_controller

import (
	"net/http"
	Dto "server/src/User/dto"
	repository "server/src/User/repository"
	"server/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var newUser Dto.CreateUserDTO

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !repository.IsEmailUnique(newUser.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "O email informado já está cadastrado"})
		return
	}
	if err := utils.Validate(newUser); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}
	repository.StoreUser(newUser)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Criado com sucesso",
	})
}

func ListUsers(c *gin.Context) {
	users := repository.FindUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
