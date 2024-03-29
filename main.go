package main

import (
	"fmt"

	bcrypt "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/data/crypto"
	http "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/data/http"
	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/data/postgres"
	jwt "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/data/token"
	docs "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/docs"
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	postgresDb, err := postgres.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = postgres.CreateDb(postgresDb)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer postgresDb.Close()

	userRepo := postgres.NewUserRepository(postgresDb)
	projectRepo := postgres.NewProjectRepository(postgresDb)

	crypto := bcrypt.NewBcrypt()
	token := jwt.NewJWT("secret")

	userUseCase := usecases.NewUserUseCase(userRepo)
	projectUseCase := usecases.NewProjectUseCase(projectRepo, userRepo)
	tagUseCase := usecases.NewTagUseCase(postgres.NewTagRepository(postgresDb))

	googleService := services.NewGoogleService()

	userController := http.NewUserController(userUseCase, crypto, token, *googleService)
	projectController := http.NewProjectController(projectUseCase, userUseCase, token)
	tagController := http.NewTagController(tagUseCase, token)

	docs.SwaggerInfo.Title = "Orange Portfolio"
	docs.SwaggerInfo.Description = "This provide endpoints to create a portofolio manager."
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/"
	r := http.Router()

	r.POST("/users", userController.CreateUser)
	r.POST("/login", userController.Login)
	r.POST("/google-login", userController.GoogleLogin)
	r.GET("/me", userController.Me)

	r.POST("/projects", projectController.CreateProject)
	r.GET("/projects", projectController.FindAllProjects)
	r.GET("/projects/users/:userId", projectController.FindProjectsByUserId)
	r.PUT("/projects/:projectId", projectController.UpdateProject)
	r.DELETE("/projects/:projectId", projectController.DeleteProject)

	r.GET("/tags", tagController.FindAllTags)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
		return
	}
}
