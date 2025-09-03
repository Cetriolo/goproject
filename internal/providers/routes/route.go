package routes

import (
	articlesRoutes "v2/internal/modules/article/routes"
	homeRoutes "v2/internal/modules/home/routes"
	userRoutes "v2/internal/modules/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articlesRoutes.Routes(router)
	userRoutes.Routes(router)
}
