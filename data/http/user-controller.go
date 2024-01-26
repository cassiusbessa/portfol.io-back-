package http

import (
	"net/http"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userUseCase usecases.UserUseCase
}

func NewUserController(userUseCase usecases.UserUseCase) UserController {
	return UserController{
		userUseCase: userUseCase,
	}
}

// @Param user body CreateUserDTO true "User object to be created"
// @Router /users [post]
// @Success 201 {object} Response "User created successfully" {"message": "User created successfully"}
// @Failure 400 {object} Response "Bad Request" {"message": "Bad Request"}
// @Failure 500 {object} Response "Internal Server Error" {"message": "Internal Server Error"}
func (u UserController) CreateUser(c *gin.Context) {
	var user entities.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server error"})
		return
	}

	newUser, err := entities.NewUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newUser.ID = uuid.New().String()

	err = u.userUseCase.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
