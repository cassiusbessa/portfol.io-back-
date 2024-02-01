package http

import (
	"fmt"
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

// @Security ApiKeyAuth
// @Summary Create a new project
// @Description Create a new project with the provided information
// @Accept json
// @Produce json
// @Param project body CreateProjectDTO true "Project object to be created"
// @Success 201 {object} Response "Project created successfully" {"message": "Projeto criado com sucesso"}
// @Failure 400 {object} Response "Bad Request" {"message": "Bad Request"}
// @Failure 401 {object} Response "Unauthorized" {"message": "Unauthorized"}
// @Failure 500 {object} Response "Internal Server Error" {"message": "Internal Server Error"}
// @Router /projects [post]
func (p ProjectController) CreateProject(c *gin.Context) {
	var projectDTO CreateProjectDTO
	err := c.ShouldBind(&projectDTO)
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

	var project entities.Project
	project.Name = projectDTO.Name
	project.Description = projectDTO.Description
	project.Image = projectDTO.Image

	newProject, err := entities.NewProject(project)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newProject.ID = uuid.New().String()
	err = p.projectUseCase.CreateProject(newProject, userId, projectDTO.Tags)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Projeto criado com sucesso"})
}

// @Security ApiKeyAuth
// @Summary Get all projects
// @Description Get all projects of all users
// @Produce json
// @Success 200 {object} ProjectDTO "List of projects"
// @Failure 401 {object} Response "Unauthorized" {"message": "Unauthorized"}
// @Failure 500 {object} Response "Internal Server Error" {"message": "Internal Server Error"}
// @Router /projects [get]
func (p ProjectController) FindAllProjects(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	_, err := p.token.GetPayload(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

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

// @Security ApiKeyAuth
// @Summary Get all projects of a user
// @Description Get all projects of a user
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} ProjectDTO "List of projects"
// @Failure 401 {object} Response "Unauthorized" {"message": "Unauthorized"}
// @Failure 500 {object} Response "Internal Server Error" {"message": "Internal Server Error"}
// @Router /projects/users/{userId} [get]
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

// @Security ApiKeyAuth
// @Summary Update a project
// @Description Update a project with the provided information
// @Accept json
// @Produce json
// @Param projectId path string true "Project ID"
// @Param project body CreateProjectDTO true "Project object to be updated"
// @Success 200 {object} Response "Project updated successfully" {"message": "Projeto atualizado com sucesso"}
// @Failure 400 {object} Response "Bad Request" {"message": "Bad Request"}
// @Failure 401 {object} Response "Unauthorized" {"message": "Unauthorized"}
// @Failure 500 {object} Response "Internal Server Error" {"message": "Internal Server Error"}
// @Router /projects/{projectId} [put]
func (p ProjectController) UpdateProject(c *gin.Context) {
	var projectDTO CreateProjectDTO
	err := c.ShouldBind(&projectDTO)
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

	projectId := c.Param("projectId")
	var project entities.Project
	project.Name = projectDTO.Name
	project.Description = projectDTO.Description
	project.Image = projectDTO.Image

	newProject, err := entities.NewProject(project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newProject.ID = projectId
	fmt.Println(projectDTO.Tags)
	err = p.projectUseCase.UpdateProject(newProject, userId, projectDTO.Tags)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Projeto atualizado com sucesso"})
}

// @Security ApiKeyAuth
// @Summary Delete a project
// @Description Delete a project with the provided information
// @Produce json
// @Param projectId path string true "Project ID"
// @Success 200 {object} Response "Project deleted successfully" {"message": "Projeto deletado com sucesso"}
// @Failure 401 {object} Response "Unauthorized" {"message": "Unauthorized"}
// @Failure 500 {object} Response "Internal Server Error" {"message": "Internal Server Error"}
// @Router /projects/{projectId} [delete]
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
