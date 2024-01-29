package http

import (
	"net/http"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectController struct {
	projectUseCase usecases.ProjectUseCase
	userUseCase    usecases.UserUseCase
	token          usecases.Token
}

func NewProjectController(projectUseCase usecases.ProjectUseCase, userUseCase usecases.UserUseCase, token usecases.Token) ProjectController {
	return ProjectController{
		projectUseCase: projectUseCase,
		userUseCase:    userUseCase,
		token:          token,
	}
}

func (p ProjectController) CreateProject(c *gin.Context) {
	var project entities.Project
	err := c.ShouldBind(&project)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	userId, err := p.token.GetPayload(token)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

	newProject, err := entities.NewProject(project)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newProject.ID = uuid.New().String()
	err = p.projectUseCase.CreateProject(newProject, userId)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Projeto criado com sucesso"})

}
