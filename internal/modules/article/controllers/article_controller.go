package controllers

import (
	"net/http"
	"strconv"
	ArticlesService "v2/internal/modules/article/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articlesService ArticlesService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articlesService: ArticlesService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid ID"})
		return
	}

	article, err := controller.articlesService.Find(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"article": article})

}
