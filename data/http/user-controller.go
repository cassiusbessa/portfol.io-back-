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

// @Summary Create a new user
// @Description Create a new user with the provided data
// @Accept json
// @Produce json
// @Param user body object true "User object to be created"
// @Param email body string true "Email address" format(email) example(email@email.com)
// @Param password body string true "User Password (8-32 characters, at least one uppercase letter, one lowercase letter, one number and one special character)" example(Password123!)
// @Param fullName body string true "Full name of the user" example("Só Mais Silva")
// @Success 201 {string} string "User created successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [post]
func (u UserController) CreateUser(c *gin.Context) {
	var user entities.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server error"})
		return
	}

	newUser, err := entities.NewUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = uuid.New().String()

	err = u.userUseCase.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusContinue, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
