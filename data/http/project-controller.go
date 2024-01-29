package http

import (
	"net/http"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectController struct {
	projectUseCase usecases.ProjectUseCases
	userUseCase    usecases.UserUseCases
	token          usecases.Token
}

func NewProjectController(projectUseCase usecases.ProjectUseCases, userUseCase usecases.UserUseCases, token usecases.Token) ProjectController {
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

func (p ProjectController) FindAllProjects(c *gin.Context) {
	projects, err := p.projectUseCase.FindAllProjects()
	var projectsDTO []ProjectDTO
	for _, project := range projects {
		projectsDTO = append(projectsDTO, projectAggregateToDTO(project))
	}
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, projectsDTO)
}

func (p ProjectController) FindProjectsByUserId(c *gin.Context) {
	userId := c.Param("userId")
	projects, err := p.projectUseCase.FindProjectsByUserId(userId)
	var projectsDTO []ProjectDTO
	for _, project := range projects {
		projectsDTO = append(projectsDTO, projectAggregateToDTO(project))
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projectsDTO)
}

func (p ProjectController) UpdateProject(c *gin.Context) {
	var project entities.Project
	err := c.ShouldBind(&project)
	if err != nil {
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

	newProject, err := entities.NewProject(project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newProject.ID = c.Param("projectId")
	err = p.projectUseCase.UpdateProject(newProject, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Projeto atualizado com sucesso"})
}

func (p ProjectController) DeleteProject(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	userId, err := p.token.GetPayload(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

	projectId := c.Param("projectId")
	err = p.projectUseCase.DeleteProject(projectId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Projeto deletado com sucesso"})
}
