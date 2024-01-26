package main

import (
	"fmt"

	bcrypt "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/data/crypto"
	http "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/data/http"
	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/data/postgres"
	docs "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/docs"
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
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
	crypto := bcrypt.NewBcrypt()
	userUseCase := usecases.NewUserUseCase(userRepo, crypto)
	userController := http.NewUserController(userUseCase)

	docs.SwaggerInfo.Title = "Orange Portfolio"
	docs.SwaggerInfo.Description = "This provide endpoints to create a portofolio manager."
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/"
	r := http.Router()
	r.POST("/users", userController.CreateUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
		return
	}
}
