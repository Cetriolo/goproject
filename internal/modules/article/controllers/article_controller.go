package controllers

import (
	"net/http"
	"strconv"
	ArticlesService "v2/internal/modules/article/services"
	"v2/pkg/html"

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
		html.Render(c, http.StatusInternalServerError, "templates/errors/500", gin.H{"title": "Internal Server Error", "message": "Invalid ID"})
		return
	}

	article, err := controller.articlesService.Find(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/404", gin.H{"title": "Page not found", "message": err.Error()})

		return
	}
	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{"title": "Show article", "article": article})

}
