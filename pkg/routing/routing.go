package routing

import (
	homeRoutes "v2/internal/providers/routes"

	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	homeRoutes.RegisterRoutes(GetRouter())
}
