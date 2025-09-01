package controllers

import (
	"net/http"
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

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home page",
		"featured": controller.articlesService.GetFeaturedArticles(),
		"stored":   controller.articlesService.GetStoredArticles(),
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"featured": controller.articlesService.GetFeaturedArticles(),
	// 	"stored":   controller.articlesService.GetStoredArticles(),
	// })

}
