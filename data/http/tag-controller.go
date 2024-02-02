package http

import (
	"net/http"

	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagUseCases usecases.TagUseCases
	token       usecases.Token
}

func NewTagController(tagUseCases usecases.TagUseCases, token usecases.Token) TagController {
	return TagController{
		tagUseCases: tagUseCases,
		token:       token,
	}
}

func (t TagController) FindAllTags(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	_, err := t.token.GenerateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	tags, err := t.tagUseCases.FindAllTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, tags)
}
