package routes

import (
	articlesController "v2/internal/modules/article/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	articlesController := articlesController.New()

	router.GET("/articles/:id", articlesController.Show)

}
